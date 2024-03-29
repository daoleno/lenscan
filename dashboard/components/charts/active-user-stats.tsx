"use client"

import { BarChart } from "@tremor/react"
import { useState } from "react"

import { ChartCard } from "./chart-card"

interface ActiveUserStatsProps {
  className?: string
  title: string
  allStats: any
}

export default function ActiveUserStats({
  className,
  title,
  allStats,
}: ActiveUserStatsProps) {
  const [range, setRange] = useState("ALL")
  const data = allStats[range]

  return (
    <ChartCard
      chartTitle={title}
      range={range}
      setRange={setRange}
      className={className}
    >
      <BarChart
        data={data?.stats}
        index="time"
        categories={data?.apps}
        // showAnimation
        showGridLines={false}
        stack
      />
    </ChartCard>
  )
}
