"use client"

import { useState } from "react"
import { AreaChart } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

export default function PublicationReactions() {
  const [range, setRange] = useState("ALL")
  const { data, error } = useSWR(
    `/api/analystics/publication-reactions?range=${range}`,
    fetcher
  )

  return (
    <ChartCard
      chartTitle="Publication Reactions"
      range={range}
      setRange={setRange}
    >
      <AreaChart
        data={data}
        index="date"
        categories={[
          "UPVOTE",
          "DOWNVOTE",
          // ... other reaction types
        ]}
        showAnimation
        stack
      />
    </ChartCard>
  )
}
