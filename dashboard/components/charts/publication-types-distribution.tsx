import { DonutChart } from "@tremor/react"

import { getPublicationTypesDistribution } from "@/app/api/analystics/getPublicationTypesDistribution"

import { ChartCard } from "./chart-card"

export const revalidate = 60 * 60 * 5

export default async function PublicationTypesDistribution() {
  const data = await getPublicationTypesDistribution()

  return (
    <ChartCard chartTitle="Publication Types Distribution">
      <DonutChart
        className="mt-24"
        data={data}
        index="publication_type"
        category="count"
        showAnimation
      />
    </ChartCard>
  )
}
