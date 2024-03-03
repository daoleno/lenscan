import { getTotalAppRevenue } from "@/app/api/analystics/revenue/getAppRevenue";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

type Categories = {
	title: string;
	metric: string;
	// icon: any;
}[];

export default async function AppRevenueCards() {
	const revenue = await getTotalAppRevenue();

	const categories: Categories = revenue.map((revenue) => ({
		title: revenue.app,
		metric: revenue.revenue.toLocaleString("en-US"),
	}));

	return (
		<div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
			{categories.map((category) => (
				<Card key={category.title}>
					<CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
						<CardTitle className="text-sm font-medium">
							{category.title.charAt(0).toUpperCase() + category.title.slice(1)}
						</CardTitle>
						{/* <div className="h-4 w-4 text-muted-foreground">{category.icon}</div> */}
					</CardHeader>
					<CardContent>
						<div className="text-2xl font-bold">
							{category.metric}
							<span className="text-sm font-normal text-muted-foreground">
								{" "}
								MATIC{" "}
							</span>
						</div>
					</CardContent>
				</Card>
			))}
		</div>
	);
}
