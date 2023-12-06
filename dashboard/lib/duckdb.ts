import * as duck from "duckdb"
import { Database } from "duckdb-async"

const duckdb = await Database.create("v2_polygon.db", duck.OPEN_READONLY)

export default duckdb
