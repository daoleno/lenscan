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
	filterName: string,
): string {
	const range = dateRange[rangeKey]

	let dateCondition = ""
	if (range !== 0) {
		const currentDate = new Date()
		const pastDate = new Date(
			currentDate.setDate(currentDate.getDate() - range),
		)

		dateCondition = `WHERE ${filterName} >= '${pastDate.toISOString()}'`
	}

	return dateCondition
}

export function getDateRangeAndCondition(
	rangeKey: DateRangeKey,
	filterName: string,
): string {
	const range = dateRange[rangeKey]

	let dateCondition = ""
	if (range !== 0) {
		const currentDate = new Date()
		const pastDate = new Date(
			currentDate.setDate(currentDate.getDate() - range),
		)

		dateCondition = `AND ${filterName} >= '${pastDate.toISOString()}'`
	}

	return dateCondition
}

export function getPreviousDateRangeCondition(
	rangeKey: DateRangeKey,
	filterName: string,
): string {
	const range = dateRange[rangeKey]

	let previousDateCondition = ""
	if (range !== 0) {
		// Get the current date and set the time to the start of the day (midnight)
		const endDate = new Date()
		endDate.setHours(0, 0, 0, 0)

		// Calculate the start date of the current period
		const startDateCurrentPeriod = new Date(endDate.getTime())
		startDateCurrentPeriod.setDate(endDate.getDate() - range)

		// Calculate the start date of the previous period
		const startDatePreviousPeriod = new Date(startDateCurrentPeriod.getTime())
		startDatePreviousPeriod.setDate(startDateCurrentPeriod.getDate() - range)

		previousDateCondition = `WHERE ${filterName} >= '${startDatePreviousPeriod.toISOString()}' AND ${filterName} < '${startDateCurrentPeriod.toISOString()}'`
	}

	return previousDateCondition
}

export function getRangeKey(request: NextRequest): DateRangeKey {
	const rangeKey = request.nextUrl.searchParams.get("range") as DateRangeKey
	return ["1D", "1W", "1M", "3M", "1Y", "ALL"].includes(rangeKey)
		? rangeKey
		: "ALL"
}

export async function fetchData(
	getDataFunc: any,
	rangeKey: string,
	profileId: string | null = null,
) {
	try {
		const data = await getDataFunc(rangeKey, profileId)
		return NextResponse.json(data)
	} catch (error) {
		console.error(error)
		return NextResponse.json(
			{
				error: "Internal server error",
			},
			{ status: 500 },
		)
	}
}
