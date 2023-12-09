"use client"

import { useState } from "react"
import { LineChart } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

export default function UserActivity() {
  const [range, setRange] = useState("ALL")
  const { data, error } = useSWR(
    `/api/analystics/user-activities?range=${range}`,
    fetcher
  )

  return (
    <ChartCard chartTitle="Users Activity" range={range} setRange={setRange}>
      <LineChart
        data={data}
        index="date"
        categories={["count"]}
        showAnimation
      />
    </ChartCard>
  )
}
