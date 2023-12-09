import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getUserActivities } from "./getUserActivities"

export const GET = (request: NextRequest) =>
  fetchData(getUserActivities, getRangeKey(request))
