import argparse
import os

import duckdb

# Parsing command line arguments
parser = argparse.ArgumentParser()
parser.add_argument("--db", help="Path to the DuckDB database", required=True)
args = parser.parse_args()

# Initialize connection
conn = duckdb.connect(args.db)

# Get all tables
tables = conn.execute(
    "SELECT table_name FROM information_schema.tables WHERE table_schema='main'"
).fetchall()

# ensure export directory exists
if not os.path.exists("sample"):
    os.makedirs("sample")


# Iterate over tables
for table in tables:
    table_name = table[0]
    # Get random 1000 rows from table and export it to parquet file
    conn.execute(
        f"COPY (SELECT * FROM {table_name} ORDER BY RANDOM() LIMIT 1000) TO 'sample/{table_name}.parquet' (FORMAT 'parquet')"
    )

print("Export completed.")
