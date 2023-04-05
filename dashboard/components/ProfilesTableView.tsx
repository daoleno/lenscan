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

import { ProfileFragment, useExploreProfiles } from "@lens-protocol/react-web";

export default function ProfilesTableView() {
  const { data: profiles, loading, error } = useExploreProfiles({ limit: 30 });

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  if (!profiles) {
    return <div>No profiles found</div>;
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

          <TableBody>
            {profiles.map((item: ProfileFragment) => (
              <TableRow key={item.id}>
                <TableCell>
                  <Link
                    href={`/profiles/${item.id}`}
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
          </TableBody>
        </Table>
      </Card>
    </div>
  );
}
