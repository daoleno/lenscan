import { RevalidateTime } from "@/config/ssr"
import { getAllAppUserStats } from "@/app/api/analystics/active-users/getActiveUserStats"

import ActiveUserStats from "./active-user-stats"

interface DAUStaticProps {
  className?: string
}

export const revalidate = RevalidateTime

export default async function MAU({ className }: DAUStaticProps) {
  const allStats = await getAllAppUserStats("MAU")

  return (
    <ActiveUserStats
      title="Monthly Active Users"
      allStats={allStats}
      className={className}
    />
  )
}
