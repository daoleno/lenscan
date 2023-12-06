import duckdb from "@/lib/duckdb"

import "server-only"

import { PublicationTypesDistribution } from "./publicationTypesDistribution"

export async function getPublicationTypesDistribution() {
  const sql = `
    SELECT 
    publication_type,
    COUNT(*) AS count
    FROM 
    publication_record
    GROUP BY 
    publication_type
    ORDER BY 
    count DESC;
  `
  console.log(sql)
  const pubs = await duckdb.all(sql)

  // convert bigint to number and format date to ISO string
  pubs.forEach((a) => {
    a.count = Number(a.count)
  })

  return pubs as PublicationTypesDistribution[]
}
