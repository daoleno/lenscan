import { duckdb, toParquetSql } from "@/lib/duckdb"

import { Profile } from "./profile"

type getProfilesSort = {
  column?: string
  order?: "ASC" | "DESC" | "asc" | "desc"
}

type getProfilesParams = {
  limit: number
  offset: number
  sort?: getProfilesSort
}

export default async function getProfiles(
  params: getProfilesParams
): Promise<{ totalCount: number; profiles: Profile[] }> {
  const { limit, offset, sort } = params

  let sortOrder = sort ? `ORDER BY ${sort.column} ${sort.order}` : ""
  const sql = `SELECT * FROM profile_record ${sortOrder} LIMIT ${limit} OFFSET ${offset}`
  const profiles = await duckdb.all(toParquetSql(sql))

  const totalCount = await duckdb.all(
    toParquetSql(`SELECT COUNT(*) AS count FROM profile_record`)
  )

  return {
    totalCount: totalCount[0].count,
    profiles: profiles as Profile[],
  }
}
