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

Export all tables into a postgres database:

```sh
poetry shell
python sync.py -p postgres://postgres:postgres@localhost:5432/lens_v2 -c 10
```

Export Sample Data into a postgres database:

```sh
poetry shell
python sync.py -p postgres://postgres:postgres@localhost:5432/lens_v2 -s 1000 -c 10
```

## How it Works

- On script start and at every scheduled interval, the script checks for new or updated rows in each table of the specified BigQuery dataset.
- Checks whether each table exists in the postgres database, creating it if not.
- Retrieves the BigQuery table schema, converts it from RECORD to individual fields and creates the table in postgres with this schema if it does not already exist.
- Retrieves the maximum `source_timestamp` value from the postgres table and fetches all rows from the BigQuery table that have a `source_timestamp` greater than this.
- Inserts the fetched rows into the corresponding postgres table.
