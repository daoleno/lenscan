import { Card, Color, Flex, Grid, Icon, Metric, Text } from "@tremor/react";

import lensClient from "@/lib/lensclient";
import {
  ChatBubbleLeftEllipsisIcon,
  ChatBubbleLeftRightIcon,
  ClipboardDocumentIcon,
  RectangleStackIcon,
  UserGroupIcon,
  UserPlusIcon,
} from "@heroicons/react/24/solid";
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
        icon: UserGroupIcon, // replace with your own icon component
        color: "blue", // replace with your desired color
      },
      {
        title: "Posts",
        metric: stats.totalPosts.toString(),
        icon: ChatBubbleLeftEllipsisIcon,
        color: "green",
      },
      {
        title: "Follows",
        metric: stats.totalFollows.toString(),
        icon: UserPlusIcon,
        color: "yellow",
      },
      {
        title: "Collects",
        metric: stats.totalCollects.toString(),
        icon: RectangleStackIcon,
        color: "purple",
      },
      {
        title: "Mirrors",
        metric: stats.totalMirrors.toString(),
        icon: ClipboardDocumentIcon,
        color: "red",
      },
      {
        title: "Comments",
        metric: stats.totalComments.toString(),
        icon: ChatBubbleLeftRightIcon,
        color: "orange",
      },
      // {
      //   title: "Burnt Profiles",
      //   metric: stats.totalBurntProfiles.toString(),
      //   icon: FireIcon,
      //   color: "teal",
      // },
    ];
    return categories as Categories;
  });

  useEffect(() => {
    if (data) {
      setCategories(data);
    }
  }, [data]);

  console.log(categories);

  return (
    <Grid numColsSm={2} numColsLg={3} className="mt-6 gap-6">
      {categories.map((item: any) => (
        <Card key={item.title} decoration="top" decorationColor={item.color}>
          <Flex justifyContent="start" className="space-x-4">
            <Icon
              icon={item.icon}
              variant="light"
              size="xl"
              color={item.color}
            />
            <div className="truncate">
              <Text>{item.title}</Text>
              <Metric className="truncate">{item.metric}</Metric>
            </div>
          </Flex>
        </Card>
      ))}
    </Grid>
  );
}
