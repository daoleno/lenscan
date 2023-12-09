import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getPopularHashtags } from "./getPopularHashtags"

export const GET = (request: NextRequest) =>
  fetchData(getPopularHashtags, getRangeKey(request))
