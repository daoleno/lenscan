import { Database } from "duckdb-async"

const duckdb = await Database.create(":memory:")

const parquetPath = "v2_polygon"

function toParquetSql(sql: string) {
  // Regular expression to find table names in the SQL query
  // This might need to be adjusted based on your specific SQL syntax and patterns
  const tableNameRegex = /FROM\s+([a-zA-Z0-9_]+)/g

  // Replace each table name with the read_parquet function
  const convertedSQL = sql.replace(tableNameRegex, function (match, tableName) {
    return `FROM read_parquet('${parquetPath}/${tableName}.parquet')`
  })
  console.log(convertedSQL)
  return convertedSQL
}

export { duckdb, parquetPath, toParquetSql }
