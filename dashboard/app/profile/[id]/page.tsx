import { getProfileRevenue } from "@/app/api/analystics/revenue/getProfileRevenue"
import NotFound from "@/components/404"
import ProfileRecord from "@/components/profile-records"
import ProfileSummaryCard from "@/components/profile-summary-card"
import { RevalidateTime } from "@/config/ssr"
import lensClient from "@/lib/lensclient"
import { isHex } from "viem"

export const revalidate = RevalidateTime

interface PageProps {
	params: {
		id: string
	}
	searchParams: {
		[key: string]: string | string[] | undefined
	}
}
export default async function Page({ params, searchParams }: PageProps) {
	const fetchParams = isHex(params.id)
		? {
				forProfileId: params.id,
		  }
		: {
				forHandle: `lens/${params.id}`,
		  }
	const profile = await lensClient.profile.fetch(fetchParams)

	if (!profile) {
		return <NotFound type="Profile" />
	}

	// add profile_id to filter
	searchParams.profile_id = profile.id

	const revenue = await getProfileRevenue(profile.id)

	return (
		<>
			<ProfileSummaryCard profile={profile} revenue={revenue} />
			<ProfileRecord searchParams={searchParams} />
		</>
	)
}
