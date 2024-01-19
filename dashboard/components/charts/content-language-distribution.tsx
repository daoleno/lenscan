import { DonutChart } from "@tremor/react"

import { RevalidateTime } from "@/config/ssr"
import { getContentLanguageDistribution } from "@/app/api/analystics/getContentLanguageDistribution"

import { ChartCard } from "./chart-card"

export const revalidate = RevalidateTime

export default async function ContentLanguageDistribution() {
  const data = await getContentLanguageDistribution()

  return (
    <ChartCard chartTitle="Content Language Distribution">
      <DonutChart
        className="mt-24"
        data={data}
        index="language"
        category="frequency"
        variant="pie"
        showAnimation
      />
    </ChartCard>
  )
}
