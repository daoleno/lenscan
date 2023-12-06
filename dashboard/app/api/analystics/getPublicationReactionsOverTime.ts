import duckdb from "@/lib/duckdb"

import "server-only"

export async function getPublicationReactionsOvertime() {
  const sql = `
    SELECT
      DATE_TRUNC('day', action_at) AS reaction_day,
      type AS reaction_type,
      COUNT(*) AS reaction_count
    FROM
      publication_reaction
    GROUP BY
      reaction_day, reaction_type
    ORDER BY
      reaction_day;
  `

  console.log(sql)
  const reactionsRaw = await duckdb.all(sql)

  // Convert bigint to number and format date to ISO string
  reactionsRaw.forEach((a) => {
    a.reaction_count = Number(a.reaction_count)
    a.reaction_day = new Date(a.reaction_day).toISOString().split("T")[0] // Format to 'YYYY-MM-DD'
  })

  // Transform data for the AreaChart
  const transformedData = reactionsRaw.reduce(
    (acc, { reaction_day, reaction_type, reaction_count }) => {
      let dayEntry = acc.find((entry: any) => entry.date === reaction_day)
      if (!dayEntry) {
        dayEntry = { date: reaction_day }
        acc.push(dayEntry)
      }
      dayEntry[reaction_type] = reaction_count
      return acc
    },
    []
  )

  return transformedData
}
