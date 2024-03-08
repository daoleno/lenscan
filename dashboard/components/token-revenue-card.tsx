import { getTotalProfileRevenueBySymbol } from "@/app/api/analystics/revenue/getProfileRevenue"
import { formatCryptoValue } from "@/lib/utils"
import Image from "next/image"
import Link from "next/link"

export default async function TokenRevenueCard({ symbol }: { symbol: string }) {
	const revenue = await getTotalProfileRevenueBySymbol(symbol)

	if (!revenue) {
		return null
	}

	return (
		<div className="flex gap-7 items-center p-3">
			<div className="flex gap-2 items-center font-bold text-lg">
				<Image
					className="text-muted-foreground"
					src={`/tokens/${revenue.token_symbol.toLowerCase()}.svg`}
					alt={revenue.token_symbol}
					width={32}
					height={32}
				/>
				{revenue.token_name}
			</div>

			{/* dot separator */}
			<div className="w-1 h-1 bg-muted-foreground rounded-full" />

			<div className="flex gap-2">
				<Link
					href={`https://polygonscan.com/token/${revenue.token_address}`}
					target="_blank"
					rel="noopener noreferrer"
				>
					<Image
						src="/apps/polygon.svg"
						alt="polygon"
						width={20}
						height={20}
						className="hover:opacity-80"
					/>
				</Link>
				<Link
					href={`https://app.uniswap.org/explore/tokens/polygon/${revenue.token_address}`}
					target="_blank"
					rel="noopener noreferrer"
				>
					<Image
						src="/apps/uniswap.svg"
						alt="uniswap"
						width={20}
						height={20}
						className="hover:opacity-80"
					/>
				</Link>
			</div>

			<div className="w-1 h-1 bg-muted-foreground rounded-full" />

			<div className="flex items-baseline gap-1">
				<span className="font-semibold text-md">
					{formatCryptoValue(revenue.total_revenue, revenue.token_decimals)}
				</span>
				<span className="text-xs font-normal text-muted-foreground">
					{revenue.token_symbol}
				</span>
			</div>
		</div>
	)
}
