import EventTableView from "@/components/EventTableView";
import Layout from "@/components/Layout";
import StatCardGrid from "@/components/StatCardGrid";
import {
  PageHeader,
  PageHeaderDescription,
  PageHeaderHeading,
} from "@/components/page-header";

export default function Home() {
  return (
    <Layout>
      <PageHeader>
        <PageHeaderHeading>Lens Protocol Explorer</PageHeaderHeading>
        <PageHeaderDescription>
          Explore the interesting events happening on the Lens Protocol.
        </PageHeaderDescription>
      </PageHeader>
      <StatCardGrid />
      <div className="mt-6 space-y-6">
        {/* <ChartView /> */}
        <EventTableView showPagination={false} itemsPerPage={10} />
      </div>
    </Layout>
  );
}
