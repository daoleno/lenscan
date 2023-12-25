"use client"

import { useState } from "react"
import { BarList } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

interface TopAppsProps {
  className?: string
}

export default function TopApps({ className }: TopAppsProps) {
  const [range, setRange] = useState("ALL")
  const queryString = `/api/analystics/topapps?range=${range}`
  const { data, error } = useSWR(queryString, fetcher)

  return (
    <ChartCard
      chartTitle={`Top ${data?.length} Apps`}
      range={range}
      setRange={setRange}
      className={className}
    >
      <div className="mb-3 flex flex-row justify-between text-sm font-semibold text-muted-foreground">
        <span>Name</span>
        <span>Publications</span>
      </div>
      <BarList data={data} showAnimation />
    </ChartCard>
  )
}
