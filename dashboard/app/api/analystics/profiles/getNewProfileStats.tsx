import { duckdb } from "@/lib/duckdb"

import "server-only"

import {
  DateRangeKey,
  getDateRangeCondition,
  getPreviousDateRangeCondition,
} from "../utils"

export type NewProfileStat = {
  day: string
  newProfiles: number
}

export async function getDailyNewProfileStats(rangeKey: DateRangeKey) {
  let sql = `
    SELECT 
      DATE_TRUNC('day', block_timestamp)::date as day, 
      COUNT(*) AS newProfiles
    FROM profile_record
  `

  sql += getDateRangeCondition(rangeKey, "block_timestamp")

  sql += `
    GROUP BY day
    ORDER BY day
  `

  const activities = await duckdb.all(sql)

  const chartData = activities.map((a) => ({
    day: new Date(a.day).toLocaleDateString(),
    newProfiles: Number(a.newProfiles),
  }))

  return chartData as NewProfileStat[]
}

export async function getProfilesGrowthPercentage(
  rangeKey: DateRangeKey = "ALL"
) {
  // Define SQL query for the current period
  let currentPeriodSql = `
    SELECT COUNT(*) as total
    FROM profile_record
  `
  currentPeriodSql += getDateRangeCondition(rangeKey, "block_timestamp")

  // Define SQL query for the previous period
  let previousPeriodSql = `
    SELECT COUNT(*) as total
    FROM profile_record
  `
  previousPeriodSql += getPreviousDateRangeCondition(
    rangeKey,
    "block_timestamp"
  )

  // Execute queries
  const currentPeriodData = await duckdb.all(currentPeriodSql)
  const previousPeriodData = await duckdb.all(previousPeriodSql)

  // Extract total counts
  const currentTotal = Number(currentPeriodData[0].total || 0)
  const previousTotal = Number(previousPeriodData[0].total || 0)

  // Calculate growth percentage
  let growthPercentage = 0
  if (previousTotal > 0) {
    growthPercentage = ((currentTotal - previousTotal) / previousTotal) * 100
    growthPercentage = parseFloat(growthPercentage.toFixed(2)) // Formats to 2 decimal places
  }

  return growthPercentage
}
