import { drizzle } from "drizzle-orm/postgres-js";
import postgres from "postgres";

// class MyLogger implements Logger {
// 	logQuery(query: string, params: unknown[]): void {
// 		console.log({ query, params });
// 	}
// }

const queryClient = postgres(
	process.env.DATABASE_URL ||
		"postgres://postgres:postgres@localhost:5432/lens_v2",
	{
		max: 10,
	},
);

let db: ReturnType<typeof drizzle>;
if (process.env.NODE_ENV === "production") {
	db = drizzle(
		queryClient,
		// { logger: new MyLogger() }
	);
} else {
	if (!(global as any).db) {
		(global as any).db = drizzle(
			queryClient,
			// { logger: new MyLogger() }
		);
	}
	db = (global as any).db;
}

export default db;
