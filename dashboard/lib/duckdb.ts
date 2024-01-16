import { Database } from "duckdb-async"

const readonly = 1 // 1 = readonly, 0 = readwrite
const duckdb = await Database.create(
  process.env.DUCKDB_PATH as string,
  readonly
)

export { duckdb }
