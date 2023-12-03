import Publication from "@/components/publication"

interface PageProps {
  params: {
    id: string
  }
}
export default function Page({ params }: PageProps) {
  return <Publication id={params.id} />
}
