import ContentLanguageDistribution from "@/components/charts/content-language-distribution"
import DAU from "@/components/charts/dau"
import PopularHashtags from "@/components/charts/hashtags"
import MAU from "@/components/charts/mau"
import PublicationTypesDistribution from "@/components/charts/publication-types-distribution"
import UserActivity from "@/components/charts/user-activity"

export default function Page() {
  return (
    <div className="grid grid-cols-1 gap-4 p-7 md:grid-cols-2">
      <DAU />
      <MAU />
      <UserActivity />
      <PopularHashtags />
      <PublicationTypesDistribution />
      <ContentLanguageDistribution />
    </div>
  )
}
