import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getUserActivityStats } from "./getUserActivity"

export const GET = (request: NextRequest) => {
  const profileId = request.nextUrl.searchParams.get("profile_id") || null
  return fetchData(getUserActivityStats, getRangeKey(request), profileId)
}
