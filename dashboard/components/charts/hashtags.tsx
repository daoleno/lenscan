"use client"

import { useState } from "react"
import { BarChart } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

export default function UserActivity() {
  const [range, setRange] = useState("ALL")
  const { data, error } = useSWR(
    `/api/analystics/popular-hashtags?range=${range}`,
    fetcher
  )

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
