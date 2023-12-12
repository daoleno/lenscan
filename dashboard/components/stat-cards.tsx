import { Copy, Loader2, MessageSquare, Scroll, Users } from "lucide-react"

import { formatNumber } from "@/lib/utils"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { getGlobalStats } from "@/app/api/analystics/getGlobalStats"

type Categories = {
  title: string
  metric?: string
  icon: any
}[]

export default async function StatCards() {
  const globalStats = await getGlobalStats()

  // console.log("stats", res)

  const categories: Categories = [
    {
      title: "Profiles",
      metric: globalStats.totalProfiles.toString(),
      icon: <Users />,
    },
    {
      title: "Posts",
      metric: globalStats.totalPosts.toString(),
      icon: <Scroll />,
    },
    {
      title: "Mirrors",
      metric: globalStats.totalMirrors.toString(),
      icon: <Copy />,
    },
    {
      title: "Comments",
      metric: globalStats.totalComments.toString(),
      icon: <MessageSquare />,
    },
  ]

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {categories.map((category) => (
        <Card key={category.title}>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              {category.title}
            </CardTitle>
            <div className="h-4 w-4 text-muted-foreground">{category.icon}</div>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {category.metric ? (
                formatNumber(category.metric)
              ) : (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              )}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
