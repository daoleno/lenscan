import datetime
import os
import shutil

import duckdb


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


def main():
    try:
        # Connect to the DuckDB database
        conn = duckdb.connect("v2_polygon.db")
        print("Connected to database successfully.")

        # Create a dated directory for this export
        date_str = datetime.datetime.now().strftime("%Y%m%d_%H%M%S")
        output_directory = f"v2_polygon_{date_str}"
        os.makedirs(output_directory, exist_ok=True)
        print(f"Output directory set to: {output_directory}")

        # Export tables
        export_tables(conn, output_directory)

        # Create or update the symbolic link to the latest export
        symlink_path = "v2_polygon_latest"
        if os.path.islink(symlink_path):
            os.unlink(symlink_path)
        os.symlink(output_directory, symlink_path)
        print(f"Updated symbolic link to: {symlink_path}")

    except Exception as e:
        print(f"An error occurred: {e}")

    # Optionally, add code to clean up old directories here


if __name__ == "__main__":
    main()
