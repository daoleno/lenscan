import TopEarningProfiles from "@/components/charts/top-earning-profiles"
import DailyProfileRevenue from "@/components/daily-profile-revenue"
import Revenue from "@/components/revenue"
import TokenRevenueCard from "@/components/token-revenue-card"

export default function ProfileRevenue({
	symbol,
	searchParams,
}: { symbol: string; searchParams: any }) {
	return (
		<div className="flex flex-col gap-7 p-8">
			<TokenRevenueCard symbol={symbol} />
			<DailyProfileRevenue symbol={symbol} />
			<TopEarningProfiles symbol={symbol} />
			<Revenue
				searchParams={searchParams}
				showToolbar
				showPagination
				showTitle={false}
			/>
		</div>
	)
}
