import { duckdb, toParquetSql } from "@/lib/duckdb"
import lensClient from "@/lib/lensclient"
import { getIPFSURL } from "@/lib/utils"

export type Profile = {
  profile_id: string
  owned_by: string
  is_burnt: boolean
  tx_hash: string
  block_hash: string
  block_number: number
  log_index: number
  tx_index: number
  block_timestamp: Date
  source_timestamp: bigint

  profile_picture: string
  profile_handle: string
}

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

  // fetch profile other metadata
  const profileIds = profiles.map((p) => p.profile_id)
  const fullProfiles = await lensClient.profile.fetchAll({
    where: {
      profileIds,
    },
  })

  // merge
  profiles.forEach((p) => {
    const fullProfile = fullProfiles.items.find((fp) => fp.id === p.profile_id)
    if (fullProfile) {
      p.profile_picture = getIPFSURL(fullProfile?.metadata?.picture as any)
      p.profile_handle = fullProfile?.handle?.suggestedFormatted.localName
    }
  })

  console.log("profiles", profiles)

  return {
    totalCount: totalCount[0].count,
    profiles: profiles as Profile[],
  }
}
