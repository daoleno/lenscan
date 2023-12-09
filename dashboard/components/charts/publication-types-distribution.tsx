import { DonutChart } from "@tremor/react"

import { getPublicationTypesDistribution } from "@/app/api/analystics/getPublicationTypesDistribution"

import { ChartCard } from "./chart-card"

export default async function PublicationTypesDistribution() {
  const data = await getPublicationTypesDistribution()

  return (
    <ChartCard chartTitle="Publication Types Distribution">
      <DonutChart
        data={data}
        index="publication_type"
        category="count"
        showAnimation
      />
    </ChartCard>
  )
}
