"use client"

import Image from "next/image"
import Link from "next/link"
import {
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@radix-ui/react-tooltip"
import { ColumnDef } from "@tanstack/react-table"
import * as timeago from "timeago.js"

import { Profile } from "@/app/api/profiles/profile"
import { type Publication } from "@/app/api/publications/getPublications"

import { Tooltip } from "../ui/tooltip"
import { DataTableColumnHeader } from "./data-table-column-header"

export const publicationColumns: ColumnDef<Publication>[] = [
  {
    accessorKey: "publication_id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Id" />
    ),
    cell: ({ row }) => (
      <Link
        className="underline"
        href={`/publication/${row.original.publication_id}`}
      >
        {row.original.publication_id}
      </Link>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "profile_id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Profile" />
    ),
    cell: ({ row }) => (
      <Link
        className="flex items-center gap-2"
        href={`/profile/${row.original.profile_id}`}
      >
        <img
          src={row.original.profile_picture || "/images/default-profile.png"}
          alt={row.original.profile_handle}
          className="h-8 w-8 rounded-full object-cover"
        />
        <div className="underline">{row.original.profile_handle}</div>
      </Link>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "is_momoka",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Network" />
    ),
    cell: ({ row }) => (
      <div className="capitalize">
        {row.getValue("is_momoka") ? "momoka" : "polygon"}
      </div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "app",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="App" />
    ),
    cell: ({ row }) => (
      <div className="flex items-center gap-1">
        {[
          "hey",
          "tape",
          "phaver",
          "orb",
          "t2",
          "buttrfly",
          "dumpling",
        ].includes(row.original.app) ? (
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger>
                <Image
                  src={`/apps/${row.original.app.toLowerCase()}.png`}
                  width={28}
                  height={28}
                  alt={row.original.app}
                  className="rounded-full object-cover"
                />
              </TooltipTrigger>
              <TooltipContent>{row.original.app}</TooltipContent>
            </Tooltip>
          </TooltipProvider>
        ) : (
          <div className="capitalize">{row.original.app}</div>
        )}
      </div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "publication_type",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Type" />
    ),
    cell: ({ row }) => (
      <div className="capitalize">
        {row.original.publication_type.toLowerCase()}
      </div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "block_number",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Block" />
    ),
    cell: ({ row }) => <div>{row.getValue("block_number")}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "block_timestamp",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Age" />
    ),
    cell: ({ row }) => (
      <div>{timeago.format(row.getValue("block_timestamp"))}</div>
    ),
  },
]

export const profileColumns: ColumnDef<Profile>[] = [
  {
    accessorKey: "profile_id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Id" />
    ),
    cell: ({ row }) => (
      <Link className="underline" href={`/profile/${row.original.profile_id}`}>
        {row.original.profile_id}
      </Link>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "owned_by",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Owner" />
    ),
    cell: ({ row }) => (
      <Link
        className="underline"
        href={`https://polygonscan.com/address/${row.original.owned_by}`}
      >
        {row.original.owned_by}
      </Link>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "tx_hash",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Tx Hash" />
    ),
    cell: ({ row }) => (
      <Link
        href={`https://polygonscan.com/tx/${row.original.tx_hash}`}
        className="underline"
      >
        {row.original.tx_hash}
      </Link>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "block_number",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Block" />
    ),
    cell: ({ row }) => <div>{row.getValue("block_number")}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "block_timestamp",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Age" />
    ),
    cell: ({ row }) => (
      <div>{timeago.format(row.getValue("block_timestamp"))}</div>
    ),
  },
]
