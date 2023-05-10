import { publicProcedure, router } from "../trpc";
import { eventRouter } from "./event";
import { momokaRouter } from "./momoka";

export const appRouter = router({
  healthcheck: publicProcedure.query(() => "yay!"),
  event: eventRouter,
  momoka: momokaRouter,
});

// export type definition of API
export type AppRouter = typeof appRouter;
