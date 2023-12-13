import { SparkAreaChart } from "@tremor/react"

import { Card, CardContent, CardHeader } from "@/components/ui/card"
import {
  getDailyNewProfileStats,
  getProfilesGrowthPercentage,
} from "@/app/api/analystics/profiles/getNewProfileStats"

export default async function DailyProfileStats() {
  const dateRange = "1M"
  const chartdata = await getDailyNewProfileStats(dateRange)
  const totalCount = chartdata.reduce((acc, cur) => {
    return acc + cur.newProfiles
  }, 0)
  const growth = await getProfilesGrowthPercentage(dateRange)
  return (
    <Card className="px-4">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <span>New Profiles</span>
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
            categories={["newProfiles"]}
            index={"day"}
            className="h-10 w-36"
          />
        </div>
      </CardContent>
    </Card>
  )
}
