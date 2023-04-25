import Layout from "@/components/Layout";
import { Loading } from "@/components/loading";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { db } from "@/lib/postgrest";
import { age } from "@/lib/utils";
import Link from "next/link";
import { useRouter } from "next/router";
import useSWR from "swr";

export default function Event() {
  const { id } = useRouter().query;
  const { data, error } = useSWR(`event-${id}`, async () => {
    return await db.Event(parseInt(id as string));
  });
  if (error) return <div>Error: {error.message} </div>;
  if (!data || !data.data)
    return (
      <Layout>
        <Loading />
      </Layout>
    );
  const { data: event }: any = data;

  return (
    <Layout>
      <h1 className="text-xl font-bold py-7">Event #{event.id}</h1>
      <Card>
        <div className="flex flex-col">
          <div className="flex flex-col gap-4">
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">Block Number:</label>
              <Link
                href={`https://polygonscan.com/block/${event.blockNumber}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12 text-blue-500 hover:text-blue-600"
              >
                {event.blockNumber}
              </Link>
            </div>
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">
                Transaction Hash:
              </label>
              <Link
                href={`https://polygonscan.com/tx/${event.txHash}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12 text-blue-500 hover:text-blue-600"
              >
                {event.txHash}
              </Link>
            </div>
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">
                Transaction Index:
              </label>
              <Link
                href={`https://polygonscan.com/tx/${event.txHash}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12 text-blue-500 hover:text-blue-600"
              >
                {event.txIndex}
              </Link>
            </div>
            <div className="flex items-center">
              <label className="text-gray-500 basis-3/12">Log Index:</label>
              <Link
                href={`https://polygonscan.com/tx/${event.txHash}`}
                target="_blank"
                rel="noreferrer"
                className="basis-9/12 text-blue-500 hover:text-blue-600"
              >
                {event.logIndex}
              </Link>
            </div>
          </div>
          <div className="flex items-center mt-4">
            <label className="text-gray-500 basis-3/12">Removed:</label>
            <span className="basis-9/12">{event.removed ? "Yes" : "No"}</span>
          </div>
          <div className="flex items-center mt-4">
            <label className="text-gray-500 basis-3/12">Timestamp:</label>
            <span className="basis-9/12">{age(event.timestamp)}</span>
          </div>
          <div className="flex mt-4">
            <label className="text-gray-500 basis-3/12">Event:</label>
            <div className="flex items-center basis-9/12">
              <Badge variant={"outline"}>{event.type}</Badge>
            </div>
          </div>
        </div>
      </Card>
      <Card className="mt-3">
        {event.data && (
          <div className="flex flex-col gap-2 basis-9/12">
            {Object.entries(event.data).map(([key, value]) => (
              <div key={key} className="flex items-center">
                <label className="text-gray-500 basis-3/12">{key}:</label>
                <span className="basis-9/12">
                  {value !== undefined ? String(value) : "-"}
                </span>
              </div>
            ))}
          </div>
        )}
      </Card>
    </Layout>
  );
}
