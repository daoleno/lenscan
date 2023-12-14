import { searchParamsSchema } from "@/lib/validations/params"
import PublicationsTable from "@/components/publications-table"
import getPublications, {
  Publication,
} from "@/app/api/publications/getPublications"

export interface PublicationsProps {
  title?: string
  searchParams: {
    [key: string]: string | string[] | undefined
  }
  showToolbar?: boolean
  showPagination?: boolean
}
export default async function Publications({
  title = "Publications",
  searchParams,
  showToolbar,
  showPagination,
}: PublicationsProps) {
  const { page, per_page, sort, app, publication_type, is_momoka, profile_id } =
    searchParamsSchema.parse(searchParams)
  const pageAsNumber = Number(page)
  const fallbackPage =
    isNaN(pageAsNumber) || pageAsNumber < 1 ? 1 : pageAsNumber
  const perPageAsNumber = Number(per_page)
  const limit = isNaN(perPageAsNumber) ? 10 : perPageAsNumber
  const offset = fallbackPage > 0 ? (fallbackPage - 1) * limit : 0
  const [column, order] = (sort?.split(".") as [
    keyof Publication | undefined,
    "asc" | "desc" | undefined,
  ]) ?? ["block_timestamp", "desc"]
  const apps = app?.split(".") ?? []
  const publication_types = publication_type?.split(".") ?? []
  const { publications, totalCount } = await getPublications({
    limit,
    offset,
    sort: {
      column,
      order,
    },
    filter: {
      app: apps,
      publication_type: publication_types,
      is_momoka: is_momoka,
      profile_id: profile_id ? [profile_id] : undefined,
    },
  })
  const maxCount = totalCount ?? 1000000
  const pageCount = Math.ceil(Number(maxCount) / limit)
  return (
    <PublicationsTable
      title={title}
      data={publications}
      pageCount={pageCount}
      totalCount={maxCount}
      showToolbar={showToolbar}
      showPagination={showPagination}
    />
  )
}
