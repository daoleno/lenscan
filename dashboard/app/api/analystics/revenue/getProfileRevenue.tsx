import "server-only"

import db from "@/lib/db"
import lensClient from "@/lib/lensclient"
import { getIPFSURL } from "@/lib/utils"
import { sql } from "drizzle-orm"

export type ProfileRevenue = {
	profile_id: string
	currency: string
	total_revenue: bigint
	currency_name: string
	currency_symbol: string
	decimals: number
}

export async function getProfileRevenue(
	profileId: string,
): Promise<ProfileRevenue[]> {
	const statement = `
		SELECT 
            pr.profile_id,
            poamh.currency, 
            SUM(COALESCE(NULLIF(poamh.amount, '')::numeric, 0) * act.count) AS total_revenue,
            ec.name AS currency_name, 
            ec.symbol AS currency_symbol, 
            ec.decimals
        FROM publication_record pr
        JOIN publication_open_action_module_history poamh ON pr.publication_id = poamh.publication_id
        JOIN (
            SELECT 
                poamh_inner.history_id, 
                COUNT(*) AS count
            FROM publication_open_action_module_acted_record poamr_inner
            JOIN publication_open_action_module_history poamh_inner ON poamr_inner.history_id = poamh_inner.history_id
            JOIN publication_record pr_inner ON poamh_inner.publication_id = pr_inner.publication_id
            WHERE pr_inner.profile_id = '${profileId}'
            GROUP BY poamh_inner.history_id
        ) act ON act.history_id = poamh.history_id
        JOIN enabled_currency ec ON poamh.currency = ec.currency
        WHERE pr.profile_id = '${profileId}'
        GROUP BY pr.profile_id, poamh.currency, ec.name, ec.symbol, ec.decimals
        ORDER BY total_revenue DESC
	`

	return await db.execute(sql.raw(statement))
}

type getRevenueRecordsSort = {
	column?: string
	order?: "ASC" | "DESC" | "asc" | "desc"
}

type getRevenueRecordsFilter = {
	profile_id?: string[]
	symbol?: string[] // Changed from currency to symbol
}

type getRevenueRecordsParams = {
	limit: number
	offset: number
	sort?: getRevenueRecordsSort
	filter?: getRevenueRecordsFilter
}

export type RevenueRecord = {
	profile_id: string
	publication_id: string
	profile_handle: string
	profile_picture: string
	currency: string
	amount: bigint
	tx_hash: string
	block_timestamp: Date
	currency_name: string
	currency_symbol: string
	currency_decimals: number
}

