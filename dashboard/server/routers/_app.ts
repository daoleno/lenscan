import { publicProcedure, router } from "../trpc";
import { eventRouter } from "./event";

export const appRouter = router({
  healthcheck: publicProcedure.query(() => "yay!"),
  event: eventRouter,
});

// export type definition of API
export type AppRouter = typeof appRouter;
