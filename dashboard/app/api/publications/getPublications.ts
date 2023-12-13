import { duckdb, toParquetSql } from "@/lib/duckdb"
import lensClient from "@/lib/lensclient"
import { getIPFSURL } from "@/lib/utils"

type getPublicationsSort = {
  column?: string
  order?: "ASC" | "DESC" | "asc" | "desc"
}

type getPublicationsFilter = {
  app?: string[]
  publication_type?: string[]
  is_momoka?: string
  profile_id?: string[]
}

type getPublicationsParams = {
  limit: number
  offset: number
  sort?: getPublicationsSort
  filter?: getPublicationsFilter
}

export type Publication = {
  publication_id: string
  contract_publication_id: string
  publication_type: string
  profile_id: string
  content_uri: string
  parent_publication_id: string
  root_publication_id: string
  is_hidden: boolean
  is_momoka: boolean
  momoka_proof: string
  app: string
  gardener_flagged: boolean
  transaction_executor: string
  tx_hash: string
  block_hash: string
  block_number: number
  log_index: number
  tx_index: number
  block_timestamp: Date
  source_timestamp: number

  profile_picture: string
  profile_handle: string
}

export default async function getPublications(
  params: getPublicationsParams
): Promise<{
  totalCount: number | undefined
  publications: Publication[]
}> {
  const { limit, offset, sort, filter } = params

  let sortOrder = sort ? `ORDER BY ${sort.column} ${sort.order}` : ""

  let conditions = []
  if (filter) {
    for (const key in filter) {
      const value = filter[key as keyof getPublicationsFilter]

      if (value !== undefined && value.length > 0) {
        // Make sure the array is not empty
        if (Array.isArray(value)) {
          conditions.push(`${key} IN ('${value.join("','")}')`)
        } else {
          conditions.push(`${key}='${value}'`)
        }
      }
    }
  }

  let filterCondition = conditions.length
    ? " WHERE " + conditions.join(" AND ")
    : ""
  const sql = `SELECT * FROM publication_record ${filterCondition} ${sortOrder} LIMIT ${limit} OFFSET ${offset}`
  const publications = await duckdb.all(toParquetSql(sql))

  // fetch profile pictures
  const profileIds = publications.map((p) => p.profile_id)
  const profilePictures = await lensClient.profile.fetchAll({
    where: {
      profileIds,
    },
  })
  publications.forEach((p) => {
    const profile = profilePictures.items.find(
      (profile) => profile.id === p.profile_id
    )
    p.profile_picture = getIPFSURL(profile?.metadata?.picture as any)
    p.profile_handle = profile?.handle?.suggestedFormatted.localName
  })

  console.log("getPublications", sql, publications)

  // Get total count if we have profile_id filter
  let totalCount
  if (filter && filter.profile_id) {
    const sql = `SELECT COUNT(*) AS count FROM publication_record ${filterCondition}`
    const result = await duckdb.all(toParquetSql(sql))
    totalCount = result[0].count
  }

  return {
    totalCount,
    publications: publications as Publication[],
  }
}
