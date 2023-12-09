import {
  PageHeader,
  PageHeaderDescription,
  PageHeaderHeading,
} from "@/components/page-header"
import Publications, { PublicationsProps } from "@/components/publications"
import StatCards from "@/components/statcards"

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
      <PageHeader>
        <PageHeaderHeading>Lens Protocol Explorer</PageHeaderHeading>
        <PageHeaderDescription>
          Explore the interesting events happening on the Lens Protocol.
        </PageHeaderDescription>
      </PageHeader>
      <StatCards />
      <div className="mt-6 space-y-6">
        {/* <ChartView /> */}
        <Publications searchParams={params.searchParams} />
      </div>
    </div>
  )
}
