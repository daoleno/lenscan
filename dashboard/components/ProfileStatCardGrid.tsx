import { Card, Color, Flex, Grid, Icon, Metric, Text } from "@tremor/react";

import {
  ChatBubbleLeftEllipsisIcon,
  ChatBubbleLeftRightIcon,
  ClipboardDocumentIcon,
  RectangleStackIcon,
  UserPlusIcon,
} from "@heroicons/react/24/solid";
import { useProfile } from "@lens-protocol/react-web";

type Categories = {
  title: string;
  metric: string;
  icon: any;
  color: Color;
}[];

export default function ProfileStatCardGrid({
  profileId,
}: {
  profileId: string;
}) {
  const { data: profile, loading, error } = useProfile({ profileId });
  console.log("profile", profile);

  const categories: Categories = [
    {
      title: "Followers",
      metric: profile?.stats.totalFollowers.toString() || "-",
      icon: UserPlusIcon,
      color: "yellow",
    },
    {
      title: "Following",
      metric: profile?.stats.totalFollowing.toString() || "-",
      icon: UserPlusIcon,
      color: "yellow",
    },
    {
      title: "Publications",
      metric: profile?.stats.totalPublications.toString() || "-",
      icon: ChatBubbleLeftEllipsisIcon,
      color: "green",
    },
    {
      title: "Posts",
      metric: profile?.stats.totalPosts.toString() || "-",
      icon: ChatBubbleLeftEllipsisIcon,
      color: "green",
    },
    {
      title: "Collects",
      metric: profile?.stats.totalCollects.toString() || "-",
      icon: RectangleStackIcon,
      color: "purple",
    },
    {
      title: "Mirrors",
      metric: profile?.stats.totalMirrors.toString() || "-",
      icon: ClipboardDocumentIcon,
      color: "red",
    },
    {
      title: "Comments",
      metric: profile?.stats.totalComments.toString() || "-",
      icon: ChatBubbleLeftRightIcon,
      color: "orange",
    },
  ];

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
