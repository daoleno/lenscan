import NotFound from "./404";
import { Loading } from "./loading";

import { Tip } from "@/components/tip";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { getJSONObj } from "@/lib/utils";
import {
  DataAvailabilityTransactionQuery,
  useDataAvailabilityTransactionQuery,
  usePublicationQuery,
} from "@/src/generated";
import { AlertOctagon, HelpCircle, Verified } from "lucide-react";
import Link from "next/link";
import DynamicReactJson from "./dynamic-react-json";

export default function MomokaTx({ id }: { id: string }) {
  const { data, loading, error } = useDataAvailabilityTransactionQuery({
    variables: { request: { id } },
    skip: !id,
  });

  const { data: pub } = usePublicationQuery({
    variables: {
      request: {
        publicationId: data?.dataAvailabilityTransaction?.publicationId,
      },
    },
  });

  if (loading) {
    return <Loading />;
  }

  if (error || !data?.dataAvailabilityTransaction) {
    console.error(error);
    return <NotFound type="Momoka Transaction" />;
  }

  const { dataAvailabilityTransaction } = data;
  const isVerified =
    dataAvailabilityTransaction.verificationStatus.__typename ===
    "DataAvailabilityVerificationStatusSuccess";

  console.log(pub);

  const txFields = [
    {
      label: "Transaction ID:",
      tip: "The data availability transaction ID.",
      href: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => `https://arweave.app/tx/${tx?.transactionId}`,
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => tx?.transactionId,
    },
    {
      label: "Publication ID:",
      tip: "The transaction hash in which the tx was emitted.",
      href: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => `/publication/${tx?.publicationId}`,
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => tx?.publicationId,
    },
    {
      label: "Verified:",
      tip: "Whether the transaction was verified.",
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => (isVerified ? <Verified /> : <AlertOctagon />),
    },
    {
      label: "Submitter:",
      tip: "The address of the submitter.",
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => tx?.submitter,
    },
    {
      label: "Created At:",
      tip: "The date and time of the transaction.",
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => tx?.createdAt,
    },
    {
      label: "Posted via",
      tip: "Publication posted via.",
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => tx?.appId,
    },
    {
      label: "Event:",
      tip: "The type of event emitted.",
      text: (
        tx: DataAvailabilityTransactionQuery["dataAvailabilityTransaction"]
      ) => <Badge variant={"outline"}>{tx?.__typename}</Badge>,
    },
  ];

  return (
    <>
      <div className="flex items-center space-x-3 py-7 text-2xl font-bold text-gray-800">
        <span>Transaction</span>
        <Badge>MOMOKA</Badge>
      </div>
      <Card>
        <CardHeader className="font-semibold">Overview</CardHeader>
        <CardContent className="flex flex-col">
          <div className="flex flex-col">
            <div className="flex flex-col gap-4">
              {txFields.map((field, index) => (
                <div className="flex items-center" key={index}>
                  <Tip text={field.tip}>
                    <HelpCircle className="h-4 text-gray-500" />
                  </Tip>
                  <label className="basis-3/12 text-gray-500">
                    {field.label}
                  </label>
                  {field.href ? (
                    <Link
                      href={field.href(dataAvailabilityTransaction)}
                      target="_blank"
                      rel="noreferrer"
                      className="basis-9/12 underline underline-offset-4"
                    >
                      {field.text(dataAvailabilityTransaction)}
                    </Link>
                  ) : (
                    <span className="basis-9/12">
                      {field.text(dataAvailabilityTransaction)}
                    </span>
                  )}
                </div>
              ))}
            </div>
          </div>
        </CardContent>
      </Card>
      <Card className="mt-3">
        <CardHeader className="flex flex-col space-y-3 font-semibold">
          <CardTitle>Detail</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex gap-4">
            <div className="basis-3/12  text-gray-800">
              On-chain Content URI
            </div>
            <div className="flex basis-9/12 items-center">
              <Link
                href={pub?.publication?.onChainContentURI || ""}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12 underline underline-offset-4"
              >
                {pub?.publication?.onChainContentURI}
              </Link>
            </div>
          </div>
          <DynamicReactJson
            src={getJSONObj(pub?.publication?.onChainContentURI || "")}
          />
        </CardContent>
      </Card>
    </>
  );
}
