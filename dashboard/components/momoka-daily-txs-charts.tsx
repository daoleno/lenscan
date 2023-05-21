import { trpc } from "@/lib/trpc";
import { BarChart, Card, Text, Title } from "@tremor/react";
import dayjs from "dayjs";
import { useState } from "react";
import { Loading } from "./loading";

export type StatisticsTimeQueryType = {
  timeStart: number;
  timeEnd: number;
};

const defaultTimeQuery = {
  timeStart: dayjs().subtract(1, "year").unix(),
  timeEnd: dayjs().unix(),
};

const valueFormatter = (number: number) =>
  `${Intl.NumberFormat("us").format(number).toString()}`;

export default function MomokaDailyTxsChart() {
  const [timeRange, setTimeRange] =
    useState<StatisticsTimeQueryType>(defaultTimeQuery);

  const { isFetching: dailyTxCountFetching, data: dailyTxCount } =
    trpc.momoka.getDailyTxCount.useQuery(timeRange, {
      enabled: !!timeRange,
    });

  if (!dailyTxCount) {
    return <Loading fixed={false} />;
  }

  return (
    <Card>
      <Title>Daily Momoka Transactions</Title>
      <Text>Comparison of daily transactions, posts, comments and mirrors</Text>
      <BarChart
        className="mt-4 h-80"
        data={dailyTxCount}
        categories={["posts", "mirrors", "comments"]}
        index="date"
        colors={["indigo", "sky", "teal"]}
        // curveType="natural"
        // startEndOnly
        valueFormatter={valueFormatter}
        stack
      />
    </Card>
  );
}
