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

import { age, formatEventType, getIPFSURL, shortHash } from "@/lib/utils";
import Image from "next/image";

import { trpc } from "@/lib/trpc";
import { AlertOctagon, Verified } from "lucide-react";
import { useState } from "react";
import { Loading } from "./loading";
import MomokaDailyTxsChart from "./momoka-daily-txs-charts";
import MomokaStats from "./momoka-stats";
import Pagination from "./pagination";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { Badge } from "./ui/badge";

export default function MomokaTxsTableView({
  showPagination = true,
  itemsPerPage = 25,
}) {
  const [cursor, setCursor] = useState(null);
  const { data, error, isLoading } = trpc.momoka.getTxs.useQuery({
    take: itemsPerPage,
    cursor,
  });

  if (error) return <div>Error: {error.message}</div>;
  if (isLoading) return <Loading fixed={false} />;

  const eventTypes = ["PostCreated", "CommentCreated", "MirrorCreated"];

  return (
    <div>
      <div className="flex flex-col justify-between space-y-7 pb-7">
        <h2 className="text-3xl font-bold tracking-tight">
          Momoka Transactions
        </h2>
        <div className="flex flex-wrap gap-2">
          {eventTypes.map((type) => (
            <Badge key={type} variant="outline">
              {type}
            </Badge>
          ))}
        </div>
        <MomokaStats />
        <MomokaDailyTxsChart />
      </div>

      <Card>
        <Table>
          <TableHead>
            <TableRow>
              <TableHeaderCell>Txn Id</TableHeaderCell>
              <TableHeaderCell>Publication Id</TableHeaderCell>
              <TableHeaderCell>Age</TableHeaderCell>
              <TableHeaderCell>Sender</TableHeaderCell>
              <TableHeaderCell>Verified</TableHeaderCell>
              <TableHeaderCell>Event</TableHeaderCell>
            </TableRow>
          </TableHead>

          <TableBody>
            {data.list.map((item: any) => (
              <TableRow key={item.proofTxId}>
                <TableCell>
                  <Link
                    href={`/momoka-tx/${item.proofTxId}`}
                    target="_blank"
                    className="font-medium underline underline-offset-4"
                  >
                    {shortHash(item.proofTxId)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    href={`/publication/${item.publicationId}`}
                    target="_blank"
                    className="font-medium underline underline-offset-4"
                  >
                    {item.publicationId}
                  </Link>
                </TableCell>
                <TableCell>{age(Number(item.timestamp))}</TableCell>
                <TableCell>
                  <Link
                    href={`/profile/${(item.event as any).profileId}`}
                    target="_blank"
                    className="flex items-center space-x-3 font-medium"
                  >
                    <Avatar>
                      <AvatarImage
                        src={getIPFSURL(item.profile.picture)}
                        alt={item.profile.handle}
                      />
                      <AvatarFallback>
                        <Image
                          src="/images/default-profile.png"
                          alt={item.profile.handle}
                          width={32}
                          height={32}
                        />
                      </AvatarFallback>
                    </Avatar>
                    <span> {item.profile.handle}</span>
                  </Link>
                </TableCell>
                <TableCell>
                  {item.success ? <Verified /> : <AlertOctagon />}
                </TableCell>
                <TableCell>
                  <Badge variant={"outline"}>
                    {formatEventType(item.type)}
                  </Badge>
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
