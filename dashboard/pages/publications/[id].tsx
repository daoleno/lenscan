import Layout from "@/components/Layout";
import ProfileSummaryCard from "@/components/ProfileSummaryCard";
import ProfileTableView from "@/components/ProfileTableView";
import { useRouter } from "next/router";

export default function Event() {
  const { id } = useRouter().query;

  return (
    <Layout>
      <ProfileSummaryCard profileId={id as string} />
      {/* <ProfileStatCardGrid profileId={id as string} /> */}
      <ProfileTableView profileId={id as string} />
    </Layout>
  );
}
