import Layout from "@/components/Layout";
import ProfilesTableView from "@/components/ProfilesTableView";

export default function Profiles() {
  return (
    <Layout>
      <div className="mt-6">
        <ProfilesTableView />
      </div>
    </Layout>
  );
}
