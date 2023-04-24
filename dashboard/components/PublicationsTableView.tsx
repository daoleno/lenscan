import { db } from "@/lib/postgrest";
import {
  Card,
  Table,
  TableHead,
  TableHeaderCell,
  TableRow,
} from "@tremor/react";
import { useState } from "react";
import useSWR from "swr";
import Pagination from "./Pagination";

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
  if (!data || !data.data) return <div>Loading...</div>;
  const { data: events, count }: any = data;

  console.log("events(pub)", events);

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
              <TableHeaderCell>Name</TableHeaderCell>
              <TableHeaderCell>Handle</TableHeaderCell>
              <TableHeaderCell>ENS</TableHeaderCell>
              <TableHeaderCell>Total Followers</TableHeaderCell>
              <TableHeaderCell>Total Publications</TableHeaderCell>
            </TableRow>
          </TableHead>

          {/* <TableBody>
            {events.map((item: Event) => (
              <TableRow key={item.id}>
                <TableCell>
                  <Link
                    href={`/publications/${item.id}`}
                    target="_blank"
                    className="text-blue-500 hover:text-blue-600"
                  >
                    {item.id}
                  </Link>
                </TableCell>
                <TableCell>{item.name ?? "-"}</TableCell>
                <TableCell>{item.handle}</TableCell>
                <TableCell>
                  {item.onChainIdentity.ens?.name ? (
                    <Link
                      href={`https://app.ens.domains/name/${item.onChainIdentity.ens.name}`}
                      target="_blank"
                      className="text-blue-500 hover:text-blue-600"
                    >
                      {String(item.onChainIdentity.ens.name)}
                    </Link>
                  ) : (
                    "-"
                  )}
                </TableCell>
                <TableCell>{item.stats.totalFollowers}</TableCell>
                <TableCell>{item.stats.totalPublications}</TableCell>
              </TableRow>
            ))}
          </TableBody> */}
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
