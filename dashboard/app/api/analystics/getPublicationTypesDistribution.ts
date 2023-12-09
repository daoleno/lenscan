import duckdb from "@/lib/duckdb"

export type PublicationTypesDistribution = {
  publication_type: string
  count: number
}

export async function getPublicationTypesDistribution() {
  let sql = `
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

  pubs.forEach((a) => {
    a.count = Number(a.count)
  })

  return pubs as PublicationTypesDistribution[]
}
