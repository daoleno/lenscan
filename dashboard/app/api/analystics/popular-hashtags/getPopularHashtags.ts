import { duckdb } from "@/lib/duckdb"

import { DateRangeKey, getDateRangeCondition } from "../utils"

import "server-only"

export type Hashtag = {
  hashtag: string
  count: number
}

export async function getPopularHashtags(rangeKey: DateRangeKey = "ALL") {
  let sql = `
    SELECT hashtag, COUNT(*) AS count
    FROM publication_hashtag
  `

  sql += getDateRangeCondition(rangeKey, "timestamp")

  sql += `
    GROUP BY hashtag
    ORDER BY count DESC
    LIMIT 20
  `

  const hashtags = await duckdb.all(sql)

  hashtags.forEach((a) => {
    a.count = Number(a.count)
  })

  return hashtags as Hashtag[]
}
