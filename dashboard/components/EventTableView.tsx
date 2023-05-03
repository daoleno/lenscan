import {
  Card,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeaderCell,
  TableRow,
} from "@tremor/react";
import Link from "next/link";

import { age, shortHash } from "@/lib/utils";

import { trpc } from "@/lib/trpc";
import { useState } from "react";
import Pagination from "./Pagination";
import { Loading } from "./loading";
import { Badge } from "./ui/badge";

export default function EventTableView({
  showPagination = true,
  itemsPerPage = 25,
}) {
  const [cursor, setCursor] = useState(null);
  const { data, error, isLoading } = trpc.event.getEvents.useQuery({
    take: itemsPerPage,
    cursor,
  });

  if (error) return <div>Error: {error.message}</div>;
  if (isLoading) return <Loading fixed={false} />;

  const eventTypes = [
    "Followed",
    "PostCreated",
    "CommentCreated",
    "MirrorCreated",
    "Collected",
  ];

  return (
    <div>
      <div className="flex flex-col justify-between space-y-7 pb-7">
        <h2 className="text-3xl font-bold tracking-tight">Events</h2>
        <div className="flex flex-wrap gap-2">
          {eventTypes.map((type) => (
            <Badge key={type} variant="outline">
              {type}
            </Badge>
          ))}
        </div>
      </div>
      <Card>
        <Table>
          <TableHead>
            <TableRow>
              <TableHeaderCell>Id</TableHeaderCell>
              <TableHeaderCell>Block</TableHeaderCell>
              <TableHeaderCell>Age</TableHeaderCell>
              <TableHeaderCell>Txn Hash</TableHeaderCell>
              <TableHeaderCell>Log Index</TableHeaderCell>
              <TableHeaderCell>Event</TableHeaderCell>
            </TableRow>
          </TableHead>

          <TableBody>
            {data.events.map((item) => (
              <TableRow key={item.id}>
                <TableCell>
                  <Link
                    href={`/event/${item.id}`}
                    target="_blank"
                    className="font-medium underline underline-offset-4"
                  >
                    {item.id}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    href={`https://polygonscan.com/block/${item.blockNumber}`}
                    target="_blank"
                    className="font-medium underline underline-offset-4"
                  >
                    {Number(item.blockNumber)}
                  </Link>
                </TableCell>
                <TableCell>{age(Number(item.timestamp))}</TableCell>
                <TableCell>
                  <Link
                    href={`https://polygonscan.com/tx/${item.txHash}`}
                    target="_blank"
                    className="font-medium underline underline-offset-4"
                  >
                    {shortHash(item.txHash!)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    href={`https://polygonscan.com/tx/${item.txHash}#eventlog`}
                    target="_blank"
                    className="font-medium underline underline-offset-4"
                  >
                    {item.logIndex}
                  </Link>
                </TableCell>
                <TableCell>
                  <Badge variant={"outline"}>{item.type}</Badge>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Card>
      {showPagination && (
        <Pagination
          curCursor={cursor}
          nextCursor={data.nextCursor}
          totalResults={data.count}
          resultsPerPage={itemsPerPage}
          setCursor={setCursor}
        />
      )}
    </div>
  );
}
