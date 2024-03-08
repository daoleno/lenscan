import Publications from "@/components/publications";
import Revenue from "@/components/revenue";
import { cn } from "@/lib/utils";
import Link from "next/link";

export default function ProfileRecord({ searchParams }: { searchParams: any }) {
	const { tab, profile_id } = searchParams;
	const searchedTab = tab ? tab : "publications";

	const linkBaseClass =
		"inline-flex items-center justify-center whitespace-nowrap rounded-sm px-3 py-1.5 text-sm font-medium ring-offset-background transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50";

	const getLinkClassName = (currentTab: string) =>
		cn(linkBaseClass, {
			"bg-background text-foreground shadow-sm": searchedTab === currentTab,
		});

	return (
		<>
			<div className="inline-flex h-10 items-center justify-center rounded-md bg-muted p-1 text-muted-foreground mt-7 mb-3">
				<Link
					className={getLinkClassName("publications")}
					href={`/profile/${profile_id}?tab=publications`}
				>
					Publications
				</Link>
				<Link
					className={getLinkClassName("revenue")}
					href={`/profile/${profile_id}?tab=revenue`}
				>
					Revenue
				</Link>
			</div>

			{searchedTab === "revenue" ? (
				<Revenue
					searchParams={searchParams}
					showToolbar
					showPagination
					showTitle={false}
				/>
			) : (
				<Publications
					searchParams={searchParams}
					showToolbar
					showPagination
					showTitle={false}
				/>
			)}
		</>
	);
}
