import duckdb from "@/lib/duckdb"

import "server-only"

import { UserActivity } from "./userActivity"

export async function getUesrActivities() {
  const sql = `
    SELECT DATE_TRUNC('day', block_timestamp)::date as day, COUNT(*) AS count
    FROM publication_record
    GROUP BY day
    ORDER BY day
  `
  console.log(sql)
  const activities = await duckdb.all(sql)

  // convert bigint to number and format date to ISO string
  activities.forEach((a) => {
    a.count = Number(a.count)
    a.day = new Date(a.day).toLocaleDateString()
  })

  return activities as UserActivity[]
}
