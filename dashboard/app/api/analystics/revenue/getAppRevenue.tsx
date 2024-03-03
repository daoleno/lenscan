import "server-only";

import db from "@/lib/db";
import { sql } from "drizzle-orm";
import { formatEther } from "viem";
import { permissionlessCreators } from "../../revenue/getRevenue";

export type AppRevenueStats = {
	apps: string[];
	stats: {
		day: string;
		[key: string]: number | string;
	}[];
};

export async function getDailyAppRevenueStats(): Promise<AppRevenueStats> {
	const statement = `
        SELECT 
          DATE_TRUNC('day', block_timestamp)::date as day, 
          to_address,
          SUM(value) AS revenue
        FROM lens_profile_created_transaction
        WHERE value > 0
        GROUP BY day, to_address
        ORDER BY day ASC;
    `;

	const results = (await db.execute(sql.raw(statement))) as any[];

	const statsMap: { [day: string]: { [appName: string]: number } } = {};
	const apps: Set<string> = new Set();

	for (const result of results) {
		const day = new Date(result.day).toLocaleDateString();
		const appName = permissionlessCreators[result.to_address] || "others";

		if (!statsMap[day]) {
			statsMap[day] = {};
		}

		statsMap[day][appName] = Number(formatEther(result.revenue));
		apps.add(appName);
	}

	const stats: AppRevenueStats["stats"] = Object.entries(statsMap).map(
		([day, appRevenues]) => ({
			day: day,
			...appRevenues,
		}),
	);

	return {
		apps: Array.from(apps),
		stats,
	};
}

type TotalAppRevenue = {
	app: string;
	revenue: number;
};

export async function getTotalAppRevenue(): Promise<TotalAppRevenue[]> {
	const statement = `
		SELECT 
		  SUM(value) AS revenue,
		  to_address
		FROM lens_profile_created_transaction
		WHERE value > 0
		GROUP BY to_address
		ORDER BY revenue DESC;
	`;

	const results = (await db.execute(sql.raw(statement))) as any[];

	const totalAppRevenue: TotalAppRevenue[] = results.map((result) => ({
		app: permissionlessCreators[result.to_address] || "others",
		revenue: Number(formatEther(result.revenue)),
	}));

	return totalAppRevenue;
}
