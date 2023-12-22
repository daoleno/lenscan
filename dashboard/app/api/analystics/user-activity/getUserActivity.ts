import { duckdb, parquetPath } from "@/lib/duckdb"

import "server-only"

import { DateRangeKey, getDateRangeAndCondition } from "../utils"

export type UserActivity = {
  day: string
  posts: number
  comments: number
  mirrors: number
  upvotes: number
  downvotes: number
}

export async function getUserActivity(
  rangeKey: DateRangeKey = "ALL",
  profileId: string | null = null
) {
  const groupByProfile = profileId ? ", profile_id" : ""
  const groupByActionedProfile = profileId ? ", actioned_by_profile_id" : ""

  let sql = `
        WITH publication_stats AS (
            SELECT
                DATE_TRUNC('day', block_timestamp)::date AS day,
                COUNT(*) FILTER (WHERE publication_type = 'POST') AS posts,
                COUNT(*) FILTER (WHERE publication_type = 'COMMENT') AS comments,
                COUNT(*) FILTER (WHERE publication_type = 'MIRROR') AS mirrors
                ${profileId ? ", profile_id" : ""}
            FROM
                read_parquet('${parquetPath}/publication_record/*.parquet')
            WHERE
                1 = 1
                ${profileId ? `AND profile_id = '${profileId}'` : ""}
                ${getDateRangeAndCondition(rangeKey, "block_timestamp")}
            GROUP BY day${groupByProfile}
        ), reaction_stats AS (
            SELECT
                DATE_TRUNC('day', action_at)::date AS day,
                COUNT(*) FILTER (WHERE type = 'UPVOTE') AS upvotes,
                COUNT(*) FILTER (WHERE type = 'DOWNVOTE') AS downvotes
                ${profileId ? ", actioned_by_profile_id" : ""}
            FROM
                read_parquet('${parquetPath}/publication_reaction/*.parquet')
            WHERE
                1 = 1
                ${
                  profileId ? `AND actioned_by_profile_id = '${profileId}'` : ""
                }
                ${getDateRangeAndCondition(rangeKey, "action_at")}
            GROUP BY day${groupByActionedProfile}
        )
        SELECT
            COALESCE(p.day, r.day) AS day,
            COALESCE(p.posts, 0) AS posts,
            COALESCE(p.comments, 0) AS comments,
            COALESCE(p.mirrors, 0) AS mirrors,
            COALESCE(r.upvotes, 0) AS upvotes,
            COALESCE(r.downvotes, 0) AS downvotes
        FROM
            publication_stats p
        FULL OUTER JOIN
            reaction_stats r ON p.day = r.day ${
              profileId ? "AND p.profile_id = r.actioned_by_profile_id" : ""
            }
        ORDER BY
            day;
      `

  console.log("--->", sql)

  const activities = await duckdb.all(sql)

  activities.forEach((a) => {
    a.posts = Number(a.posts)
    a.comments = Number(a.comments)
    a.mirrors = Number(a.mirrors)
    a.upvotes = Number(a.upvotes)
    a.downvotes = Number(a.downvotes)
    a.day = new Date(a.day).toLocaleDateString()
  })

  return activities as UserActivity[]
}
