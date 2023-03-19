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

import { shortHash } from "@/lib/utils";

import { db } from "@/lib/postgrest";
import { useState } from "react";
import useSWR from "swr";
import Pagination from "./Pagination";

interface PaginationProps {
  currentPage: number;
  totalResults: number;
  resultsPerPage: number;
  onPageChange: (pageNumber: number) => void;
}

export default function TableView({
  start,
  end,
}: {
  start: number;
  end: number;
}) {
  const { data, error } = useSWR(
    `latest-events-${start}-${end}`,
    async () => {
      return await db.Event(start, end);
    },
    {
      refreshInterval: 5000, // refresh data every 5 seconds
    }
  );

  console.log("data", data);

  const [currentPage, setCurrentPage] = useState(1);

  if (error) return <div>Error: {error.message}</div>;
  if (!data || !data.data) return <div>Loading...</div>;
  const { data: events, count }: any = data;

  console.log("events", events);

  function handlePageChange(page: number) {
    setCurrentPage(page);
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
            {events
              .slice((currentPage - 1) * 25, currentPage * 25)
              .map((item: Event) => (
                <TableRow key={item.id}>
                  <TableCell>
                    <Link
                      href={`https://polygonscan.com/block/${item.id}`}
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
                  <TableCell>a few seconds ago</TableCell>
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
                      {item.event}
                    </Badge>
                  </TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </Card>
      <Pagination
        currentPage={currentPage}
        totalResults={1}
        resultsPerPage={end - start}
        onPageChange={handlePageChange}
      />
    </div>
  );
}
