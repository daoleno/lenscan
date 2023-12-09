"use client"

import { useState } from "react"
import { DonutChart } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

export default function PublicationTypesDistribution() {
  const [range, setRange] = useState("ALL")
  const { data, error } = useSWR(
    `/api/analystics/publication-types-distribution?range=${range}`,
    fetcher
  )

  return (
    <ChartCard
      chartTitle="Publication Types Distribution"
      range={range}
      setRange={setRange}
    >
      <DonutChart
        data={data}
        index="publication_type"
        category="count"
        showAnimation
      />
    </ChartCard>
  )
}
