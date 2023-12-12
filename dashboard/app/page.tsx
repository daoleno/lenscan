import Publications, { PublicationsProps } from "@/components/publications"
import StatCards from "@/components/stat-cards"
import StatChartCard from "@/components/stat-chart-card"
import TopProfiles from "@/components/top-profiles"

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
      <div className="flex gap-4">
        <div className="w-2/3">
          <TopProfiles />
        </div>
        <div className="flex w-full flex-col gap-1.5">
          <StatChartCard />
          <StatChartCard />
          <StatChartCard />
        </div>
      </div>

      <div className="mt-6 space-y-6">
        {/* <ChartView /> */}
        <Publications
          title="Recent Publications"
          searchParams={params.searchParams}
        />
      </div>
    </div>
  )
}
