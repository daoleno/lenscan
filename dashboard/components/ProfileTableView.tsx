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
import { Loading } from "./loading";
import { Badge } from "./ui/badge";

import { age, shortHash } from "@/lib/utils";

import { trpc } from "@/lib/trpc";
import { useState } from "react";
import Pagination from "./Pagination";

interface ProfileTableViewProps {
  profileId: string;
  showPagination?: boolean;
  itemsPerPage?: number;
}
export default function ProfileTableView({
  profileId,
  showPagination = true,
  itemsPerPage = 25,
}: ProfileTableViewProps) {
  let id = parseInt(profileId);
  const [cursor, setCursor] = useState(null);
  const { data, error, isLoading } = trpc.event.getEventsByProfileId.useQuery(
    {
      profileId: id,
      take: itemsPerPage,
      cursor,
    },
    {
      enabled: !!id,
    }
  );

  if (error) return <div>Error: {error.message}</div>;
  if (isLoading) return <Loading fixed={false} />;

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
