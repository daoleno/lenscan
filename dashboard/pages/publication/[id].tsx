import Layout from "@/components/Layout";
import Publication from "@/components/publication";
import { useRouter } from "next/router";

export default function Event() {
  const { id } = useRouter().query;

  return (
    <Layout>
      <Publication id={id as string} />
    </Layout>
  );
}
