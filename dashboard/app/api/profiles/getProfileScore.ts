import db from "@/lib/db";
import { sql } from "drizzle-orm";

export type ProfileScore = {
	profile_id: string;
	score: number;
	generated_at: Date;
};

export default async function getProfileScore(
	profileId: string,
): Promise<ProfileScore> {
	const res = (await db.execute(
		sql`SELECT profile_id, score, generated_at FROM machine_learning_quality_profiles WHERE profile_id = ${profileId} ORDER BY generated_at DESC LIMIT 1`,
	)) as ProfileScore[];

	if (res.length === 0) {
		throw new Error("Profile not found");
	}

	return res[0];
}
