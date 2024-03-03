import db from "@/lib/db";
import { sql } from "drizzle-orm";

export const permissionlessCreators: Record<string, string> = {
	"0x0b5e6100243f793e480DE6088dE6bA70aA9f3872": "Lens Protocol",
	"0x4b8845ACb8148dE64D1D99Cf27A3890a91F55E53": "Hey",
	"0xac79e9ef69021f9bf7ccb175611f3115ff65a44d": "Orb",
	"0xD0f6d9676d36F5f4AF5765fCb78c388B51577327": "Tape",
	"0x57B7bf6f792a6181Ec5aFB88cE7bcE330a9d1b67": "Phaver",
	"0xE7404472E8C7cEaf15c1cA64d4fe3eA78818fF5a": "Buttrfly",
};

export type Revenue = {
	id: number;
	block_number: number;
	block_timestamp: Date;
	transaction_hash: string;
	to_address: string;
	app: string;
	value: number;
};

type getRevenueSort = {
	column?: string;
	order?: "ASC" | "DESC" | "asc" | "desc";
};

type getRevenueFilter = {
	app?: string[];
};

type getRevenueParams = {
	limit: number;
	offset: number;
	sort?: getRevenueSort;
	filter?: getRevenueFilter;
};

export default async function getRevenue(params: getRevenueParams) {
	const { limit, offset, sort, filter } = params;

	console.log("filter", filter);

	const sortOrder = sort ? `ORDER BY ${sort.column} ${sort.order}` : "";

	const conditions = [];
	if (filter) {
		for (const key in filter) {
			const value = filter[key as keyof getRevenueFilter];

			// Special handling for the 'app' filter
			if (key === "app" && Array.isArray(value) && value.length > 0) {
				// Map 'app' values to corresponding 'to_address' values
				const toAddresses = Object.entries(permissionlessCreators)
					.filter(([address, app]) => value.includes(app))
					.map(([address]) => address);

				// Ensure we have corresponding 'to_address' values to filter by
				if (toAddresses.length > 0) {
					conditions.push(`to_address IN ('${toAddresses.join("','")}')`);
				}
			} else if (value !== undefined && value.length > 0) {
				if (Array.isArray(value)) {
					conditions.push(`${key} IN ('${value.join("','")}')`);
				} else {
					conditions.push(`${key}='${value}'`);
				}
			}
		}
	}

	const filterCondition = conditions.length
		? " WHERE " + conditions.join(" AND ") + " AND value > 0"
		: " WHERE value > 0";

	const statement = `SELECT id, block_number, block_timestamp, transaction_hash, to_address, value FROM lens_profile_created_transaction ${filterCondition} ${sortOrder} LIMIT ${limit} OFFSET ${offset}`;

	const revenueData = (await db.execute(sql.raw(statement))) as any[];

	const totalCount = (await db.execute(
		sql`SELECT COUNT(*) AS count FROM lens_profile_created_transaction WHERE value > 0`,
	)) as any[];

	const revenue = revenueData.map((result) => {
		const app = permissionlessCreators[result.to_address] || "others";
		return {
			...result,
			app,
		};
	});

	return {
		totalCount: totalCount[0].count,
		revenue: revenue,
	};
}
