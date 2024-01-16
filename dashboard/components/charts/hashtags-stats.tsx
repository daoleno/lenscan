"use client"

import { useState } from "react"
import { BarChart } from "@tremor/react"

import { ChartCard } from "./chart-card"

export default function HashTagsStats({ allStats }: any) {
  const [range, setRange] = useState("ALL")
  const data = allStats[range]

  return (
    <ChartCard chartTitle="Popular Hashtags" range={range} setRange={setRange}>
      <BarChart
        data={data}
        index="hashtag"
        categories={["count"]}
        showAnimation
      />
    </ChartCard>
  )
}
