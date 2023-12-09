import os

import duckdb

try:
    # Connect to the DuckDB database
    conn = duckdb.connect("v2_polygon.db")
    print("Connected to database successfully.")

    # Specify the output directory
    output_directory = "v2_polygon"
    os.makedirs(output_directory, exist_ok=True)
    print(f"Output directory set to: {output_directory}")

    # Retrieve a list of all tables
    tables = conn.execute(
        "SELECT table_name FROM information_schema.tables WHERE table_schema = 'main'"
    ).fetchall()
    if not tables:
        print("No tables found in the database.")
    else:
        print(f"Found {len(tables)} tables.")

    # Export each table to a Parquet file
    for table in tables:
        table_name = table[0]
        parquet_file_path = os.path.join(output_directory, f"{table_name}.parquet")
        conn.execute(f"COPY {table_name} TO '{parquet_file_path}' (FORMAT 'parquet')")
        print(f"Exported {table_name} to {parquet_file_path}")

except Exception as e:
    print(f"An error occurred: {e}")
