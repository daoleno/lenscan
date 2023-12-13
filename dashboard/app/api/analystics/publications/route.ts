import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getDailyPublicationStats } from "./getDailyPublicationStats"

export const GET = (request: NextRequest) =>
  fetchData(getDailyPublicationStats, getRangeKey(request))
