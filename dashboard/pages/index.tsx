import Layout from "@/components/Layout";
import StatCardGrid from "@/components/StatCardGrid";
import TableView from "@/components/TableView";

export default function Home() {
  return (
    <Layout>
      <StatCardGrid />
      <div className="mt-6 space-y-6">
        {/* <ChartView /> */}
        <TableView start={0} end={9} />
      </div>
    </Layout>
  );
}
