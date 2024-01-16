import DailyProfileStats from "@/components/daily-profile-stats"
import DailyPublicationStats from "@/components/daily-publication-stats"
import DailyApps from "@/components/dailyapps"
import DauStats from "@/components/dau-stats"
import Publications, { PublicationsProps } from "@/components/publications"
import StatCards from "@/components/stat-cards"

export const revalidate = 60 * 60 * 5

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
      <div className="grid grid-cols-1 items-stretch gap-4 sm:grid-cols-3">
        <div className="sm:col-span-2">
          <DailyApps />
        </div>
        <div className="grid gap-3">
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
