import { duckdb } from "@/lib/duckdb"

import "server-only"

export type ContentLanguageDistribution = {
  language: string
  frequency: number
}

export async function getContentLanguageDistribution() {
  const sql = `
      SELECT language, COUNT(*) AS frequency
      FROM publication_metadata
      GROUP BY language
      ORDER BY frequency DESC;
  `

  const result = await duckdb.all(sql)

  // convert bigint to number
  result.forEach((row) => {
    row.frequency = Number(row.frequency)
  })
  return result as ContentLanguageDistribution[]
}
