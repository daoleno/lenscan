import argparse
import logging
from concurrent.futures import ThreadPoolExecutor

from google.cloud import bigquery
from google.oauth2 import service_account
from psycopg_pool import ConnectionPool

logging.basicConfig(level=logging.INFO)


# Parse command line arguments
parser = argparse.ArgumentParser(description="Sync data from BigQuery to PostgreSQL.")
parser.add_argument(
    "-p", "--pg-dsn", help="PostgreSQL DSN (Data Source Name).", required=True
)
parser.add_argument(
    "-c",
    "--concurrency",
    help="Number of concurrent tasks.",
    type=int,
    default=1,
)
parser.add_argument(
    "-s",
    "--sample-size",
    help="Number of records to sample for each table (for debugging). 0 for no limit.",
    type=int,
    default=0,
)
args = parser.parse_args()


# Initialize BigQuery client
service_account_path = "service_account.json"
credentials = service_account.Credentials.from_service_account_file(
    service_account_path
)
bqclient = bigquery.Client(credentials=credentials, project=credentials.project_id)


dataset_ref = bqclient.dataset("v2_polygon", project="lens-public-data")
dataset = bqclient.get_dataset(dataset_ref)


supported_tables = {
    "publication_metadata",
    "publication_record",
    "publication_hashtag",
    "profile_record",
    "publication_reaction",
    "profile_revenue",
    "profile_revenue_record",
}


is_task_running = False


def convert_schema(table_schema):
    converted_schema = []
    type_mapping = {
        "STRING": "TEXT",
        "INTEGER": "INTEGER",
        "FLOAT": "REAL",
        "BOOLEAN": "BOOLEAN",
        "TIMESTAMP": "TIMESTAMPTZ",  # Mapping TIMESTAMP to TIMESTAMPTZ
        "NUMERIC": "NUMERIC",  # Mapping NUMERIC directly
        "JSON": "JSONB",  # Mapping JSON to JSONB
        # Add other known BigQuery data types as needed
    }

    for field in table_schema:
        if field.field_type == "RECORD" and field.name == "datastream_metadata":
            # Specifically handle the 'datastream_metadata' RECORD type
            for subfield in field.fields:
                if subfield.name == "source_timestamp":
                    # Convert 'source_timestamp' within 'datastream_metadata' to BIGINT
                    converted_schema.append(f"{subfield.name} BIGINT")
        elif field.field_type in type_mapping:
            # Handle other known types normally
            converted_type = type_mapping[field.field_type]
            converted_schema.append(f"{field.name} {converted_type}")
        else:
            # Raise an error for unsupported or unknown types
            raise ValueError(
                f"Unsupported or unknown BigQuery data type: {field.field_type}"
            )

    return converted_schema


def process_table(table, pool, dataset_ref, bqclient, idx, sample_size):
    try:
        with pool.connection() as conn:
            table_id = table.table_id
            table_ref = dataset_ref.table(table_id)
            table_schema = bqclient.get_table(table_ref).schema
            cursor = conn.cursor()

            # Convert BigQuery schema to PostgreSQL schema
            converted_schema = convert_schema(table_schema)

            # Check if table exists in PostgreSQL and create it if not
            cursor.execute(f"SELECT to_regclass('public.{table_id}');")
            if cursor.fetchone()[0] is None:
                logging.info(f"Creating table: {table_id}")
                ddl = f"CREATE TABLE {table_id} ({', '.join(converted_schema)});"
                cursor.execute(ddl)
                conn.commit()

            # Retrieve last synced timestamp
            cursor.execute(f"SELECT MAX(source_timestamp) FROM {table_id};")
            last_timestamp_result = cursor.fetchone()
            last_timestamp = (
                last_timestamp_result[0] if last_timestamp_result[0] is not None else 0
            )

            # Build and execute BigQuery query
            fields = [
                f.name
                if f.name != "datastream_metadata"
                else "datastream_metadata.source_timestamp"
                for f in table_schema
            ]

            limit_clause = f"LIMIT {sample_size}" if sample_size > 0 else ""
            query = f"""
            SELECT {', '.join(fields)}
            FROM `{table_ref}`
            WHERE datastream_metadata.source_timestamp > {last_timestamp}
            ORDER BY datastream_metadata.source_timestamp ASC
            {limit_clause}
            """
            query_job = bqclient.query(query)

            # Process and insert data into PostgreSQL
            pageNum = 0
            for page in query_job.result(page_size=10000).pages:
                pageNum += 1
                items = list(page)
                if items:
                    # Convert items to a format suitable for insertion into PostgreSQL
                    # This step needs to be adapted based on your actual table schema and data types
                    insert_values = [
                        # Convert each item to a tuple or another format as needed for insertion
                        tuple(item.values())
                        for item in items
                    ]
                    # Generate the INSERT INTO statement with placeholders for values
                    placeholders = ", ".join(["%s"] * len(items[0]))
                    insert_query = f"INSERT INTO {table_id} VALUES ({placeholders});"
                    cursor.executemany(insert_query, insert_values)
                    conn.commit()
                    logging.info(
                        f"Synced table {idx}: {table_id} - page {pageNum} - {len(items)} rows"
                    )

            logging.info(f"Completed syncing table {idx}: {table_id}")

    except Exception as e:
        logging.error(f"An error occurred while processing table {table_id}: {e}")
    finally:
        cursor.close()


def perform_sync_task():
    global is_task_running

    if is_task_running:
        logging.info("Another sync task is already running. Skipping this cycle.")
        return

    is_task_running = True
    logging.info("Starting data sync...")

    pg_dsn = args.pg_dsn
    sample_size = args.sample_size
    with ConnectionPool(conninfo=pg_dsn, min_size=1, max_size=100) as pool:
        try:
            all_tables = list(bqclient.list_tables(dataset))
            # Filter tables to include only the supported ones
            tables = [
                table for table in all_tables if table.table_id in supported_tables
            ]
            logging.info(f"Found {len(tables)} supported tables in BigQuery.")

            with ThreadPoolExecutor(max_workers=args.concurrency) as executor:
                for idx, table in enumerate(tables, start=1):
                    executor.submit(
                        process_table,
                        table,
                        pool,
                        dataset_ref,
                        bqclient,
                        idx,
                        sample_size,
                    )

            logging.info(f"Synced {len(tables)} tables.")

        except Exception as e:
            logging.error(f"An error occurred: {e}")
        finally:
            is_task_running = False


perform_sync_task()
