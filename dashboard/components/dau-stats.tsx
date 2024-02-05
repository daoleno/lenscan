import { SparkAreaChart } from "@tremor/react"

import { Card, CardContent, CardHeader } from "@/components/ui/card"
import {
  getDauGrowthPercentages,
  getNetworkUserStats,
} from "@/app/api/analystics/active-users/getActiveUserStats"

export default async function DauStats() {
  const chartdata = await getNetworkUserStats("1Y", "MAU")
  const curMonthCount = chartdata[chartdata.length - 1].total
  const growth = await getDauGrowthPercentages("1M")
  return (
    <Card className="px-4">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <span>Active Users</span>
        <span className="px-3 text-xs text-muted-foreground">{"1M"}</span>
      </CardHeader>
      <CardContent className="flex items-center justify-between gap-7">
        <div>
          <div className="text-2xl font-bold">{curMonthCount}</div>
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
