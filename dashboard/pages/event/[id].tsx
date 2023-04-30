import { Loading } from "@/components/loading";
import { Tip } from "@/components/tip";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { db } from "@/lib/postgrest";
import { age } from "@/lib/utils";
import { definitions } from "@/types/generated-types";
import { ethers } from "ethers";
import { HelpCircle } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/router";
import useSWR from "swr";

export default function Event() {
  const { id } = useRouter().query;
  const { data, error } = useSWR(`event-${id}`, async () => {
    return await db.Event(parseInt(id as string));
  });
  if (error) return <div>Error: {error.message} </div>;
  if (!data || !data.data) return <Loading />;
  const { data: event }: any = data;
  const eventFields = [
    {
      label: "Block Number:",
      tip: "The block number in which the event was emitted.",
      href: (event: definitions["Event"]) =>
        `https://polygonscan.com/block/${event.blockNumber}`,
      text: (event: definitions["Event"]) => event.blockNumber,
    },
    {
      label: "Transaction Hash:",
      tip: "The transaction hash in which the event was emitted.",
      href: (event: definitions["Event"]) =>
        `https://polygonscan.com/tx/${event.txHash}`,
      text: (event: definitions["Event"]) => event.txHash,
    },
    {
      label: "Transaction Index:",
      tip: "The transaction index in which the event was emitted.",
      href: (event: definitions["Event"]) =>
        `https://polygonscan.com/tx/${event.txHash}`,
      text: (event: definitions["Event"]) => event.txIndex,
    },
    {
      label: "Log Index:",
      tip: "The log index in which the event was emitted.",
      href: (event: definitions["Event"]) =>
        `https://polygonscan.com/tx/${event.txHash}`,
      text: (event: definitions["Event"]) => event.logIndex,
    },
    {
      label: "Removed:",
      tip: "Whether the event was removed from the blockchain.",
      text: (event: definitions["Event"]) => (event.removed ? "Yes" : "No"),
    },
    {
      label: "Timestamp:",
      tip: "The timestamp in which the event was emitted.",
      text: (event: definitions["Event"]) => age(event.timestamp),
    },
    {
      label: "Event:",
      tip: "The type of event emitted.",
      text: (event: definitions["Event"]) => (
        <Badge variant={"outline"}>{event.type}</Badge>
      ),
    },
  ];

  return (
    <>
      <h1 className="text-xl font-bold py-7">Event #{event.id}</h1>
      <Card>
        <CardHeader className="font-semibold">Transaction Details</CardHeader>
        <CardContent className="flex flex-col">
          <div className="flex flex-col">
            <div className="flex flex-col gap-4">
              {eventFields.map((field, index) => (
                <div className="flex items-center" key={index}>
                  <Tip text={field.tip}>
                    <HelpCircle className="text-gray-500 h-4" />
                  </Tip>
                  <label className="text-gray-500 basis-3/12">
                    {field.label}
                  </label>
                  {field.href ? (
                    <Link
                      href={field.href(event)}
                      target="_blank"
                      rel="noreferrer"
                      className="basis-9/12 underline underline-offset-4"
                    >
                      {field.text(event)}
                    </Link>
                  ) : (
                    <span className="basis-9/12">{field.text(event)}</span>
                  )}
                </div>
              ))}
            </div>
          </div>
        </CardContent>
      </Card>
      <Card className="mt-3">
        <CardHeader className="flex flex-col font-semibold space-y-3">
          <CardTitle>Event Data</CardTitle>
          {event.data &&
            (event.type == "PostCreated" ||
              event.type == "CommentCreated" ||
              event.type == "MirrorCreated" ||
              event.type == "Collected") && (
              <div className="flex items-center space-x-3">
                <Badge>{event.type}</Badge>
                <Link
                  href={`/publication/${ethers.utils.hexlify(
                    (event.data as any)?.ProfileId
                  )}-${ethers.utils.hexlify((event.data as any)?.PubId)}`}
                  className="font-medium underline underline-offset-4"
                >
                  {ethers.utils.hexlify((event.data as any)?.ProfileId)}-
                  {ethers.utils.hexlify((event.data as any)?.PubId)}
                </Link>
              </div>
            )}
        </CardHeader>
        <CardContent>
          <div className="flex flex-col gap-2 basis-9/12">
            {Object.entries(event.data).map(([key, value]) => (
              <div key={key} className="flex items-center">
                <Tip text={key}>
                  <HelpCircle className="text-gray-500 h-4" />
                </Tip>
                <label className="text-gray-500 basis-3/12">{key}:</label>
                <span className="basis-9/12">
                  {value !== undefined ? String(value) : "-"}
                </span>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </>
  );
}
