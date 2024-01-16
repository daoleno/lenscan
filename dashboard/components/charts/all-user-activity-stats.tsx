"use client"

import { useState } from "react"
import { BarChart } from "@tremor/react"

import { ChartCard } from "./chart-card"

interface UserActivityProps {
  className?: string
  allStats: any
}

export default function AllUserActivityStats({
  className,
  allStats,
}: UserActivityProps) {
  const [range, setRange] = useState("ALL")
  const data = allStats[range]

  return (
    <ChartCard
      chartTitle="Users Activity"
      range={range}
      setRange={setRange}
      className={className}
    >
      <BarChart
        data={data}
        index="day"
        categories={["posts", "comments", "mirrors", "upvotes", "downvotes"]}
        // showAnimation
        showGridLines={false}
        stack
      />
    </ChartCard>
  )
}
