import { SparkAreaChart } from "@tremor/react"

import { Card, CardContent, CardHeader } from "@/components/ui/card"
import {
  getDauGrowthPercentages,
  getDauStats,
} from "@/app/api/analystics/dau/getDailyActiveUser"

export default async function DauStats() {
  const dateRange = "1M"
  const chartdata = await getDauStats(dateRange)
  const totalCount = chartdata.reduce((acc, cur) => {
    return acc + cur.total
  }, 0)
  const growth = await getDauGrowthPercentages(dateRange)
  return (
    <Card className="px-4">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <span>Active Users</span>
        <span className="px-3 text-xs text-muted-foreground">{dateRange}</span>
      </CardHeader>
      <CardContent className="flex items-center justify-between gap-7">
        <div>
          <div className="text-2xl font-bold">{totalCount}</div>
          <p className="text-xs text-muted-foreground">
            {growth}% from last month
          </p>
        </div>
        <div>
          <SparkAreaChart
            data={chartdata}
            categories={["total"]}
            index={"day"}
            className="h-10 w-36"
          />
        </div>
      </CardContent>
    </Card>
  )
}
