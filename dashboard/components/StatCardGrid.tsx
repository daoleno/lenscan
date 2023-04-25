import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

import lensClient from "@/lib/lensclient";
import { formatNumber } from "@/lib/utils";
import {
  Copy,
  Flame,
  Layers,
  Loader2,
  MessageSquare,
  Scroll,
  UserPlus,
  Users,
} from "lucide-react";
import useSWR from "swr";

type Categories = {
  title: string;
  metric?: string;
  icon: any;
}[];

export default function StatCardGrid() {
  const { data: stats, error } = useSWR(
    "lens-protocol-global-stats",
    async () => {
      return await lensClient.stats.fetch();
    }
  );
  const categories: Categories = [
    {
      title: "Profiles",
      metric: stats?.totalProfiles.toString(),
      icon: <Users />,
    },
    {
      title: "Posts",
      metric: stats?.totalPosts.toString(),
      icon: <Scroll />,
    },
    {
      title: "Follows",
      metric: stats?.totalFollows.toString(),
      icon: <UserPlus />,
    },
    {
      title: "Collects",
      metric: stats?.totalCollects.toString(),
      icon: <Layers />,
    },
    {
      title: "Mirrors",
      metric: stats?.totalMirrors.toString(),
      icon: <Copy />,
    },
    {
      title: "Comments",
      metric: stats?.totalComments.toString(),
      icon: <MessageSquare />,
    },
    {
      title: "Burnt Profiles",
      metric: stats?.totalBurntProfiles.toString(),
      icon: <Flame />,
    },
  ];

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
  );
}
