import { isHex } from "viem"

import NotFound from "@/components/404"
import ProfileSummaryCard from "@/components/profile-summary-card"
import Publications from "@/components/publications"
import { RevalidateTime } from "@/config/ssr"
import lensClient from "@/lib/lensclient"

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

	return (
		<>
			<ProfileSummaryCard profile={profile} />
			<Publications searchParams={searchParams} showToolbar showPagination />
		</>
	)
}
