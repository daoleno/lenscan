import { BarChart } from "@tremor/react"

import { getTopContributors } from "@/app/api/analystics/getTopContributors"

import { ChartCard } from "./chart-card"

export default async function TopContributors() {
  const data = await getTopContributors()

  return (
    <ChartCard chartTitle="Top Contributors">
      <BarChart
        data={data}
        index="profile_id"
        categories={[
          "publication_count",
          "reaction_count",
          "mention_count",
          "total_count",
        ]}
        stack
        showAnimation
      />
    </ChartCard>
  )
}
