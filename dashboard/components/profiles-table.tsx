import { Profile } from "@/app/api/profiles/getProfiles"

import { profileColumns } from "./data-table/columns"
import { DataTable } from "./data-table/data-table"

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
      <DataTable
        columns={profileColumns}
        data={data}
        pageCount={pageCount}
        totalCount={totalCount}
        showToolbar={showToolbar}
        showPagination={showPagination}
      />
    </div>
  )
}
