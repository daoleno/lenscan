import { searchParamsSchema } from "@/lib/validations/params";

import AppRevenueCards from "@/components/app-revenue-cards";
import DailyAppRevenue from "@/components/daily-app-revenue";
import RevenueTable from "@/components/revenue-table";
import { RevalidateTime } from "@/config/ssr";
import { Publication } from "../api/publications/getPublications";
import getRevenue from "../api/revenue/getRevenue";

export const revalidate = RevalidateTime;

interface PageProps {
	searchParams: {
		[key: string]: string | string[] | undefined;
	};
}

export default async function Page({ searchParams }: PageProps) {
	// Parse search params using zod schema
	const { page, per_page, sort, app } = searchParamsSchema.parse(searchParams);

	// Fallback page for invalid page numbers
	const pageAsNumber = Number(page);
	const fallbackPage =
		isNaN(pageAsNumber) || pageAsNumber < 1 ? 1 : pageAsNumber;
	// Number of items per page
	const perPageAsNumber = Number(per_page);
	const limit = isNaN(perPageAsNumber) ? 10 : perPageAsNumber;
	// Number of items to skip
	const offset = fallbackPage > 0 ? (fallbackPage - 1) * limit : 0;
	// Column and order to sort by
	// Spliting the sort string by "." to get the column and order
	// Example: "title.desc" => ["title", "desc"]
	const [column, order] = (sort?.split(".") as [
		keyof Publication | undefined,
		"asc" | "desc" | undefined,
	]) ?? ["block_timestamp", "desc"];
	const apps = app?.split(".") ?? [];
	const { revenue, totalCount } = await getRevenue({
		limit,
		offset,
		sort: {
			column,
			order,
		},
		filter: {
			app: apps,
		},
	});
	const pageCount = Math.ceil(Number(totalCount) / limit);

	return (
		<div className="flex flex-col gap-7 p-8">
			<h2 className="text-3xl font-bold tracking-tight">App Revenue</h2>
			<AppRevenueCards />
			<DailyAppRevenue />
			<RevenueTable
				data={revenue}
				pageCount={pageCount}
				totalCount={totalCount}
				showToolbar={true}
				showPagination={true}
			/>
		</div>
	);
}
