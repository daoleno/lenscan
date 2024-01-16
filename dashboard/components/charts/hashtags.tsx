import { getAllPopularHashtags } from "@/app/api/analystics/popular-hashtags/getPopularHashtags"

import HashTagsStats from "./hashtags-stats"

export const revalidate = 60 * 60 * 5

export default async function HashTags() {
  const allStats = await getAllPopularHashtags()

  return <HashTagsStats allStats={allStats} />
}
