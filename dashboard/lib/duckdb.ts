import { Database } from "duckdb-async";

const duckdb = await Database.create("v2_polygon.db");

export default duckdb;