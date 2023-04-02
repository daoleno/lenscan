import Layout from "@/components/Layout";
import { db } from "@/lib/postgrest";
import { Badge, Card } from "@tremor/react";
import { useRouter } from "next/router";
import useSWR from "swr";

export default function Event() {
  const { id } = useRouter().query;
  const { data, error } = useSWR(`event-${id}`, async () => {
    return await db.Event(parseInt(id as string));
  });
  if (error) return <div>Error: {error.message} </div>;
  if (!data || !data.data) return <div>Loading...</div>;
  const { data: event }: any = data;

  console.log("event", event);

  return (
    <Layout>
      <h1 className="text-xl font-bold py-7">Event #{event.id}</h1>
      <Card>
        <div className="flex flex-col">
          <div className="flex flex-col gap-4">
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">Block Number:</label>
              <a
                href={`https://polygonscan.com/block/${event.blockNumber}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12"
              >
                {event.blockNumber}
              </a>
            </div>
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">
                Transaction Hash:
              </label>
              <a
                href={`https://polygonscan.com/tx/${event.txHash}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12"
              >
                {event.txHash}
              </a>
            </div>
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">
                Transaction Index:
              </label>
              <a
                href={`https://polygonscan.com/tx/${event.txHash}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12"
              >
                {event.txIndex}
              </a>
            </div>
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">Log Index:</label>
              <a
                href={`https://polygonscan.com/tx/${event.txHash}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12"
              >
                {event.logIndex}
              </a>
            </div>
          </div>
          <div className="flex items-center mt-4">
            <label className="text-gray-500 basis-3/12">Removed:</label>
            <span className="basis-9/12">{event.removed ? "Yes" : "No"}</span>
          </div>
          <div className="flex mt-4">
            <label className="text-gray-500 basis-3/12">Event:</label>
            <div className="flex items-center basis-9/12">
              <Badge size="xs" color="green">
                {event.event}
              </Badge>
            </div>
          </div>
        </div>
      </Card>
    </Layout>
  );
}
