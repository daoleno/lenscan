"use client"

import { useState } from "react"
import { BarChart } from "@tremor/react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { ChartCard } from "./chart-card"

interface UserActivityProps {
  profileId?: string | null
}

export default function UserActivity({ profileId = null }: UserActivityProps) {
  const [range, setRange] = useState("ALL")
  const queryString = profileId
    ? `/api/analystics/user-activity?range=${range}&profile_id=${profileId}`
    : `/api/analystics/user-activity?range=${range}`
  const { data, error } = useSWR(queryString, fetcher)

  return (
    <ChartCard
      chartTitle={profileId ? "Activity" : "Users Activity"}
      range={range}
      setRange={setRange}
    >
      <BarChart
        data={data}
        index="day"
        categories={["posts", "comments", "mirrors", "upvotes", "downvotes"]}
        showAnimation
        showGridLines={false}
        stack
      />
    </ChartCard>
  )
}
