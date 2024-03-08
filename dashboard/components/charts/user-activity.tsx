"use client";

import { BarChart } from "@tremor/react";
import { useState } from "react";
import useSWR from "swr";

import fetcher from "@/lib/fetcher";

import { Loader } from "../loader";
import { ChartCard } from "./chart-card";

interface UserActivityProps {
	profileId?: string | null;
	className?: string;
}

export default function UserActivity({
	profileId = null,
	className,
}: UserActivityProps) {
	const [range, setRange] = useState("ALL");
	const queryString = profileId
		? `/api/analystics/user-activity?range=${range}&profile_id=${profileId}`
		: `/api/analystics/user-activity?range=${range}`;
	const { data, error, isLoading } = useSWR(queryString, fetcher);

	return (
		<ChartCard
			chartTitle={profileId ? "Activity" : "Users Activity"}
			range={range}
			setRange={setRange}
			className={className}
		>
			{isLoading ? (
				<Loader fixed={false} />
			) : (
				<BarChart
					data={data}
					index="day"
					categories={["posts", "comments", "mirrors", "upvotes", "downvotes"]}
					// showAnimation
					showGridLines={false}
					stack
				/>
			)}
		</ChartCard>
	);
}
