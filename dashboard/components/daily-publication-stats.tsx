import { SparkAreaChart } from "@tremor/react"

import { Card, CardContent, CardHeader } from "@/components/ui/card"
import {
  getDailyPublicationStats,
  getPublicationsGrowthPercentage,
} from "@/app/api/analystics/publications/getDailyPublicationStats"

export default async function DailyPublicationStats() {
  const dateRange = "1M"
  const chartdata = await getDailyPublicationStats(dateRange)
  const totalCount = chartdata.reduce((acc, cur) => {
    return acc + cur.total
  }, 0)
  const growth = await getPublicationsGrowthPercentage(dateRange)
  return (
    <Card className="px-4">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <span>New Publications</span>
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
