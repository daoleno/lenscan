import { BarChart } from "@tremor/react"

import { getTopProfiles } from "@/app/api/analystics/getTopProfiles"

import { ChartCard } from "./chart-card"

export default async function TopContributors() {
  const data = await getTopProfiles()

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
