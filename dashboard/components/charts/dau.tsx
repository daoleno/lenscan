"use client"

import { useState } from "react"
import { BarChart } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

interface UserActivityProps {
  profileId?: string | null
  className?: string
}

export default function DAU({ className }: UserActivityProps) {
  const [range, setRange] = useState("ALL")
  const queryString = `/api/analystics/active-users?range=${range}&statType=DAU`
  const { data, error } = useSWR(queryString, fetcher)

  console.log("data", data)

  return (
    <ChartCard
      chartTitle="Daily Active Users"
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
