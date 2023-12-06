import duckdb from "@/lib/duckdb"

import "server-only"

import { Hashtag } from "./hashtag"

export async function getPopularHashtags() {
  const sql = `
    SELECT hashtag, COUNT(*) AS count
    FROM publication_hashtag
    GROUP BY hashtag
    ORDER BY count DESC
    LIMIT 20;

  `
  console.log(sql)
  const hashtags = await duckdb.all(sql)

  // convert bigint to number and format date to ISO string
  hashtags.forEach((a) => {
    a.count = Number(a.count)
  })

  return hashtags as Hashtag[]
}
