import Layout from "@/components/Layout";
import ProfileTableView from "@/components/ProfileTableView";

export default function Profiles() {
  return (
    <Layout>
      <div className="mt-6">
        <ProfileTableView />
      </div>
    </Layout>
  );
}
