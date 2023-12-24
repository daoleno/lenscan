import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getTopApps } from "./getAppStats"

export const GET = (request: NextRequest) => {
  return fetchData(getTopApps, getRangeKey(request))
}
