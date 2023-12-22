import { duckdb, parquetPath } from "@/lib/duckdb"

import "server-only"

import { ProfileFragment } from "@lens-protocol/client"

import lensClient from "@/lib/lensclient"

import { DateRangeKey, getDateRangeCondition } from "./utils"

type TopProfile = ProfileFragment & { weighted_score: number }

export async function getTopProfiles(rangeKey: DateRangeKey = "ALL") {
  let baseSql = `
    FROM read_parquet('${parquetPath}/[FILENAME]/*.parquet')
  `

  let sql = `
    WITH publication_counts AS (
        SELECT profile_id, COUNT(*) AS publication_count
        ${baseSql.replace("[FILENAME]", "publication_record")}
        ${getDateRangeCondition(rangeKey, "block_timestamp")}
        GROUP BY profile_id
    ), 
    reaction_counts AS (
        SELECT actioned_by_profile_id AS profile_id, COUNT(*) AS reaction_count
        ${baseSql.replace("[FILENAME]", "publication_reaction")}
        ${getDateRangeCondition(rangeKey, "action_at")}
        GROUP BY actioned_by_profile_id
    ),
    mention_counts AS (
        SELECT profile_id, COUNT(*) AS mention_count
        ${baseSql.replace("[FILENAME]", "publication_mention")}
        ${getDateRangeCondition(rangeKey, "timestamp")}
        GROUP BY profile_id
    ),
    combined_counts AS (
        SELECT 
            COALESCE(p.profile_id, r.profile_id, m.profile_id) AS profile_id,
            COALESCE(p.publication_count, 0) AS publication_count,
            COALESCE(r.reaction_count, 0) AS reaction_count,
            COALESCE(m.mention_count, 0) AS mention_count,
            0.5 * COALESCE(p.publication_count, 0) + 0.3 * COALESCE(r.reaction_count, 0) + 0.2 * COALESCE(m.mention_count, 0) AS weighted_score
        FROM publication_counts p
        FULL OUTER JOIN reaction_counts r ON p.profile_id = r.profile_id
        FULL OUTER JOIN mention_counts m ON p.profile_id = m.profile_id
    )
  `

  sql += `
    SELECT profile_id, publication_count, reaction_count, mention_count, weighted_score
    FROM combined_counts
    ORDER BY weighted_score DESC
    LIMIT 5
  `

  const topProfiles = await duckdb.all(sql)
  if (!topProfiles.length) return []

  const fetchResults = await lensClient.profile.fetchAll({
    where: {
      profileIds: topProfiles.map((c) => c.profile_id),
    },
  })

  // add weighted score to profiles
  const profiles = fetchResults.items.map((profile) => {
    const contributor = topProfiles.find((c) => c.profile_id === profile.id)
    return {
      ...profile,
      weighted_score: Math.round(contributor?.weighted_score),
    }
  })

  return profiles as TopProfile[]
}