export async function getRevenueRecords(
	params: getRevenueRecordsParams,
): Promise<{
	totalCount: number | undefined
	revenueRecords: RevenueRecord[]
}> {
	const { limit, offset, sort, filter } = params

	const sortOrder = sort ? `ORDER BY ${sort.column} ${sort.order}` : ""

	const conditions = ["COALESCE(amount::numeric, 0) > 0"]
	if (filter) {
		if (filter.profile_id && filter.profile_id.length > 0) {
			conditions.push(`pr.profile_id IN ('${filter.profile_id.join("','")}')`)
		}

		if (filter.symbol && filter.symbol.length > 0) {
			conditions.push(`ec.symbol IN ('${filter.symbol.join("','")}')`)
		}
	}

	const filterCondition = conditions.length
		? ` WHERE ${conditions.join(" AND ")}`
		: ""

	const statement = `
        SELECT pr.profile_id, pr.publication_id, poamh.currency, poamr.tx_hash, poamr.block_hash, 
               poamr.block_number, poamr.log_index, poamr.tx_index, poamr.block_timestamp, 
               poamh.amount, ec.name AS currency_name, ec.symbol AS currency_symbol, ec.decimals AS currency_decimals
        FROM publication_open_action_module_history poamh
        JOIN publication_record pr ON poamh.publication_id = pr.publication_id
        JOIN publication_open_action_module_acted_record poamr ON poamr.history_id = poamh.history_id
        LEFT JOIN enabled_currency ec ON poamh.currency = ec.currency
        ${filterCondition}
        ${sortOrder} 
        LIMIT ${limit} OFFSET ${offset}
    `
	const revenueRecords = (await db.execute(sql.raw(statement))) as any[]

	if (!revenueRecords.length) {
		return {
			totalCount: 0,
			revenueRecords: [],
		}
	}

	// fetch profile pictures
	const profileIds = revenueRecords.map((p) => p.profile_id)
	const fullProfiles = await lensClient.profile.fetchAll({
		where: {
			profileIds,
		},
	})

	for (const p of revenueRecords) {
		const profile = fullProfiles.items.find(
			(profile) => profile.id === p.profile_id,
		)
		p.profile_handle = profile?.handle?.suggestedFormatted.localName || ""
		p.profile_picture = getIPFSURL(profile?.metadata?.picture as any) || ""
	}

	// Get total count if we have profile_id filter
	const countStatement = `
		SELECT COUNT(*) AS count
		FROM publication_open_action_module_acted_record poamr
		JOIN publication_open_action_module_history poamh ON poamr.history_id = poamh.history_id
		JOIN publication_record pr ON poamh.publication_id = pr.publication_id
		LEFT JOIN enabled_currency ec ON poamh.currency = ec.currency
		${filterCondition}
    `

	const result = (await db.execute(sql.raw(countStatement))) as any[]
	const totalCount = result[0].count

	return {
		totalCount,
		revenueRecords: revenueRecords as RevenueRecord[],
	}
}

export type TotalProfileRevenue = {
	token_symbol: string
	token_name: string
	token_decimals: number
	token_address: string
	total_revenue: bigint
}

export async function getTotalProfileRevenueBySymbol(
	symbol: string,
): Promise<TotalProfileRevenue> {
	const statement = `
		SELECT
            ec.symbol AS token_symbol,
            ec.name AS token_name,
            ec.decimals AS token_decimals,
            ec.currency AS token_address,
            SUM(poamh.amount::numeric * act.count) AS total_revenue
        FROM
            enabled_currency ec
        JOIN
            publication_open_action_module_history poamh
        ON
            ec.currency = poamh.currency
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
            GROUP BY 
                poamh_inner.history_id
        ) act 
        ON 
            act.history_id = poamh.history_id
        WHERE
            ec.symbol = '${symbol}'
        GROUP BY
            ec.symbol, ec.name, ec.decimals, ec.currency
	`

	const result = await db.execute(sql.raw(statement))

	return result[0] as TotalProfileRevenue
}

export type TotalDailyRevenue = {
	token_symbol: string
	day: string
	amount: number
}

export async function getTotalDailyRevenueBySymbol(
	symbol: string,
): Promise<TotalDailyRevenue[]> {
	const statement = `
		SELECT
            ec.symbol AS token_symbol,
            DATE_TRUNC('day', poamh.block_timestamp) AS day,
            SUM(poamh.amount::numeric * act.count) / POWER(10, MAX(ec.decimals)) AS amount
        FROM
            publication_open_action_module_history poamh
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
            JOIN
                enabled_currency ec_inner
            ON
                poamh_inner.currency = ec_inner.currency
            WHERE 
                ec_inner.symbol = '${symbol}'
            GROUP BY 
                poamh_inner.history_id
        ) act 
        ON 
            act.history_id = poamh.history_id
        WHERE
            ec.symbol = '${symbol}'
        GROUP BY
            day, ec.symbol
        ORDER BY
            day ASC;
	`

	const result = (await db.execute(sql.raw(statement))) as TotalDailyRevenue[]

	const chartData = result.map((a) => ({
		day: new Date(a.day).toLocaleDateString(),
		amount: Number(a.amount),
		token_symbol: symbol,
	}))

	return chartData as TotalDailyRevenue[]
}
