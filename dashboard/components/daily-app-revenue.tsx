import { BarChart } from "@tremor/react";

import { getDailyAppRevenueStats } from "@/app/api/analystics/revenue/getAppRevenue";
import { Card, CardContent, CardHeader } from "@/components/ui/card";

export default async function DailyAppRevenue() {
	const chartdata = await getDailyAppRevenueStats();
	return (
		<Card>
			<CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
				<span>
					Daily App Revenue
					<span className="text-sm text-muted-foreground font-normal">
						{" "}
						(MATIC){" "}
					</span>
				</span>
			</CardHeader>
			<CardContent className="flex items-center justify-between gap-7">
				<BarChart
					data={chartdata.stats}
					categories={chartdata.apps}
					index={"day"}
					showAnimation
					stack
				/>
			</CardContent>
		</Card>
	);
}
