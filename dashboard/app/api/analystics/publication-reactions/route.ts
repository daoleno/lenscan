import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getPublicationReactionsOvertime } from "./getPublicationReactionsOverTime"

export const GET = (request: NextRequest) =>
  fetchData(getPublicationReactionsOvertime, getRangeKey(request))
