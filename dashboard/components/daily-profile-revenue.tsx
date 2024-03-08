import { AreaChart } from "@tremor/react"

import { getTotalDailyRevenueBySymbol } from "@/app/api/analystics/revenue/getProfileRevenue"
import { Card, CardContent, CardHeader } from "@/components/ui/card"

export default async function DailyProfileRevenue({
	symbol,
}: { symbol: string }) {
	const chartdata = await getTotalDailyRevenueBySymbol(symbol)
	return (
		<Card>
			<CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2 font-bold">
				Daily Profile Revenue
			</CardHeader>
			<CardContent className="flex items-center justify-between gap-7">
				<AreaChart
					data={chartdata}
					categories={["amount"]}
					index={"day"}
					showAnimation
					stack
				/>
			</CardContent>
		</Card>
	)
}
