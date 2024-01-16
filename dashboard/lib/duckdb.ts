import { Database } from "duckdb-async"

const accessMode = 0 // 1 = read/write, 0 = read-only
const duckdb = await Database.create(process.env.DUCKDB_PATH, {
  access_mode: "READ_ONLY",
})

export { duckdb }
