import { Suspense } from "react"

import { RevenueRecord } from "@/app/api/analystics/revenue/getProfileRevenue"
import { profileRevenueColumns } from "./data-table/columns"
import { DataTable } from "./data-table/data-table"
import { DataTableLoading } from "./data-table/data-table-loading"

interface ProfileRevenueTableProps {
	title: string
	data: RevenueRecord[]
	pageCount: number
	totalCount?: number
	showToolbar?: boolean
	showPagination?: boolean
	showTitle?: boolean
}

export default async function ProfileRevenueTable({
	title = "Revenue",
	data,
	pageCount,
	totalCount,
	showToolbar,
	showPagination,
	showTitle = true,
}: ProfileRevenueTableProps) {
	return (
		<div className="my-3 flex-1 flex-col gap-3 md:flex">
			{showTitle && (
				<h2 className="p-3 text-3xl font-bold tracking-tight">{title}</h2>
			)}
			<Suspense fallback={<DataTableLoading columnCount={6} rowCount={10} />}>
				<DataTable
					columns={profileRevenueColumns}
					data={data}
					pageCount={pageCount}
					totalCount={totalCount}
					showToolbar={showToolbar}
					showPagination={showPagination}
					// Render dynamic faceted filters
					filterableColumns={[
						{
							id: "currency_symbol",
							title: "Symbol",
							options: [
								{ label: "WMATIC", value: "WMATIC" },
								{ label: "BONSAI", value: "BONSAI" },
								{ label: "WETH", value: "WETH" },
								{ label: "USDC", value: "USDC" },
								{ label: "USDT", value: "USDT" },
								{ label: "DAI", value: "DAI" },
								{ label: "Pointless", value: "pointless" },
							],
						},
					]}
				/>
			</Suspense>
		</div>
	)
}
