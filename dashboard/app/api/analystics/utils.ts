import { NextRequest, NextResponse } from "next/server"

// Define a type for dateRange keys
export type DateRangeKey = "1D" | "1W" | "1M" | "3M" | "1Y" | "ALL"

export const dateRange = {
  "1D": 1,
  "1W": 7,
  "1M": 30,
  "3M": 90,
  "1Y": 365,
  ALL: 0,
}

export function getDateRangeCondition(
  rangeKey: DateRangeKey,
  filterName: string
): string {
  const range = dateRange[rangeKey]

  let dateCondition = ""
  if (range !== 0) {
    const currentDate = new Date()
    const pastDate = new Date(
      currentDate.setDate(currentDate.getDate() - range)
    )

    dateCondition = `WHERE ${filterName} >= '${pastDate.toISOString()}'`
  }

  return dateCondition
}

export function getRangeKey(request: NextRequest): string {
  let rangeKey = request.nextUrl.searchParams.get("range") as DateRangeKey
  return ["1D", "1W", "1M", "3M", "1Y", "ALL"].includes(rangeKey)
    ? rangeKey
    : "ALL"
}

export async function fetchData(getDataFunc: any, rangeKey: string) {
  try {
    const data = await getDataFunc(rangeKey)
    return NextResponse.json(data)
  } catch (error) {
    console.error(error)
    return NextResponse.json(
      {
        error: "Internal server error",
      },
      { status: 500 }
    )
  }
}
