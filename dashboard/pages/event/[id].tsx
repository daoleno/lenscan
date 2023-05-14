import NotFound from "@/components/404";
import { Loading } from "@/components/loading";
import { Tip } from "@/components/tip";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { trpc } from "@/lib/trpc";
import { age } from "@/lib/utils";
import { Event } from "@prisma/client";
import { ethers } from "ethers";
import { HelpCircle } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/router";

export default function Event() {
  const { id } = useRouter().query;
  const {
    data: event,
    error,
    isLoading,
  } = trpc.event.getEvent.useQuery(Number(id));

  if (error) return <div>Error: {error.message} </div>;
  if (isLoading) return <Loading />;
  if (!event) return <NotFound type="Event" />;

  const eventFields = [
    {
      label: "Block Number:",
      tip: "The block number in which the event was emitted.",
      href: (event: Event) =>
        `https://polygonscan.com/block/${event.blockNumber}`,
      text: (event: Event) => event.blockNumber,
    },
    {
      label: "Transaction Hash:",
      tip: "The transaction hash in which the event was emitted.",
      href: (event: Event) => `https://polygonscan.com/tx/${event.txHash}`,
      text: (event: Event) => event.txHash,
    },
    {
      label: "Transaction Index:",
      tip: "The transaction index in which the event was emitted.",
      href: (event: Event) => `https://polygonscan.com/tx/${event.txHash}`,
      text: (event: Event) => event.txIndex,
    },
    {
      label: "Log Index:",
      tip: "The log index in which the event was emitted.",
      href: (event: Event) => `https://polygonscan.com/tx/${event.txHash}`,
      text: (event: Event) => event.logIndex,
    },
    {
      label: "Removed:",
      tip: "Whether the event was removed from the blockchain.",
      text: (event: Event) => (event.removed ? "Yes" : "No"),
    },
    {
      label: "Timestamp:",
      tip: "The timestamp in which the event was emitted.",
      text: (event: Event) => age(Number(event.timestamp)),
    },
    {
      label: "Event:",
      tip: "The type of event emitted.",
      text: (event: Event) => <Badge variant={"outline"}>{event.type}</Badge>,
    },
  ];

  return (
    <>
      <h1 className="py-7 text-xl font-bold">Event #{event.id}</h1>
      <Card>
        <CardHeader className="font-semibold">Transaction Details</CardHeader>
        <CardContent className="flex flex-col">
          <div className="flex flex-col">
            <div className="flex flex-col gap-4">
              {eventFields.map((field, index) => (
                <div className="flex items-center" key={index}>
                  <Tip text={field.tip}>
                    <HelpCircle className="h-4 text-gray-500" />
                  </Tip>
                  <label className="basis-3/12 text-gray-500">
                    {field.label}
                  </label>
                  {field.href ? (
                    <Link
                      href={field.href(event)}
                      target="_blank"
                      rel="noreferrer"
                      className="basis-9/12 underline underline-offset-4"
                    >
                      {field.text(event) as any}
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
        <CardHeader className="flex flex-col space-y-3 font-semibold">
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
          <div className="flex basis-9/12 flex-col gap-2">
            {Object.entries(event.data || {}).map(([key, value]) => (
              <div key={key} className="flex items-center">
                <Tip text={key}>
                  <HelpCircle className="h-4 text-gray-500" />
                </Tip>
                <label className="basis-3/12 text-gray-500">{key}:</label>
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
