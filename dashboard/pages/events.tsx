import EventTableView from "@/components/EventTableView";
import Layout from "@/components/Layout";

export default function Events() {
  return (
    <Layout>
      <div className="mt-6">
        <EventTableView />
      </div>
    </Layout>
  );
}
