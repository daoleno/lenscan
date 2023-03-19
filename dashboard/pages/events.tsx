import Layout from "@/components/Layout";
import TableView from "@/components/TableView";

export default function Home() {
  return (
    <Layout>
      <div className="mt-6">
        <TableView start={0} end={19} />
      </div>
    </Layout>
  );
}
