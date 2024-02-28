import { Copy, MessageSquare, Scroll, Users } from "lucide-react"

import { getGlobalStats } from "@/app/api/analystics/getGlobalStats"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

type Categories = {
  title: string
  metric: string
  icon: any
}[]

export default async function StatCards() {
  const globalStats = await getGlobalStats()

  const categories: Categories = [
    {
      title: "Profiles",
      metric: globalStats.totalProfiles,
      icon: <Users />,
    },
    {
      title: "Posts",
      metric: globalStats.totalPosts,
      icon: <Scroll />,
    },
    {
      title: "Mirrors",
      metric: globalStats.totalMirrors,
      icon: <Copy />,
    },
    {
      title: "Comments",
      metric: globalStats.totalComments,
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
                {category.metric}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
