import { Suspense } from "react"

import { type Publication } from "@/app/api/publications/getPublications"

import { publicationColumns } from "./data-table/columns"
import { DataTable } from "./data-table/data-table"
import { DataTableLoading } from "./data-table/data-table-loading"

interface PublicationsTableProps {
  title: string
  data: Publication[]
  pageCount: number
  totalCount?: number
  showToolbar?: boolean
  showPagination?: boolean
}

export default async function PublicationsTable({
  title = "Publications",
  data,
  pageCount,
  totalCount,
  showToolbar,
  showPagination,
}: PublicationsTableProps) {
  return (
    <div className="my-3 flex-1 flex-col gap-3 md:flex">
      <h2 className="p-3 text-3xl font-bold tracking-tight">{title}</h2>
      <Suspense fallback={<DataTableLoading columnCount={6} rowCount={10} />}>
        <DataTable
          columns={publicationColumns}
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
                { label: "U3", value: "u3" },
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
              options: [{ label: "Momoka", value: "true" }],
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
      </Suspense>
    </div>
  )
}
