import { db } from "@/lib/postgrest";
import { age, shortHash } from "@/lib/utils";
import { definitions } from "@/types/generated-types";
import {
  Card,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeaderCell,
  TableRow,
} from "@tremor/react";
import { ethers } from "ethers";
import Link from "next/link";
import { useState } from "react";
import useSWR from "swr";
import Pagination from "./Pagination";
import { Loading } from "./loading";
import { Badge } from "./ui/badge";

export default function PublicationsTableView({
  showPagination = true,
  itemsPerPage = 25,
}) {
  const [range, setRange] = useState([0, itemsPerPage - 1]);
  const { data, error } = useSWR(
    `latest-publications-${range[0]}-${range[1]}`,
    async () => {
      return await db.Publications(range[0], range[1]);
    }
    // {
    //   refreshInterval: 5000, // refresh data every 5 seconds
    // }
  );

  const [currentPage, setCurrentPage] = useState(1);

  if (error) return <div>Error: {error.message}</div>;
  if (!data || !data.data) return <Loading />;
  const { data: events, count }: any = data;

  console.log("events(pub)", events);

  function handlePageChange(page: number) {
    setCurrentPage(page);
    setRange([(page - 1) * 25, page * 25]);
  }

  return (
    <div className="mt-6">
      <div className="flex flex-col justify-between space-y-7 pb-7">
        <h2 className="text-3xl font-bold tracking-tight">Publications</h2>
      </div>
      <Card>
        <Table>
          <TableHead>
            <TableRow>
              <TableHeaderCell>Id</TableHeaderCell>
              <TableHeaderCell>Profile Id</TableHeaderCell>
              <TableHeaderCell>block</TableHeaderCell>
              <TableHeaderCell>Age</TableHeaderCell>
              <TableHeaderCell>Txn Hash</TableHeaderCell>
              <TableHeaderCell>type</TableHeaderCell>
            </TableRow>
          </TableHead>

          <TableBody>
            {events?.map((item: definitions["Event"]) => (
              <TableRow key={item.id}>
                <TableCell>
                  <Link
                    href={`/publication/${ethers.utils.hexlify(
                      (item.data as any)?.ProfileId
                    )}-${ethers.utils.hexlify((item.data as any)?.PubId)}`}
                    className="font-medium underline underline-offset-4"
                  >
                    {ethers.utils.hexlify((item.data as any)?.ProfileId)}-
                    {ethers.utils.hexlify((item.data as any)?.PubId)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    href={`/profile/${ethers.utils.hexlify(
                      (item.data as any)?.ProfileId
                    )}`}
                    className="font-medium underline underline-offset-4"
                  >
                    {(item.data as any)?.ProfileId || "-"}
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
