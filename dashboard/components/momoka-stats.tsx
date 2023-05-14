import { useEffect } from "react";

import useSubmitterSpent from "@/hooks/useSubmitterSpent";
import getCoingeckoPrice from "@/lib/getMaticPrice";
import weiToEth from "@/lib/weiToEth";
import {
  useDaSummaryLazyQuery,
  useDataAvailabilitySubmittersLazyQuery,
} from "@/src/generated";
import { useState } from "react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { formatNumber } from "@/lib/utils";
import { Loading } from "./loading";

import { Coins, CreditCard, Loader2, Sprout } from "lucide-react";

type Categories = {
  title: string;
  metric?: string;
  icon: any;
}[];

const MomokaStats = () => {
  const [allTransactionsCount, setAllTransactionsCount] = useState(0);
  const [topSubmitter, setTopSubmitter] = useState<any>(null);
  const [totalSpent, setTotalSpent] = useState<number | null>(null);
  const [maticMarketPrice, setMaticMarketPrice] = useState(0);
  const [fetchTopSubmitter, { loading: submittersDataLoading }] =
    useDataAvailabilitySubmittersLazyQuery({
      fetchPolicy: "no-cache",
    });
  const [fetchAllCount, { loading }] = useDaSummaryLazyQuery({
    fetchPolicy: "no-cache",
  });
  const { fetchData: fetchSpentAmount, loading: fetchingSpentAmount } =
    useSubmitterSpent();

  const fetchCounts = async () => {
    const { data: countData } = await fetchAllCount();
    const { data: submittersData } = await fetchTopSubmitter();
    setAllTransactionsCount(
      countData?.dataAvailabilitySummary.totalTransactions ?? 0
    );
    if (submittersData?.dataAvailabilitySubmitters?.items[0]) {
      const submitters = submittersData?.dataAvailabilitySubmitters.items.map(
        (el) => el.address
      );
      const sum = await fetchSpentAmount(submitters);
      setTotalSpent(sum);
      setTopSubmitter(submittersData?.dataAvailabilitySubmitters?.items[0]);
    }
  };

  const fetchMaticPrice = async () => {
    const price = await getCoingeckoPrice();
    setMaticMarketPrice(price);
  };

  useEffect(() => {
    fetchMaticPrice();
    fetchCounts();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (loading || submittersDataLoading) {
    return <Loading fixed={false} />;
  }

  const getTotalSpentInUsd = () => {
    return totalSpent ? weiToEth(totalSpent.toString()) * maticMarketPrice : 0;
  };

  const categories: Categories = [
    {
      title: "Transactions",
      metric: formatNumber(allTransactionsCount),
      icon: <Sprout />,
    },
    {
      title: "Total Spent",
      metric: `$ ${getTotalSpentInUsd().toFixed(2)}`,
      icon: <CreditCard />,
    },
    {
      title: "Average Fee",
      metric: `$ ${(getTotalSpentInUsd() / allTransactionsCount).toFixed(4)}`,
      icon: <Coins />,
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
                category.metric
              ) : (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              )}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};

export default MomokaStats;
