import argparse
import concurrent.futures
import glob
import os
import time
import traceback
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
parser.add_argument(
    "-t", "--table", help="Name of specific table to sync", required=False, default=None
)
args = parser.parse_args()
dataset_ref = bqclient.dataset("v2_polygon", project="lens-public-data")
dataset = bqclient.get_dataset(dataset_ref)
is_task_running = False

output_directory = args.output
os.makedirs(output_directory, exist_ok=True)


def map_bq_type_to_python(bq_type):
    type_mapping = {
        "STRING": pl.Utf8,
        "BYTES": pl.Utf8,
        "INTEGER": pl.Int64,
        "FLOAT": pl.Float64,
        "BOOLEAN": pl.Boolean,
        "TIMESTAMP": pl.Datetime,
        "DATE": pl.Date,
        "TIME": pl.Time,
        "DATETIME": pl.Datetime,
        "NUMERIC": pl.Float64,
        "BIGNUMERIC": pl.Float64,
        "GEOGRAPHY": pl.Utf8,
    }

    return type_mapping.get(bq_type, pl.Utf8)


def get_file_size(file_path):
    """Returns the size of the file in bytes."""
    return os.path.getsize(file_path)


def merge_files(
    table_directory, table_id, max_group_size=2 * (10**9), max_file_size=1 * (10**9)
):
    """Merge files in given directory with a memory limit."""
    parquet_files = glob.glob(os.path.join(table_directory, f"{table_id}_*.parquet"))

    current_group = []
    current_group_file_paths = []  # this will hold file paths for deletion
    current_group_size = 0

    for file_path in parquet_files:
        df = pl.read_parquet(file_path)

        # Skip if single DataFrame already exceeds max_file_size
        if df.estimated_size() > max_file_size:
            print(
                f"Skipping file {file_path}. Size exceeds the maximum limit per file."
            )
            continue

        # If the addition of this df would exceed the max_group_size
        # then process current_group and start a new group
        if current_group_size + df.estimated_size() > max_group_size:
            if current_group:  # only write/merge if there's data to write
                df_merge = pl.concat(current_group)
                merged_file_name = os.path.join(
                    table_directory,
                    f"{table_id}_{datetime.now().strftime('%Y%m%d%H%M%S')}_merged.parquet",
                )
                df_merge.write_parquet(merged_file_name)
                print(f"[{datetime.now()}] Merged file saved as {merged_file_name}.")

                # Remove old files
                for old_file_path in current_group_file_paths:
                    os.remove(old_file_path)

            # Start a new group
            current_group = [df]
            current_group_file_paths = [file_path]  # reset the file paths list
            current_group_size = df.estimated_size()
        else:
            current_group.append(df)
            current_group_file_paths.append(file_path)
            current_group_size += df.estimated_size()

    # Don't forget the last group if it is not empty
    if current_group:
        df_merge = pl.concat(current_group)
        merged_file_name = os.path.join(
            table_directory,
            f"{table_id}_{datetime.now().strftime('%Y%m%d%H%M%S')}_merged.parquet",
        )
        df_merge.write_parquet(merged_file_name)
        print(f"[{datetime.now()}] Merged file saved as {merged_file_name}.")

        for old_file_path in current_group_file_paths:
            os.remove(old_file_path)


def sync_table(table_item, index, total_tables):
    try:
        print(
            f"[{datetime.now()}] Starting data sync of table {index}/{total_tables}: {table_item.table_id}"
        )
        last_timestamp = 0
        table_id = table_item.table_id
        table_directory = os.path.join(output_directory, table_id)
        os.makedirs(table_directory, exist_ok=True)  # Ensure the directory exists

        # Check if parquet files in the table directory need to be merged
        merge_files(table_directory, table_id)

        table_ref = dataset_ref.table(table_id)
        table = bqclient.get_table(table_ref)  # get table object

        # If the parquet files exist, get the maximum timestamp from the latest file
        parquet_files = sorted(
            glob.glob(os.path.join(table_directory, f"{table_id}_*.parquet"))
        )
        if parquet_files:
            df_old = pl.read_parquet(parquet_files[-1])  # read the latest file
            if "source_timestamp" in df_old.columns:
                last_timestamp = df_old["source_timestamp"].max()

        # Generate list of fields and their types, preserving the original schema's order.
        fields = [
            (
                f.name if f.name != "datastream_metadata" else "source_timestamp",
                pl.Float64
                if f.name == "datastream_metadata"
                else map_bq_type_to_python(f.field_type),
            )
            for f in table.schema
        ]

        # Initial query part
        field_names_for_query = [
            "datastream_metadata.source_timestamp"
            if field == "source_timestamp"
            else field
            for field, _ in fields
        ]
        query = f"SELECT {', '.join(field_names_for_query)} FROM `{table_ref}` WHERE datastream_metadata.source_timestamp > {last_timestamp} ORDER BY datastream_metadata.source_timestamp ASC"
        print(f"[{datetime.now()}] Query: {query}")

        # Modify the query if --sample is set
        if args.sample:
            query = f"SELECT {', '.join(fields)} FROM `{table_ref}` ORDER BY RAND() LIMIT 1000"

        query_job = bqclient.query(query)
        iterator = query_job.result(page_size=100000)

        # Configure Arrow writer to append to the .parquet file
        pages_received = 0
        for page in iterator.pages:
            pages_received += 1
            items = list(page)
            page_size = len(items)

            if page_size == 0:
                print(
                    f"[{datetime.now()}] No data received for table {index}/{total_tables}: {table_id}"
                )
                continue

            # Create DataFrame with explicit types
            data = {
                field: [item.get(field, None) for item in items] for field, _ in fields
            }

            df = pl.DataFrame(
                {
                    field: pl.Series(name=field, values=data[field]).map_elements(
                        lambda x: True
                        if str(x).lower() == "true"
                        else False
                        if str(x).lower() == "false"
                        else x,
                        return_dtype=bq_type,
                    )
                    for field, bq_type in fields
                }
            )

            # Write DataFrame to a new parquet file with timestamp in filename
            parquet_file_path = os.path.join(
                table_directory,
                f"{table_id}_{datetime.now().strftime('%Y%m%d%H%M%S')}.parquet",
            )
            df.write_parquet(parquet_file_path)

            print(
                f"[{datetime.now()}] Data sync of table {index}/{total_tables}: {table_id} - Received page {pages_received} with size {page_size}."
            )

        print(
            f"[{datetime.now()}] Data sync of table {index}/{total_tables}: {table_id} completed."
        )
    except Exception as e:
        print(
            f"[{datetime.now()}] Data sync of table {index}/{total_tables}: {table_id} failed."
        )
        traceback.print_exc()
        os._exit(1)


def perform_sync_task():
    global is_task_running

    if is_task_running:
        print(
            f"[{datetime.now()}] Another sync task is already running. Skipping this cycle."
        )
        return

    is_task_running = True
    print(f"[{datetime.now()}] Starting data sync...")
    if args.table:  # table is specified
        table_ref = dataset_ref.table(args.table)
        table = bqclient.get_table(table_ref)
        tables = [table]
    else:
        tables = list(bqclient.list_tables(dataset))
    total_tables = len(tables)

    with concurrent.futures.ThreadPoolExecutor(
        max_workers=args.concurrency
    ) as executor:
        for index, table_item in enumerate(tables, start=1):
            executor.submit(sync_table, table_item, index, total_tables)

    is_task_running = False
    print(f"[{datetime.now()}] Data sync completed.")


schedule.every(12).hours.do(perform_sync_task)

perform_sync_task()

while True:
    schedule.run_pending()
    time.sleep(1)
