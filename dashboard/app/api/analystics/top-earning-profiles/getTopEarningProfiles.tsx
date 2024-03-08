import db from "@/lib/db"
import lensClient from "@/lib/lensclient"
import { getIPFSURL } from "@/lib/utils"
import { sql } from "drizzle-orm"
import { DateRangeKey, getDateRangeAndCondition } from "../utils"

export type TopEarningProfile = {
	profile_id: string
	profile_handle: string
	profile_picture: any
	profile_link: string
	amount: number
}

export async function getTopEarningProfiles(
	symbol: string,
	rangeKey: DateRangeKey,
): Promise<TopEarningProfile[]> {
	const dateRange = getDateRangeAndCondition(rangeKey, "prr.block_timestamp")

	const statement = `
    SELECT
        prr.profile_id AS profile_id,
        SUM(CAST(prr.amount AS numeric) / POWER(10, ec.decimals)) AS amount
    FROM
        profile_revenue_record prr
    JOIN
        enabled_currency ec
    ON
        prr.currency = ec.currency
    WHERE
        ec.symbol = '${symbol}' ${dateRange}
    GROUP BY
        prr.profile_id
    ORDER BY
        "amount" DESC
    LIMIT 25;
    `

	const result = (await db.execute(sql.raw(statement))) as TopEarningProfile[]
	if (!result.length) {
		return []
	}

	// fetch profile pictures
	const profileIds = result.map((p) => p.profile_id)
	const fullProfiles = await lensClient.profile.fetchAll({
		where: {
			profileIds,
		},
	})

	for (const p of result) {
		const profile = fullProfiles.items.find(
			(profile) => profile.id === p.profile_id,
		)
		p.profile_handle = profile?.handle?.suggestedFormatted.localName || ""
		p.profile_link = `/profile/${p.profile_id}`
		p.profile_picture = getIPFSURL(profile?.metadata?.picture as any) || ""
	}

	return result as TopEarningProfile[]
}
