import { getAllAppUserStats } from "@/app/api/analystics/active-users/getActiveUserStats"

import ActiveUserStats from "./active-user-stats"

interface DAUStaticProps {
  className?: string
}

export const revalidate = 60 * 60 * 5

export default async function DAU({ className }: DAUStaticProps) {
  const allStats = await getAllAppUserStats("DAU")

  return (
    <ActiveUserStats
      title="Daily Active Users"
      allStats={allStats}
      className={className}
    />
  )
}
