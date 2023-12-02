import { searchParamsSchema } from "@/lib/validations/params"
import PublicationsTable from "@/components/publications-table"

import getPublications from "../api/publications/getPublications"
import { Publication } from "../api/publications/publication"

interface PageProps {
  searchParams: {
    [key: string]: string | string[] | undefined
  }
}

export default async function Page({ searchParams }: PageProps) {
  // Parse search params using zod schema
  const { page, per_page, sort, app, publication_type, network } =
    searchParamsSchema.parse(searchParams)

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

  const apps = app?.split(".") ?? []
  const publication_types = publication_type?.split(".") ?? []
  const networks = network?.split(".") ?? []

  const { publications } = await getPublications({
    limit,
    offset,
    sort: {
      column,
      order,
    },
    filter: {
      app: apps,
      publication_type: publication_types,
      network: networks,
    },
  })
  const maxCount = 1000000
  const pageCount = Math.ceil(Number(maxCount) / limit)

  return (
    <PublicationsTable
      data={publications}
      pageCount={pageCount}
      totalCount={maxCount}
      showToolbar={true}
      showPagination={true}
    />
  )
}
