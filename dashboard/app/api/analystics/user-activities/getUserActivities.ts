import duckdb from "@/lib/duckdb"

import "server-only"

import { DateRangeKey, getDateRangeCondition } from "../utils"

export type UserActivity = {
  day: string
  count: number
}

// Update rangeKey parameter type
export async function getUserActivities(rangeKey: DateRangeKey = "ALL") {
  let sql = `
    SELECT DATE_TRUNC('day', block_timestamp)::date as day, COUNT(*) AS count
    FROM publication_record
  `

  sql += getDateRangeCondition(rangeKey, "block_timestamp")

  sql += `
    GROUP BY day
    ORDER BY day
  `

  console.log(sql)
  const activities = await duckdb.all(sql)

  activities.forEach((a) => {
    a.count = Number(a.count)
    a.day = new Date(a.day).toLocaleDateString()
  })

  return activities as UserActivity[]
}
