import { AreaChart, SparkAreaChart } from "@tremor/react"

import { cn } from "@/lib/utils"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
} from "@/components/ui/card"
import {
  getDailyNewProfileStats,
  getProfilesGrowthPercentage,
} from "@/app/api/analystics/profiles/getNewProfileStats"

export default async function DailyProfileStats({ enableSpark = true }) {
  const dateRange = "1M"
  const chartdata = await getDailyNewProfileStats(dateRange)
  const totalCount = chartdata.reduce((acc, cur) => {
    return acc + cur.newProfiles
  }, 0)
  const growth = await getProfilesGrowthPercentage(dateRange)
  return (
    <Card className={cn(enableSpark && "px-4")}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <span>New Profiles</span>
        <span className="px-3 text-xs text-muted-foreground">{dateRange}</span>
      </CardHeader>
      {!enableSpark && (
        <CardDescription className="mx-7">
          <div className="text-2xl font-bold">{totalCount}</div>
          <p className="text-xs text-muted-foreground">
            {growth}% from last month
          </p>
        </CardDescription>
      )}
      <CardContent className="flex items-center justify-between gap-7">
        {enableSpark && (
          <div>
            <div className="text-2xl font-bold">{totalCount}</div>
            <p className="text-xs text-muted-foreground">
              {growth}% from last month
            </p>
          </div>
        )}
        {enableSpark ? (
          <SparkAreaChart
            data={chartdata}
            categories={["newProfiles"]}
            index={"day"}
            className="h-10 w-36"
          />
        ) : (
          <AreaChart
            data={chartdata}
            categories={["newProfiles"]}
            index={"day"}
          />
        )}
      </CardContent>
    </Card>
  )
}
