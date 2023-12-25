import { Suspense } from "react"

import { Profile } from "@/app/api/profiles/getProfiles"

import { profileColumns } from "./data-table/columns"
import { DataTable } from "./data-table/data-table"
import { DataTableLoading } from "./data-table/data-table-loading"

interface ProfilesTableProps {
  data: Profile[]
  pageCount: number
  totalCount?: number
  showToolbar?: boolean
  showPagination?: boolean
}

export default async function ProfilesTable({
  data,
  pageCount,
  totalCount,
  showToolbar,
  showPagination,
}: ProfilesTableProps) {
  return (
    <div className="my-3 flex-1 flex-col gap-3 md:flex">
      <Suspense fallback={<DataTableLoading columnCount={6} rowCount={10} />}>
        <DataTable
          columns={profileColumns}
          data={data}
          pageCount={pageCount}
          totalCount={totalCount}
          showToolbar={showToolbar}
          showPagination={showPagination}
        />
      </Suspense>
    </div>
  )
}
