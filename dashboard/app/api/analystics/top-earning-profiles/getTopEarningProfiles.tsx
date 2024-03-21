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
	const dateRange = getDateRangeAndCondition(rangeKey, "poamr.block_timestamp")

	const statement = `
        SELECT
            pr.profile_id AS profile_id,
            SUM(poamh.amount::numeric * act.count) / POWER(10, ec.decimals) AS amount
        FROM
            publication_record pr
        JOIN
            publication_open_action_module_history poamh
        ON
            pr.publication_id = poamh.publication_id
        JOIN
            enabled_currency ec
        ON
            poamh.currency = ec.currency
        JOIN (
            SELECT 
                poamh_inner.history_id, 
                COUNT(*) AS count
            FROM 
                publication_open_action_module_acted_record poamr
            JOIN 
                publication_open_action_module_history poamh_inner 
            ON 
                poamr.history_id = poamh_inner.history_id
            WHERE 
                poamh_inner.currency = (SELECT currency FROM enabled_currency WHERE symbol = '${symbol}')
				${dateRange}
            GROUP BY 
                poamh_inner.history_id
        ) act 
        ON 
            act.history_id = poamh.history_id
        WHERE
            ec.symbol = '${symbol}'
        GROUP BY
            pr.profile_id, ec.decimals
        ORDER BY
            amount DESC
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
