import { NextRequest, NextResponse } from "next/server"
import { getRangeKey } from "../utils"
import { getTopEarningProfiles } from "./getTopEarningProfiles"

export const GET = async (request: NextRequest) => {
	try {
		const rangeKey = getRangeKey(request)
		const symbol = request.nextUrl.searchParams.get("symbol") as string
		const data = await getTopEarningProfiles(symbol, rangeKey)
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
