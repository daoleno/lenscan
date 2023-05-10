import MomokaTx from "@/components/momoka-tx";
import { useRouter } from "next/router";

export default function Tx() {
  const { id } = useRouter().query;

  return <MomokaTx id={id as string} />;
}
