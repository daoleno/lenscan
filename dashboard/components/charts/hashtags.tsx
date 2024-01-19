import { RevalidateTime } from "@/config/ssr"
import { getAllPopularHashtags } from "@/app/api/analystics/popular-hashtags/getPopularHashtags"

import HashTagsStats from "./hashtags-stats"

export const revalidate = RevalidateTime

export default async function HashTags() {
  const allStats = await getAllPopularHashtags()

  return <HashTagsStats allStats={allStats} />
}
