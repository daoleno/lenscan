import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
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

import { age, formatEventType, getIPFSURL, shortHash } from "@/lib/utils";

import { trpc } from "@/lib/trpc";
import { AlertOctagon, Verified } from "lucide-react";
import Image from "next/image";
import { useState } from "react";
import Pagination from "./pagination";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";

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
  const {
    data: momokaTxs,
    error: momokaTxsError,
    isLoading: momokaTxsLoading,
  } = trpc.momoka.getTxsByProfileId.useQuery(
    {
      profileId: id,
      take: itemsPerPage,
      cursor,
    },
    {
      enabled: !!id,
    }
  );
  const { data: count } = trpc.event.getEventsCountByProfileId.useQuery(id);
  if (error || momokaTxsError)
    return <div>Error: {error?.message || momokaTxsError?.message}</div>;
  if (isLoading && momokaTxsLoading) return <Loading fixed={false} />;

  return (
    <Tabs className="mt-6 flex flex-col space-y-3" defaultValue="events">
      <TabsList className="mx-auto flex items-start space-x-3">
        <TabsTrigger value="events">Polygon Events</TabsTrigger>
        <TabsTrigger value="momoka-txns">Momoka Transactions</TabsTrigger>
      </TabsList>
      <TabsContent value="events">
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
              {data?.events.map((item) => (
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
            nextCursor={data?.nextCursor}
            totalResults={count}
            resultsPerPage={itemsPerPage}
            setCursor={setCursor}
          />
        )}
      </TabsContent>
      <TabsContent value="momoka-txns">
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
              {momokaTxs?.list.map((item: any) => (
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
            nextCursor={momokaTxs?.nextCursor}
            // totalResults={momokaTxs?.count}
            resultsPerPage={itemsPerPage}
            setCursor={setCursor}
          />
        )}
      </TabsContent>
    </Tabs>
  );
}
