import lensClient from "@/lib/lensclient"
import NotFound from "@/components/404"
import ProfileSummaryCard from "@/components/profile-summary-card"
import Publications from "@/components/publications"

interface PageProps {
  params: {
    id: string
  }
  searchParams: {
    [key: string]: string | string[] | undefined
  }
}
export default async function Page({ params, searchParams }: PageProps) {
  const profile = await lensClient.profile.fetch({
    forProfileId: params.id,
  })

  if (!profile) {
    return <NotFound type="Profile" />
  }

  // add profile_id to filter
  searchParams.profile_id = params.id

  return (
    <>
      <ProfileSummaryCard profile={profile} />
      <Publications searchParams={searchParams} showToolbar showPagination />
    </>
  )
}
