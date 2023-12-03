import { Publication } from "@/app/api/publications/publication"

import { columns } from "./data-table/columns"
import { DataTable } from "./data-table/data-table"

interface PublicationsTableProps {
  data: Publication[]
  pageCount: number
  totalCount?: number
  showToolbar?: boolean
  showPagination?: boolean
}

export default async function PublicationsTable({
  data,
  pageCount,
  totalCount,
  showToolbar,
  showPagination,
}: PublicationsTableProps) {
  return (
    <div className="flex-1 flex-col gap-3 md:flex">
      <h2 className="text-3xl font-bold tracking-tight">Publications</h2>
      <DataTable
        columns={columns}
        data={data}
        pageCount={pageCount}
        totalCount={totalCount}
        showToolbar={showToolbar}
        showPagination={showPagination}
        // Render dynamic faceted filters
        filterableColumns={[
          {
            id: "app",
            title: "App",
            options: [
              { label: "Hey", value: "hey" },
              { label: "Tape", value: "tape" },
              { label: "Phaver", value: "phaver" },
              { label: "Orb", value: "orb" },
            ],
          },
          {
            id: "publication_type",
            title: "Type",
            options: [
              { label: "Post", value: "POST" },
              { label: "Comment", value: "COMMENT" },
              { label: "Mirror", value: "MIRROR" },
            ],
          },
          {
            id: "is_momoka",
            title: "Network",
            options: [{ label: "Momoka", value: "false" }],
          },
        ]}
        // // Render dynamic searchable filters
        // searchableColumns={[
        //   {
        //     id: "title",
        //     title: "titles",
        //   },
        // ]}
      />
    </div>
  )
}
