import { RevalidateTime } from "@/config/ssr"
import { getAllUserActivity } from "@/app/api/analystics/user-activity/getUserActivity"

import AllUserActivityStats from "./all-user-activity-stats"

interface UserActivityProps {
  className?: string
}

export const revalidate = RevalidateTime

export default async function AllUserActivity({
  className,
}: UserActivityProps) {
  const data = await getAllUserActivity()

  return <AllUserActivityStats allStats={data} className={className} />
}
