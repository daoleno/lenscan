# BigQuery to DuckDB Sync Script

This is a Python script which syncs data from a Lens Google BigQuery dataset to a local DuckDB database at regular intervals.

## Install the necessary Python packages:

```sh
poetry install
```

## Setup the Service Account

Before running the script, make sure you have downloaded the service account JSON file:

- Follow https://cloud.google.com/docs/authentication/getting-started#creating_a_service_account to create a new service account and download the JSON key file.
- Place the JSON key file in the same directory as this script and rename it to 'service_account.json'.

## Running the Script

Run the script using a Python interpreter:

```sh
poetry shell
python sync.py
```

## How it Works

- On script start and every 15 minutes after, the script checks for new or updated rows in each table of the specified BigQuery dataset.
- Checks whether each table exists in the DuckDB database, creating it if not.
- Retrieves the BigQuery table schema, converts it from RECORD to individual fields and creates the table in DuckDB with this schema if it does not already exist.
- Retrieves the maximum `source_timestamp` value from the DuckDB table and fetches all rows from the BigQuery table that have a `source_timestamp` greater than this.
- Inserts the fetched rows into the corresponding DuckDB table.
