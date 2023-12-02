import duckdb from "@/lib/duckdb";
import { GlobalStats } from "./globalStats";

import 'server-only';


export async function getGlobalStats() {
        const stats = await duckdb.all(`SELECT 
                (SELECT COUNT(DISTINCT profile_id) FROM profile_record) AS totalProfiles,
                (SELECT COUNT(DISTINCT publication_id) FROM publication_record) AS totalPublications,
                (SELECT COUNT(*) FROM publication_record WHERE publication_type='POST') AS totalPosts,
                (SELECT COUNT(*) FROM publication_record WHERE publication_type='MIRROR') AS totalMirrors,
                (SELECT COUNT(*) FROM publication_record WHERE publication_type='COMMENT') AS totalComments;`)
        return stats[0] as GlobalStats
}