import { Event } from "@/types";
import {
  Badge,
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
  if (!data || !data.data) return <div>Loading...</div>;
  const { data: events, count }: any = data;

  function handlePageChange(page: number) {
    setCurrentPage(page);
    setRange([(page - 1) * 25, page * 25]);
  }

  return (
    <div className="mt-6">
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
                    className="text-blue-500 hover:text-blue-600"
                  >
                    {item.id}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    href={`https://polygonscan.com/block/${item.blockNumber}`}
                    target="_blank"
                    className="text-blue-500 hover:text-blue-600"
                  >
                    {item.blockNumber}
                  </Link>
                </TableCell>
                <TableCell>{age(item.timestamp)}</TableCell>
                <TableCell>
                  <Link
                    href={`https://polygonscan.com/tx/${item.txHash}`}
                    target="_blank"
                    className="text-blue-500 hover:text-blue-600"
                  >
                    {shortHash(item.txHash!)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    href={`https://polygonscan.com/tx/${item.txHash}#eventlog`}
                    target="_blank"
                    className="text-blue-500 hover:text-blue-600"
                  >
                    {item.logIndex}
                  </Link>
                </TableCell>
                <TableCell>
                  <Badge size="xs" color="green">
                    {item.type}
                  </Badge>
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
