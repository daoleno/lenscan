import React from "react"
import { SparkBarChart } from "@tremor/react"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

interface RevenueData {
  revenue: number
}

interface MyComponentProps {
  data?: RevenueData[]
  mode?: string
  theme?: {
    cssVars: {
      [key: string]: any
    }
  }
}

export const chartdata = [
  {
    month: "Jan 21",
    Performance: 4000,
    Benchmark: 3000,
  },
  {
    month: "Feb 21",
    Performance: 3000,
    Benchmark: 2000,
  },
  {
    month: "Mar 21",
    Performance: 2000,
    Benchmark: 1700,
  },
  {
    month: "Apr 21",
    Performance: 2780,
    Benchmark: 2500,
  },
  {
    month: "May 21",
    Performance: 1890,
    Benchmark: 1890,
  },
  {
    month: "Jun 21",
    Performance: 2390,
    Benchmark: 2000,
  },
  {
    month: "Jul 21",
    Performance: 3490,
    Benchmark: 3000,
  },
]

const StatChartCard: React.FC<MyComponentProps> = ({ data, mode, theme }) => (
  <Card>
    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
      <CardTitle className="text-base font-normal">Total Revenue</CardTitle>
    </CardHeader>
    <CardContent className="flex gap-7">
      <div>
        <div className="text-2xl font-bold">$15,231.89</div>
        <p className="text-xs text-muted-foreground">+20.1% from last month</p>
      </div>
      <SparkBarChart
        data={chartdata}
        categories={["Performance"]}
        index={"month"}
        colors={["emerald"]}
        className="h-10 w-36"
      />
    </CardContent>
  </Card>
)

export default StatChartCard
