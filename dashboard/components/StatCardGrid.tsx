import { Color } from "@tremor/react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

import lensClient from "@/lib/lensclient";
import { formatNumber } from "@/lib/utils";
import {
  Copy,
  Flame,
  Layers,
  MessageSquare,
  Scroll,
  UserPlus,
  Users,
} from "lucide-react";
import { useEffect, useState } from "react";
import useSWR from "swr";

type Categories = {
  title: string;
  metric: string;
  icon: any;
  color: Color;
}[];

export default function StatCardGrid() {
  const [categories, setCategories] = useState<Categories>([]);
  const { data, error } = useSWR("lens-protocol-global-stats", async () => {
    const stats = await lensClient.stats.fetch();
    const categories = [
      {
        title: "Profiles",
        metric: stats.totalProfiles.toString(),
        icon: <Users />, // replace with your own icon component
        color: "bg-blue-100", // replace with your desired color
      },
      {
        title: "Posts",
        metric: stats.totalPosts.toString(),
        icon: <Scroll />,
        color: "bg-green-100",
      },
      {
        title: "Follows",
        metric: stats.totalFollows.toString(),
        icon: <UserPlus />,
        color: "bg-yellow-100",
      },
      {
        title: "Collects",
        metric: stats.totalCollects.toString(),
        icon: <Layers />,
        color: "bg-purple-100",
      },
      {
        title: "Mirrors",
        metric: stats.totalMirrors.toString(),
        icon: <Copy />,
        color: "bg-red-100",
      },
      {
        title: "Comments",
        metric: stats.totalComments.toString(),
        icon: <MessageSquare />,
        color: "bg-orange-100",
      },
      {
        title: "Burnt Profiles",
        metric: stats.totalBurntProfiles.toString(),
        icon: <Flame />,
        color: "bg-teal-100",
      },
    ];
    return categories as Categories;
  });

  useEffect(() => {
    if (data) {
      setCategories(data);
    }
  }, [data]);

  return (
    <>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {categories.map((category) => (
          <Card key={category.title}>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                {category.title}
              </CardTitle>
              <div className="h-4 w-4 text-muted-foreground">
                {category.icon}
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">
                {formatNumber(category.metric)}
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </>
  );
}
