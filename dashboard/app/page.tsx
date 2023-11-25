import StatCardGrid from "@/components/StatCardGrid";
import {
  PageHeader,
  PageHeaderDescription,
  PageHeaderHeading,
} from "@/components/page-header";

export default function Home() {
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
        {/* <EventTableView showPagination={false} itemsPerPage={10} /> */}
      </div>
    </div>
  );
}
