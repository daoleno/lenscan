import { Database } from "duckdb-async"

const duckdb = await Database.create(":memory:")

const parquetPath = process.env.PARQUET_DIR_PATH

function toParquetSql(sql: string) {
  const tableNameRegex = /(?:FROM|JOIN)\s+([a-zA-Z0-9_]+)/g

  const convertedSQL = sql.replace(tableNameRegex, function (match, tableName) {
    return ` ${match.slice(
      0,
      4
    )} read_parquet('${parquetPath}/${tableName}/*.parquet')`
  })
  console.log(convertedSQL)
  return convertedSQL
}

export { duckdb, parquetPath, toParquetSql }
