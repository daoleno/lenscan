import os
import shutil
import time
from datetime import datetime

import duckdb
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

dataset_ref = bqclient.dataset("v2_polygon", project="lens-public-data")
dataset = bqclient.get_dataset(dataset_ref)

# DuckDB connection
conn = duckdb.connect(database="v2_polygon.db")
cursor = conn.cursor()

is_task_running = False


def convert_schema(table_schema):
    converted_schema = []
    for field in table_schema:
        if field.field_type == "RECORD" and field.name == "datastream_metadata":
            for subfield in field.fields:
                if subfield.name == "source_timestamp":
                    converted_schema.append(f"{subfield.name} BIGINT")
        else:
            converted_schema.append(f"{field.name} {field.field_type}")
    return converted_schema


def connect_database(db_name):
    """
    Connect to the specified DuckDB database.
    Return the connection object.
    """
    conn = duckdb.connect(db_name)
    print("Connected to database successfully.")
    return conn


def create_output_directory(base_name):
    """
    Create a new directory with a timestamp.
    Return the directory name.
    """
    date_str = datetime.now().strftime("%Y%m%d_%H%M%S")
    output_directory = f"{base_name}_{date_str}"
    os.makedirs(output_directory, exist_ok=True)
    print(f"Output directory set to: {output_directory}")
    return output_directory


def update_symbolic_link(symlink_path, target):
    """
    Update the symbolic link to the specified target directory.
    """
    if os.path.exists(symlink_path):
        if os.path.islink(symlink_path):
            os.unlink(symlink_path)
    os.symlink(os.path.abspath(target), symlink_path)
    print(f"Updated symbolic link {symlink_path} to: {target}")


def export_tables(conn, output_dir):
    """
    Export all tables to Parquet files in the given directory.
    """
    try:
        # Retrieve a list of all tables
        tables = conn.execute(
            "SELECT table_name FROM information_schema.tables WHERE table_schema = 'main'"
        ).fetchall()
        if not tables:
            print("No tables found in the database.")
            return

        print(f"Found {len(tables)} tables.")

        # Export each table to a Parquet file
        for table in tables:
            table_name = table[0]
            parquet_file_path = os.path.join(output_dir, f"{table_name}.parquet")
            conn.execute(
                f"COPY {table_name} TO '{parquet_file_path}' (FORMAT 'parquet')"
            )
            print(f"Exported {table_name} to {parquet_file_path}")

    except Exception as e:
        print(f"An error occurred: {e}")


def delete_old_dir(symlink_path, base_name):
    """
    Delete the old directories that are not linked by the symbolic link.
    """
    cur_link_target = os.path.realpath(symlink_path)
    parent_dir = os.path.dirname(cur_link_target)
    for item in os.listdir(parent_dir):
        dir_path = os.path.join(parent_dir, item)
        if (
            os.path.isdir(dir_path)
            and item.startswith(base_name)
            and os.path.abspath(dir_path) != cur_link_target
        ):
            shutil.rmtree(dir_path)
            print(f"Deleted old directory: {dir_path}")


def perform_sync_task():
    global conn, is_task_running

    # Check if task is already running
    if is_task_running:
        print(
            f"[{datetime.now()}] Another sync task is already running. Skipping this cycle."
        )
        return

    is_task_running = True
    print(f"[{datetime.now()}] Starting data sync...")

    tables = list(bqclient.list_tables(dataset))  # Convert to list only once
    total_tables = len(tables)
    for index, table in enumerate(tables, start=1):
        table_id = table.table_id
        table_ref = dataset_ref.table(table_id)

        # Check if table exists in DuckDB and create if not
        cursor.execute(
            f"SELECT count(*) FROM information_schema.tables WHERE table_name = '{table_id}'"
        )
        if cursor.fetchone()[0] == 0:
            # Fetch the table schema from BigQuery
            table_schema = bqclient.get_table(table_ref).schema
            # Convert schema, replacing RECORD type with individual fields
            converted_schema = convert_schema(table_schema)
            ddl = f"CREATE TABLE {table_id} ({', '.join(converted_schema)})"
            cursor.execute(ddl)

        # Attempt to retrieve the last sync timestamp from the DuckDB table
        last_timestamp_result = cursor.execute(
            f"SELECT MAX(source_timestamp) FROM {table_id}"
        ).fetchone()
        last_timestamp = (
            last_timestamp_result[0] if last_timestamp_result[0] is not None else 0
        )

        # Build query to fetch new or updated records from BigQuery
        # Fetch the table schema from BigQuery
        table_schema = bqclient.get_table(table_ref).schema

        # Generate list of fields, excluding 'datastream_metadata', but including 'datastream_metadata.source_timestamp'
        fields = [f.name for f in table_schema if f.name != "datastream_metadata"]
        fields.append("datastream_metadata.source_timestamp")

        query = f"""
        SELECT {', '.join(fields)} 
        FROM `{table_ref}`
        WHERE datastream_metadata.source_timestamp > {last_timestamp}
        """
        query_job = bqclient.query(query)
        iterator = query_job.result(page_size=10000)  # Fetch 1000 rows at a time
        page_num = 0
        for page in iterator.pages:
            page_num += 1
            items = list(page)
            df = pl.DataFrame(
                {
                    field.name: list(data)
                    for field, data in zip(table_schema, zip(*items))
                }
            )
            # insert into duckdb
            if df.height > 0:  # Check if DataFrame is not empty
                cursor.register("df", df)
                cursor.execute(f"INSERT INTO {table_id} SELECT * FROM df")
                print(
                    f"[{datetime.now()}] Synced table {index}/{total_tables}: {table_id}, Page: {page_num}, Rows: {len(items)}"
                )
            else:
                print(
                    f"[{datetime.now()}] No new records found for table {index}/{total_tables}: {table_id}"
                )

    print(f"[{datetime.now()}] Data sync completed.")
    is_task_running = False

    # export tables to parquet files
    output_directory = create_output_directory("v2_polygon")
    export_tables(conn, output_directory)
    update_symbolic_link("v2_polygon", output_directory)

    # Add this new deletion operation after updating the symbolic link.
    delete_old_dir("v2_polygon", "v2_polygon_")


# Schedule the task to run every 15 minutes
schedule.every(15).minutes.do(perform_sync_task)

# Run the task immediately on start
perform_sync_task()

# Run the scheduler
while True:
    schedule.run_pending()
    time.sleep(1)
