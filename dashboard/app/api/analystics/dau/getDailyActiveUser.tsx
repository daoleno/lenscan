import { duckdb, toParquetSql } from "@/lib/duckdb"

import {
  DateRangeKey,
  getDateRangeCondition,
  getPreviousDateRangeCondition,
} from "../utils"

export type DauStats = {
  day: string
  polygon: number
  momoka: number
  total: number
}

export async function getDauStats(rangeKey: DateRangeKey) {
  let sql = `
    SELECT 
      DATE_TRUNC('day', block_timestamp)::date as day, 
      COUNT(DISTINCT CASE WHEN is_momoka THEN profile_id END) AS momoka,
      COUNT(DISTINCT CASE WHEN NOT is_momoka THEN profile_id END) AS polygon
    FROM publication_record
  `
  sql += getDateRangeCondition(rangeKey, "block_timestamp")
  sql += `
    GROUP BY day
    ORDER BY day
  `

  const activeUsers = await duckdb.all(toParquetSql(sql))

  const chartData = activeUsers.map((a) => ({
    day: new Date(a.day).toLocaleDateString(),
    polygon: Number(a.polygon),
    momoka: Number(a.momoka),
    total: Number(a.polygon) + Number(a.momoka),
  }))

  return chartData as DauStats[]
}

export async function getDauGrowthPercentages(rangeKey: DateRangeKey) {
  // Define SQL query for the current period
  let currentPeriodSql = `
    SELECT COUNT(DISTINCT profile_id) as total
    FROM publication_record
  `
  currentPeriodSql += getDateRangeCondition(rangeKey, "block_timestamp")

  // Define SQL query for the previous period
  let previousPeriodSql = `
    SELECT COUNT(DISTINCT profile_id) as total
    FROM publication_record
  `
  previousPeriodSql += getPreviousDateRangeCondition(
    rangeKey,
    "block_timestamp"
  )

  // Execute queries
  const currentPeriodData = await duckdb.all(toParquetSql(currentPeriodSql))
  const previousPeriodData = await duckdb.all(toParquetSql(previousPeriodSql))

  // Extract total counts
  const currentTotal = Number(currentPeriodData[0].total || 0)
  const previousTotal = Number(previousPeriodData[0].total || 0)

  // Calculate growth percentage
  let growthPercentage = 0
  if (previousTotal > 0) {
    growthPercentage = parseFloat(
      (((currentTotal - previousTotal) / previousTotal) * 100).toFixed(2)
    )
  }

  return growthPercentage
}
