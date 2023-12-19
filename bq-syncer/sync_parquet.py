import argparse
import concurrent.futures
import os
import sys
import time
import traceback
from datetime import datetime

import polars as pl
import pyarrow.parquet as pq
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
parser.add_argument(
    "-s",
    "--sample",
    help="If true, it exports a maximum of 1000 random sample data from each table",
    action="store_true",
    default=False,
)
parser.add_argument(
    "-c",
    "--concurrency",
    help="Number of concurrent tasks.",
    type=int,
    default=1,
)
args = parser.parse_args()
dataset_ref = bqclient.dataset("v2_polygon", project="lens-public-data")
dataset = bqclient.get_dataset(dataset_ref)
is_task_running = False

output_directory = args.output
os.makedirs(output_directory, exist_ok=True)


def map_bq_type_to_python(bq_type):
    type_mapping = {
        "STRING": str,
        "BYTES": bytes,
        "INTEGER": int,
        "FLOAT": float,
        "BOOLEAN": bool,
        "TIMESTAMP": datetime,
    }
    return type_mapping.get(bq_type, str)


def sync_table(table_item, index, total_tables):
    if table_item.table_id != "profile_metadata":
        return

    last_timestamp = 0
    table_id = table_item.table_id
    table_ref = dataset_ref.table(table_id)
    table = bqclient.get_table(table_ref)  # get table object

    parquet_file_path = os.path.join(output_directory, f"{table_id}.parquet")

    # If the parquet file exists, get the maximum timestamp
    if os.path.exists(parquet_file_path):
        df_old = pl.read_parquet(parquet_file_path)
        if "datastream_metadata.source_timestamp" in df_old.columns:
            last_timestamp = df_old["datastream_metadata.source_timestamp"].max()

    # Generate list of fields, preserving the original schema's order.
    fields = [
        f.name
        if f.name != "datastream_metadata"
        else "datastream_metadata.source_timestamp"
        for f in table.schema
    ]
    try:
        # Initial query part
        query = f"SELECT {', '.join(fields)} FROM `{table_ref}` WHERE datastream_metadata.source_timestamp > {last_timestamp}"

        # Modify the query if --sample is set
        if args.sample:
            query = f"SELECT {', '.join(fields)} FROM `{table_ref}` ORDER BY RAND() LIMIT 1000"

        query_job = bqclient.query(query)
        iterator = query_job.result(page_size=10000)

        pages_received = 0
        for page in iterator.pages:
            pages_received += 1
            print(
                f"[{datetime.now()}] Processing table {index}/{total_tables}: {table_id} - Page {pages_received}"
            )
            items = list(page)
            if len(items) == 0:
                print(
                    f"[{datetime.now()}] No data received for table {index}/{total_tables}: {table_id}"
                )
                continue
            df = pl.DataFrame({field: data for field, data in zip(fields, zip(*items))})
            if os.path.exists(parquet_file_path):
                df_old = pl.read_parquet(parquet_file_path)
                df = df.select(
                    [
                        pl.col(col).cast(df_old[col].dtype)
                        if col in df_old.columns
                        else pl.col(col)
                        for col in df.columns
                    ]
                )

                df = pl.concat([df_old, df])
            df.write_parquet(parquet_file_path)

        print(
            f"[{datetime.now()}] Data sync of table {index}/{total_tables}: {table_id} completed."
        )
    except Exception as e:
        # Print more verbose error information
        print(f"An error occurred while processing table {table_id}: {e}")
        traceback.print_exc()  # This prints the stack trace
        sys.exit(1)


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
        with concurrent.futures.ThreadPoolExecutor(
            max_workers=args.concurrency
        ) as executor:
            for index, table_item in enumerate(tables, start=1):
                executor.submit(sync_table, table_item, index, total_tables)
    except Exception as e:
        print(f"An error occurred: {e}")
    is_task_running = False
    print(f"[{datetime.now()}] Data sync completed.")


schedule.every(60).minutes.do(perform_sync_task)

perform_sync_task()

while True:
    schedule.run_pending()
    time.sleep(1)
