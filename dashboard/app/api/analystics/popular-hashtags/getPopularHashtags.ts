import db from "@/lib/db";
import { DateRangeKey, getDateRangeCondition } from "../utils";

import { sql } from "drizzle-orm";
import "server-only";

export type Hashtag = {
	hashtag: string;
	count: number;
};

export async function getAllPopularHashtags(): Promise<{
	[key in DateRangeKey]?: Hashtag[];
}> {
	const rangeKeys: DateRangeKey[] = ["1D", "1W", "1M", "3M", "1Y", "ALL"];
	const allPopularHashtags: { [key in DateRangeKey]?: Hashtag[] } = {};

	for (const rangeKey of rangeKeys) {
		allPopularHashtags[rangeKey] = await getPopularHashtags(rangeKey);
	}

	return allPopularHashtags;
}

export async function getPopularHashtags(rangeKey: DateRangeKey = "ALL") {
	let statement = `
    SELECT hashtag, COUNT(*) AS count
    FROM publication_hashtag
  `;

	statement += getDateRangeCondition(rangeKey, "timestamp");

	statement += `
    GROUP BY hashtag
    ORDER BY count DESC
    LIMIT 20
  `;

	const hashtags = (await db.execute(sql.raw(statement))) as Hashtag[];
	return hashtags.map((a) => ({
		...a,
		count: Number(a.count),
	}));
}
