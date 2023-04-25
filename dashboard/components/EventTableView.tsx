import { Event } from "@/types";
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

import { db } from "@/lib/postgrest";
import { useState } from "react";
import useSWR from "swr";
import Pagination from "./Pagination";
import { Loading } from "./loading";
import { Badge } from "./ui/badge";

export default function EventTableView({
  showPagination = true,
  itemsPerPage = 25,
}) {
  const [range, setRange] = useState([0, itemsPerPage - 1]);

  const { data, error } = useSWR(
    `latest-events-${range[0]}-${range[1]}`,
    async () => {
      return await db.Events(range[0], range[1]);
    }
    // {
    //   refreshInterval: 5000, // refresh data every 5 seconds
    // }
  );

  const [currentPage, setCurrentPage] = useState(1);

  if (error) return <div>Error: {error.message}</div>;
  if (!data || !data.data) return <Loading />;
  const { data: events, count }: any = data;

  function handlePageChange(page: number) {
    setCurrentPage(page);
    setRange([(page - 1) * 25, page * 25]);
  }

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
        <div className="flex space-x-2">
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
            {events.map((item: Event) => (
              <TableRow key={item.id}>
                <TableCell>
                  <Link
                    href={`/events/${item.id}`}
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
                    {item.blockNumber}
                  </Link>
                </TableCell>
                <TableCell>{age(item.timestamp)}</TableCell>
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
          currentPage={currentPage}
          totalResults={count}
          resultsPerPage={range[1] - range[0]}
          onPageChange={handlePageChange}
        />
      )}
    </div>
  );
}
