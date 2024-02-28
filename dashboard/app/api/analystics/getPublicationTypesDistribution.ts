import db from "@/lib/db";
import { sql } from "drizzle-orm";

export type PublicationTypesDistribution = {
	publication_type: string;
	count: number;
};

export async function getPublicationTypesDistribution() {
	const statement = `
    SELECT 
    publication_type,
    COUNT(*) AS count
    FROM 
    publication_record
    GROUP BY 
    publication_type
    ORDER BY 
    count DESC;
  `;

	const pubs = (await db.execute(
		sql.raw(statement),
	)) as PublicationTypesDistribution[];

	return pubs.map((a) => ({
		...a,
		count: Number(a.count),
	}));
}
