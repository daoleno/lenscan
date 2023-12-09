import duckdb from "@/lib/duckdb"

import "server-only"

export type Contributor = {
  profile_id: string
  publication_count: number
  reaction_count: number
  mention_count: number
  total_count: number
}

export async function getTopContributors() {
  let sql = `
    WITH publication_counts AS (
        SELECT profile_id, COUNT(*) AS publication_count
        FROM publication_record
        GROUP BY profile_id
    ), 
    reaction_counts AS (
        SELECT actioned_by_profile_id AS profile_id, COUNT(*) AS reaction_count
        FROM publication_reaction
        GROUP BY actioned_by_profile_id
    ),
    mention_counts AS (
        SELECT profile_id, COUNT(*) AS mention_count
        FROM publication_mention
        GROUP BY profile_id
    ),
    combined_counts AS (
        SELECT 
            COALESCE(p.profile_id, r.profile_id, m.profile_id) AS profile_id,
            COALESCE(p.publication_count, 0) AS publication_count,
            COALESCE(r.reaction_count, 0) AS reaction_count,
            COALESCE(m.mention_count, 0) AS mention_count,
            COALESCE(p.publication_count, 0) + COALESCE(r.reaction_count, 0) + COALESCE(m.mention_count, 0) AS total_count
        FROM publication_counts p
        FULL OUTER JOIN reaction_counts r ON p.profile_id = r.profile_id
        FULL OUTER JOIN mention_counts m ON p.profile_id = m.profile_id
    )
  `

  sql += `
    SELECT profile_id, publication_count, reaction_count, mention_count, total_count
    FROM combined_counts
    ORDER BY total_count DESC
    LIMIT 10
  `

  console.log(sql)
  const contributors = await duckdb.all(sql)

  contributors.forEach((a) => {
    a.publication_count = Number(a.publication_count)
    a.reaction_count = Number(a.reaction_count)
    a.mention_count = Number(a.mention_count)
    a.total_count = Number(a.total_count)
  })

  return contributors as Contributor[]
}
