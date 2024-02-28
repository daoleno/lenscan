import db from "@/lib/db"
import { sql } from "drizzle-orm"

import "server-only"

export type ContentLanguageDistribution = {
  language: string
  frequency: number
}

export async function getContentLanguageDistribution() {
  const statement = `
      SELECT language, COUNT(*) AS frequency
      FROM publication_metadata
      GROUP BY language
      ORDER BY frequency DESC;
  `

  const result = await db.execute(sql.raw(statement)) as ContentLanguageDistribution[]

  return result.map((a) => ({
    ...a,
    frequency: Number(a.frequency),
  }))
}
