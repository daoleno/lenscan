import {
  PageHeader,
  PageHeaderDescription,
  PageHeaderHeading,
} from "@/components/page-header"
import PublicationsTable from "@/components/publications-table"
import StatCardGrid from "@/components/StatCardGrid"

import getPublications from "./api/publications/getPublications"

export default async function Home() {
  const { publications } = await getPublications({
    limit: 10,
    offset: 0,
  })

  return (
    <div className="mt-6 space-y-6">
      <PageHeader>
        <PageHeaderHeading>Lens Protocol Explorer</PageHeaderHeading>
        <PageHeaderDescription>
          Explore the interesting events happening on the Lens Protocol.
        </PageHeaderDescription>
      </PageHeader>
      <StatCardGrid />
      <div className="mt-6 space-y-6">
        {/* <ChartView /> */}
        <PublicationsTable
          data={publications}
          pageCount={1}
          showToolbar={false}
          showPagination={false}
        />
      </div>
    </div>
  )
}
