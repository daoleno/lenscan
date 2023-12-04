import lensClient from "@/lib/lensclient"
import NotFound from "@/components/404"
import ProfileSummaryCard from "@/components/profile-summary-card"

interface PageProps {
  params: {
    id: string
  }
}
export default async function Page({ params }: PageProps) {
  const profile = await lensClient.profile.fetch({
    forProfileId: params.id,
  })

  if (!profile) {
    return <NotFound type="Profile" />
  }

  return (
    <>
      <ProfileSummaryCard profile={profile} />
      {/* <PublicationsTable profileId={profileId} /> */}
    </>
  )
}
