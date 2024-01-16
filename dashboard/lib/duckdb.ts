import { Database } from "duckdb-async"

const duckdb = await Database.create(process.env.DUCKDB_PATH)

export { duckdb }
