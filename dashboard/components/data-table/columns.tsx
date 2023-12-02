"use client"

import { ColumnDef } from "@tanstack/react-table"
import * as timeago from "timeago.js"

import { Publication } from "@/app/api/publications/publication"

import { DataTableColumnHeader } from "./data-table-column-header"

export const columns: ColumnDef<Publication>[] = [
  {
    accessorKey: "publication_id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Id" />
    ),
    cell: ({ row }) => <div>{row.getValue("publication_id")}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "profile_id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Profile Id" />
    ),
    cell: ({ row }) => <div>{row.getValue("profile_id")}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "is_momoka",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Network" />
    ),
    cell: ({ row }) => (
      <div>{row.getValue("is_momoka") ? "momoka" : "polygon"}</div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "app",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="App" />
    ),
    cell: ({ row }) => <div>{row.getValue("app")}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "publication_type",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Type" />
    ),
    cell: ({ row }) => <div>{row.getValue("publication_type")}</div>,
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
