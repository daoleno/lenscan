import ProfileRevenue from "@/components/profile-revenue"
import { RevalidateTime } from "@/config/ssr"
import { cn } from "@/lib/utils"
import Link from "next/link"

export const revalidate = RevalidateTime

interface PageProps {
	searchParams: {
		[key: string]: string | string[] | undefined
	}
}

const tokens = [
	{ label: "Wrapped Matic", value: "WMATIC" },
	{ label: "Bonsai Token", value: "BONSAI" },
	{ label: "Pointless", value: "pointless" },
	{ label: "USD Coin", value: "USDC" },
	{ label: "Dai Stablecoin", value: "DAI" },
	{ label: "Wrapped Ether", value: "WETH" },
]

export default async function Page({ searchParams }: PageProps) {
	const { tab } = searchParams
	const searchedTab = tab ? tab : "WMATIC"

	const linkBaseClass =
		"inline-flex items-center justify-center whitespace-nowrap rounded-sm px-3 py-1.5 text-sm font-medium ring-offset-background transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"

	const getLinkClassName = (currentTab: string) =>
		cn(linkBaseClass, {
			"bg-background text-foreground shadow-sm": searchedTab === currentTab,
		})

	return (
		<div className="flex flex-col">
			<h1 className="text-3xl font-bold text-foreground m-7">
				Profile Revenue
			</h1>
			<div className="inline-flex h-10 mr-auto mx-7 rounded-md bg-muted p-1 text-muted-foreground">
				{tokens.map((token) => (
					<Link
						key={token.value}
						className={getLinkClassName(token.value)}
						href={`/revenue/profile?tab=${token.value}`}
					>
						{token.label}
					</Link>
				))}
			</div>

			<ProfileRevenue
				symbol={searchedTab as string}
				searchParams={searchParams}
			/>
		</div>
	)
}
