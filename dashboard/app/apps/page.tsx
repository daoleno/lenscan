import AppsSummary from "@/components/apps-summary"
import TopApps from "@/components/charts/top-apps"

export const revalidate = 3600

export default function Page() {
  return (
    <div className="flex flex-col gap-3 p-8">
      <AppsSummary />
      <TopApps />
    </div>
  )
}
