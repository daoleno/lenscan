import { searchParamsSchema } from "@/lib/validations/params"
import DailyProfileStats from "@/components/daily-profile-stats"
import ProfilesTable from "@/components/profiles-table"

import getProfiles from "../api/profiles/getProfiles"
import { Publication } from "../api/publications/getPublications"

interface PageProps {
  searchParams: {
    [key: string]: string | string[] | undefined
  }
}

export default async function Page({ searchParams }: PageProps) {
  // Parse search params using zod schema
  const { page, per_page, sort } = searchParamsSchema.parse(searchParams)

  // Fallback page for invalid page numbers
  const pageAsNumber = Number(page)
  const fallbackPage =
    isNaN(pageAsNumber) || pageAsNumber < 1 ? 1 : pageAsNumber
  // Number of items per page
  const perPageAsNumber = Number(per_page)
  const limit = isNaN(perPageAsNumber) ? 10 : perPageAsNumber
  // Number of items to skip
  const offset = fallbackPage > 0 ? (fallbackPage - 1) * limit : 0
  // Column and order to sort by
  // Spliting the sort string by "." to get the column and order
  // Example: "title.desc" => ["title", "desc"]
  const [column, order] = (sort?.split(".") as [
    keyof Publication | undefined,
    "asc" | "desc" | undefined,
  ]) ?? ["block_timestamp", "desc"]

  const { profiles, totalCount } = await getProfiles({
    limit,
    offset,
    sort: {
      column,
      order,
    },
  })
  const pageCount = Math.ceil(Number(totalCount) / limit)

  return (
    <div className="flex flex-col gap-3 p-8">
      <h2 className="text-3xl font-bold tracking-tight">Profiles</h2>
      <DailyProfileStats enableSpark={false} />
      <ProfilesTable
        data={profiles}
        pageCount={pageCount}
        totalCount={totalCount}
        showToolbar={false}
        showPagination={true}
      />
    </div>
  )
}
