import duckdb from "@/lib/duckdb"

import { Publication } from "./publication"

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
  console.log(sql)
  const publications = await duckdb.all(sql)

  // Get total count if we have profile_id filter
  let totalCount
  if (filter && filter.profile_id) {
    const sql = `SELECT COUNT(*) AS count FROM publication_record ${filterCondition}`
    const result = await duckdb.all(sql)
    totalCount = result[0].count
  }

  return {
    totalCount,
    publications: publications as Publication[],
  }
}
