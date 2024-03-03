import { Suspense } from "react";

import { Revenue } from "@/app/api/revenue/getRevenue";
import { revenueColumns } from "./data-table/columns";
import { DataTable } from "./data-table/data-table";
import { DataTableLoading } from "./data-table/data-table-loading";

interface RevenueTableProps {
	data: Revenue[];
	pageCount: number;
	totalCount?: number;
	showToolbar?: boolean;
	showPagination?: boolean;
}

export default async function RevenueTable({
	data,
	pageCount,
	totalCount,
	showToolbar,
	showPagination,
}: RevenueTableProps) {
	return (
		<div className="my-3 flex-1 flex-col gap-3 md:flex">
			<Suspense fallback={<DataTableLoading columnCount={6} rowCount={10} />}>
				<DataTable
					columns={revenueColumns}
					data={data}
					pageCount={pageCount}
					totalCount={totalCount}
					showToolbar={showToolbar}
					showPagination={showPagination}
					// Render dynamic faceted filters
					filterableColumns={[
						{
							id: "app",
							title: "App",
							options: [
								{ label: "Lens Protocol", value: "Lens Protocol" },
								{ label: "Hey", value: "Hey" },
								{ label: "Orb", value: "Orb" },
								{ label: "Tape", value: "Tape" },
								{ label: "Phaver", value: "Phaver" },
								{ label: "Buttrfly", value: "Buttrfly" },
							],
						},
					]}
				/>
			</Suspense>
		</div>
	);
}
