"use client"

import { BarList } from "@tremor/react"
import { useState } from "react"
import useSWR from "swr"

import fetcher from "@/lib/fetcher"

import { TopEarningProfile } from "@/app/api/analystics/top-earning-profiles/getTopEarningProfiles"
import { Error } from "../error"
import { Loader } from "../loader"
import { ChartCard } from "./chart-card"

interface TopEarningProfilesProps {
	symbol: string
	className?: string
}

export default function TopEarningProfiles({
	symbol,
	className,
}: TopEarningProfilesProps) {
	const [range, setRange] = useState("1D")
	const queryString = `/api/analystics/top-earning-profiles?symbol=${symbol}&range=${range}`
	const { data: rawData, error } = useSWR(queryString, fetcher)
	if (error) return <Error msg={error.message} />
	if (!rawData) return <Loader fixed={false} />

	// only get name, value from data
	const data: any = rawData.map((item: TopEarningProfile) => {
		return {
			name: item.profile_handle,
			value: item.amount,
			href: item.profile_link,
			icon: function Icon() {
				return (
					<img
						src={item.profile_picture}
						alt={item.profile_handle}
						className="w-6 h-6 rounded-full mx-3"
					/>
				)
			},
		}
	})

	return (
		<ChartCard
			chartTitle="Top Earning Profiles"
			range={range}
			setRange={setRange}
			className={className}
		>
			<BarList data={data} showAnimation />
		</ChartCard>
	)
}
