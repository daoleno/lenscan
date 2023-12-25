import DailyProfileStats from "@/components/daily-profile-stats"
import DailyPublicationStats from "@/components/daily-publication-stats"
import DauStats from "@/components/dau-stats"
import Publications, { PublicationsProps } from "@/components/publications"
import StatCards from "@/components/stat-cards"
import TopProfiles from "@/components/top-profiles"

export const revalidate = 900

export default async function Home() {
  const params: PublicationsProps = {
    searchParams: {
      page: "1",
      per_page: "10",
      sort: "block_timestamp.desc",
    },
  }
  return (
    <div className="mt-6 space-y-6">
      <StatCards />
      <div className="flex flex-col gap-4 sm:flex-row">
        <div className="w-2/3 ">
          <TopProfiles />
        </div>
        <div className="flex w-full flex-col gap-1.5">
          <DailyPublicationStats />
          <DailyProfileStats />
          <DauStats />
        </div>
      </div>

      <div className="mt-6 space-y-6">
        <Publications
          title="Recent Publications"
          searchParams={params.searchParams}
        />
      </div>
    </div>
  )
}
