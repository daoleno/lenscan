"use client"

import { useState } from "react"
import { BarList } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { Error } from "../error"
import { Loader } from "../loader"
import { ChartCard } from "./chart-card"

interface TopAppsProps {
  className?: string
}

export default function TopApps({ className }: TopAppsProps) {
  const [range, setRange] = useState("1D")
  const queryString = `/api/analystics/topapps?range=${range}`
  const { data: rawData, error } = useSWR(queryString, fetcher)
  if (error) return <Error msg={error.message} />
  if (!rawData) return <Loader fixed={false} />

  // only get name, value from data
  const data: any = rawData.map((item: any) => {
    return {
      name: item.name,
      value: item.value,
    }
  })

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
