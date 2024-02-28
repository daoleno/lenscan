import db from "@/lib/db";
import lensClient from "@/lib/lensclient";
import { getIPFSURL } from "@/lib/utils";
import { sql } from "drizzle-orm";

export type Profile = {
	profile_id: string;
	owned_by: string;
	is_burnt: boolean;
	tx_hash: string;
	block_hash: string;
	block_number: number;
	log_index: number;
	tx_index: number;
	block_timestamp: Date;
	source_timestamp: bigint;

	profile_picture: string;
	profile_handle: string;
};

type getProfilesSort = {
	column?: string;
	order?: "ASC" | "DESC" | "asc" | "desc";
};

type getProfilesParams = {
	limit: number;
	offset: number;
	sort?: getProfilesSort;
};

export default async function getProfiles(
	params: getProfilesParams,
): Promise<{ totalCount: number; profiles: Profile[] }> {
	const { limit, offset, sort } = params;

	const sortOrder = sort ? `ORDER BY ${sort.column} ${sort.order}` : "";
	const statement = `SELECT * FROM profile_record ${sortOrder} LIMIT ${limit} OFFSET ${offset}`;
	const profiles = (await db.execute(sql.raw(statement))) as any[];

	const totalCount = (await db.execute(
		sql`SELECT COUNT(*) AS count FROM profile_record`,
	)) as any[];

	// fetch profile other metadata
	const profileIds = profiles.map((p) => p.profile_id);
	const fullProfiles = await lensClient.profile.fetchAll({
		where: {
			profileIds,
		},
	});

	// merge
	profiles.forEach((p) => {
		const fullProfile = fullProfiles.items.find((fp) => fp.id === p.profile_id);
		if (fullProfile) {
			p.profile_picture = getIPFSURL(fullProfile?.metadata?.picture as any);
			p.profile_handle = fullProfile?.handle?.suggestedFormatted.localName;
		}
	});

	return {
		totalCount: totalCount[0].count,
		profiles: profiles as Profile[],
	};
}
