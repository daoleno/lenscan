import ContentLanguageDistribution from "@/components/charts/content-language-distribution"
import PopularHashtags from "@/components/charts/hashtags"
import PublicationReactions from "@/components/charts/publication-reactions"
import PublicationTypesDistribution from "@/components/charts/publication-types-distribution"
import TopContributors from "@/components/charts/top-contributors"
import UserActivity from "@/components/charts/user-activity"

export default async function Page() {
  return (
    <div className="grid grid-cols-1 gap-4 p-7 md:grid-cols-2">
      <UserActivity />
      <PopularHashtags />
      <PublicationTypesDistribution />
      <PublicationReactions />
      <ContentLanguageDistribution />
      <TopContributors />
    </div>
  )
}
