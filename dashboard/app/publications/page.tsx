import Publications from "@/components/publications"

interface PageProps {
  searchParams: {
    [key: string]: string | string[] | undefined
  }
}

export default async function Page({ searchParams }: PageProps) {
  return <Publications searchParams={searchParams} showToolbar showPagination />
}
