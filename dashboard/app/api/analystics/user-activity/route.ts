import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getUserActivity } from "./getUserActivity"

export const GET = (request: NextRequest) => {
  const profileId = request.nextUrl.searchParams.get("profile_id") || null
  return fetchData(getUserActivity, getRangeKey(request), profileId)
}
