import argparse
import os
import sys
import time
from datetime import datetime

import polars as pl
import schedule
from google.cloud import bigquery
from google.oauth2 import service_account

# Initialize BigQuery client
service_account_path = "service_account.json"
credentials = service_account.Credentials.from_service_account_file(
    service_account_path
)
bqclient = bigquery.Client(credentials=credentials, project=credentials.project_id)

# Parse command line arguments
parser = argparse.ArgumentParser(description="Connect to DuckDB and export tables.")
parser.add_argument(
    "-o",
    "--output",
    help="Path to the directory to which exports will be saved.",
    required=True,
)
args = parser.parse_args()
dataset_ref = bqclient.dataset("v2_polygon", project="lens-public-data")
dataset = bqclient.get_dataset(dataset_ref)
is_task_running = False

output_directory = args.output
os.makedirs(output_directory, exist_ok=True)


def perform_sync_task():
    global is_task_running

    if is_task_running:
        print(
            f"[{datetime.now()}] Another sync task is already running. Skipping this cycle."
        )
        return

    is_task_running = True
    print(f"[{datetime.now()}] Starting data sync...")
    tables = list(bqclient.list_tables(dataset))
    total_tables = len(tables)

    try:
        for index, table_item in enumerate(tables, start=1):
            last_timestamp = 0
            table_id = table_item.table_id
            table_ref = dataset_ref.table(table_id)
            table = bqclient.get_table(table_ref)  # get table object

            parquet_file_path = os.path.join(output_directory, f"{table_id}.parquet")

            # If the parquet file exists, get the maximum timestamp
            if os.path.exists(parquet_file_path):
                df_old = pl.read_parquet(parquet_file_path)

                # Check if column 'source_timestamp' exists in the DataFrame
                if "datastream_metadata.source_timestamp" in df_old.columns:
                    last_timestamp = df_old[
                        "datastream_metadata.source_timestamp"
                    ].max()

            # Generate list of fields, preserving the original schema's order.
            fields = [
                f.name
                if f.name != "datastream_metadata"
                else "datastream_metadata.source_timestamp"
                for f in table.schema
            ]

            try:
                query = f"SELECT {', '.join(fields)} FROM `{table_ref}` WHERE datastream_metadata.source_timestamp > {last_timestamp}"
                query_job = bqclient.query(query)
                iterator = query_job.result(page_size=10000)

                pages_received = 0

                for page in iterator.pages:
                    pages_received += 1
                    print(
                        f"[{datetime.now()}] Processing table {index}/{total_tables}: {table_id} - Page {pages_received}"
                    )

                    items = list(page)

                    # If no data, skip the loop
                    if len(items) == 0:
                        print(
                            f"[{datetime.now()}] No data received for table {index}/{total_tables}: {table_id}"
                        )
                        continue

                    df = pl.DataFrame(
                        {field: data for field, data in zip(fields, zip(*items))}
                    )

                    if os.path.exists(parquet_file_path):
                        df_old = pl.read_parquet(parquet_file_path)
                        df = df_old.vstack(df)

                    df.write_parquet(parquet_file_path)

                print(
                    f"[{datetime.now()}] Data sync of table {index}/{total_tables}: {table_id} completed."
                )
            except Exception as table_related_error:
                print(
                    f"An error occurred while processing table {table_id}: {table_related_error}"
                )
                sys.exit(1)

    except Exception as e:
        print(f"An error occurred: {e}")
    is_task_running = False
    print(f"[{datetime.now()}] Data sync completed.")


schedule.every(60).minutes.do(perform_sync_task)

perform_sync_task()

while True:
    schedule.run_pending()
    time.sleep(1)
