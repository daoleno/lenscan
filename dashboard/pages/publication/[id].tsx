import Publication from "@/components/publication";
import { useRouter } from "next/router";

export default function PublicationPage() {
  const { id } = useRouter().query;

  return <Publication id={id as string} />;
}
