import { DonutChart } from "@tremor/react"

import { getContentLanguageDistribution } from "@/app/api/analystics/getContentLanguageDistribution"

import { ChartCard } from "./chart-card"

export default async function ContentLanguageDistribution() {
  const data = await getContentLanguageDistribution()

  return (
    <ChartCard chartTitle="Content Language Distribution">
      <DonutChart
        data={data}
        index="language"
        category="frequency"
        variant="pie"
        showAnimation
      />
    </ChartCard>
  )
}
