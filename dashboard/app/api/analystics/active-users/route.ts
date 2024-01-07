import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getAppUserStats } from "./getDailyActiveUser"

export const GET = (request: NextRequest) => {
  const statType = request.nextUrl.searchParams.get("statType")
  return fetchData(getAppUserStats, getRangeKey(request), statType)
}
