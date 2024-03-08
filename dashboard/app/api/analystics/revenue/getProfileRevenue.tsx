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
        prr.profile_id, 
        prr.currency, 
        SUM(CAST(prr.amount AS numeric)) AS total_revenue, 
        ec.name AS currency_name, 
        ec.symbol AS currency_symbol, 
        ec.decimals
    FROM profile_revenue_record prr
    JOIN enabled_currency ec ON prr.currency = ec.currency
    WHERE prr.profile_id = '${profileId}'
    GROUP BY prr.profile_id, prr.currency, ec.name, ec.symbol, ec.decimals
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

	const conditions = []
	if (filter) {
		if (filter.profile_id && filter.profile_id.length > 0) {
			conditions.push(`prr.profile_id IN ('${filter.profile_id.join("','")}')`)
		}

		if (filter.symbol && filter.symbol.length > 0) {
			// Filter based on currency symbol
			conditions.push(`ec.symbol IN ('${filter.symbol.join("','")}')`)
		}
	}

	const filterCondition = conditions.length
		? " WHERE " + conditions.join(" AND ")
		: ""
	const statement = `
        SELECT prr.*, ec.name AS currency_name, ec.symbol AS currency_symbol, ec.decimals AS currency_decimals
        FROM profile_revenue_record prr
        JOIN enabled_currency ec ON prr.currency = ec.currency
        ${filterCondition} ${sortOrder} LIMIT ${limit} OFFSET ${offset}
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
	const countStatement = `SELECT COUNT(*) AS count FROM profile_revenue_record prr JOIN enabled_currency ec ON prr.currency = ec.currency ${filterCondition}`
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
		SUM(CAST(prr.amount AS numeric)) AS total_revenue
	FROM
		profile_revenue_record prr
	JOIN
		enabled_currency ec
	ON
		prr.currency = ec.currency
	WHERE
		ec.symbol = '${symbol}'
	GROUP BY
		ec.symbol, ec.name, ec.decimals, ec.currency; 
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
		DATE_TRUNC('day', prr.block_timestamp) AS day,
		SUM(CAST(prr.amount AS numeric) / POWER(10, ec.decimals)) AS amount
	FROM
		profile_revenue_record prr
	JOIN
		enabled_currency ec
	ON
		prr.currency = ec.currency
	WHERE
		ec.symbol = '${symbol}'
	GROUP BY
		day, ec.symbol
	ORDER BY
		day ASC, ec.symbol ASC;
	`

	const result = (await db.execute(sql.raw(statement))) as TotalDailyRevenue[]

	const chartData = result.map((a) => ({
		day: new Date(a.day).toLocaleDateString(),
		amount: Number(a.amount),
		token_symbol: symbol,
	}))

	return chartData as TotalDailyRevenue[]
}
