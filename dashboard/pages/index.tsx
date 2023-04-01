import EventTableView from "@/components/EventTableView";
import Layout from "@/components/Layout";
import StatCardGrid from "@/components/StatCardGrid";

export default function Home() {
  return (
    <Layout>
      <StatCardGrid />
      <div className="mt-6 space-y-6">
        {/* <ChartView /> */}
        <EventTableView showPagination={false} itemsPerPage={10} />
      </div>
    </Layout>
  );
}
