import getRevenueRecords, {
	RevenueRecord,
} from "@/app/api/analystics/revenue/getProfileRevenue";
import { searchProfileRevenueParamsSchema } from "@/lib/validations/params";
import ProfileRevenueTable from "./profile-revenue-table ";

export interface RevenueProps {
	title?: string;
	searchParams: {
		[key: string]: string | string[] | undefined;
	};
	showToolbar?: boolean;
	showPagination?: boolean;
	showTitle?: boolean;
}
export default async function Revenue({
	title = "Revenue",
	searchParams,
	showToolbar,
	showPagination,
	showTitle = true,
}: RevenueProps) {
	const { page, per_page, sort, profile_id, currency_symbol } =
		searchProfileRevenueParamsSchema.parse(searchParams);
	const pageAsNumber = Number(page);
	const fallbackPage =
		isNaN(pageAsNumber) || pageAsNumber < 1 ? 1 : pageAsNumber;
	const perPageAsNumber = Number(per_page);
	const limit = isNaN(perPageAsNumber) ? 10 : perPageAsNumber;
	const offset = fallbackPage > 0 ? (fallbackPage - 1) * limit : 0;
	const [column, order] = (sort?.split(".") as [
		keyof RevenueRecord | undefined,
		"asc" | "desc" | undefined,
	]) ?? ["block_timestamp", "desc"];
	const symbols = currency_symbol?.split(".") ?? [];
	const profileIds = profile_id?.split(".") ?? [];
	const { revenueRecords, totalCount } = await getRevenueRecords({
		limit,
		offset,
		sort: {
			column,
			order,
		},
		filter: {
			profile_id: profileIds,
			symbol: symbols,
		},
	});
	const maxCount = totalCount ?? 1000000;
	const pageCount = Math.ceil(Number(maxCount) / limit);
	return (
		<ProfileRevenueTable
			title={title}
			data={revenueRecords}
			pageCount={pageCount}
			totalCount={maxCount}
			showToolbar={showToolbar}
			showPagination={showPagination}
			showTitle={showTitle}
		/>
	);
}
