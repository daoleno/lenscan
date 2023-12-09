import PopularHashtags from "@/components/charts/hashtags"
import PublicationReactions from "@/components/charts/publication-reactions"
import PublicationTypesDistribution from "@/components/charts/publication-types-distribution"
import UserActivity from "@/components/charts/user-activity"

export default async function Page() {
  return (
    <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
      <UserActivity />
      <PopularHashtags />
      <PublicationTypesDistribution />
      <PublicationReactions />
    </div>
  )
}
