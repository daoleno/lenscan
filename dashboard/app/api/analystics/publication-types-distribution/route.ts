import { NextRequest } from "next/server"

import { fetchData, getRangeKey } from "../utils"
import { getPublicationTypesDistribution } from "./getPublicationTypesDistribution"

export const GET = (request: NextRequest) =>
  fetchData(getPublicationTypesDistribution, getRangeKey(request))
