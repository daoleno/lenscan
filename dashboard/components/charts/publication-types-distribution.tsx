import { DonutChart } from "@tremor/react"

import { RevalidateTime } from "@/config/ssr"
import { getPublicationTypesDistribution } from "@/app/api/analystics/getPublicationTypesDistribution"

import { ChartCard } from "./chart-card"

export const revalidate = RevalidateTime

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
