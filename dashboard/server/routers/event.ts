import { z } from "zod";
import prisma from "../prisma";
import { publicProcedure, router } from "../trpc";

export const eventRouter = router({
  getEvents: publicProcedure
    .input(
      z.object({
        take: z.number(),
        cursor: z.number().nullable(),
      })
    )
    .query(async ({ input }) => {
      const res = (await prisma.$queryRaw`
        SELECT reltuples::bigint AS estimate
        FROM pg_class
        WHERE relname = 'Event';
      `) as { estimate: string }[];
      const count = Number(res[0].estimate);

      if (!input.cursor) {
        const firstQueryEvents = await prisma.event.findMany({
          take: input.take,
          orderBy: { id: "desc" },
        });
        const firstQueryLastEvent =
          firstQueryEvents[firstQueryEvents.length - 1];
        const firstQueryNextCursor = firstQueryLastEvent
          ? firstQueryLastEvent.id
          : null;
        return {
          count,
          events: firstQueryEvents,
          nextCursor: firstQueryNextCursor,
        };
      }

      const events = await prisma.event.findMany({
        take: input.take,
        skip: 1,
        cursor: {
          id: input.cursor,
        },
        orderBy: {
          id: "desc",
        },
      });
      const lastEvent = events[events.length - 1];
      const nextCursor = lastEvent ? lastEvent.id : null;
      return {
        count,
        events,
        nextCursor,
      };
    }),
  getEvent: publicProcedure.input(z.number()).query(async ({ input }) => {
    const event = await prisma.event.findUnique({
      where: {
        id: input,
      },
    });
    return event;
  }),
  getEventsCountByProfileId: publicProcedure
    .input(z.number())
    .query(async ({ input }) => {
      const res = (await prisma.$queryRawUnsafe(
        `SELECT COUNT(*) FROM "Event" WHERE data->'ProfileId' = '${input}'`
      )) as any[];
      const count = Number(res[0].count);
      return count;
    }),
  getEventsByProfileId: publicProcedure
    .input(
      z.object({
        profileId: z.number(),
        take: z.number(),
        cursor: z.number().nullable(),
      })
    )
    .query(async ({ input }) => {
      if (!input.cursor) {
        const firstQueryEvents = (await prisma.$queryRawUnsafe(
          `SELECT * FROM "Event" WHERE data->'ProfileId' = '${input.profileId}' ORDER BY timestamp DESC LIMIT ${input.take}`
        )) as any[];
        const firstQueryLastEvent =
          firstQueryEvents[firstQueryEvents.length - 1];
        const firstQueryNextCursor = firstQueryLastEvent
          ? firstQueryLastEvent.id
          : null;
        return {
          events: firstQueryEvents,
          nextCursor: firstQueryNextCursor,
        };
      }

      const events = (await prisma.$queryRawUnsafe(
        `SELECT * FROM "Event" WHERE data->'ProfileId' = '${input.profileId}' AND id < ${input.cursor} ORDER BY timestamp DESC LIMIT ${input.take}`
      )) as any[];
      const lastEvent = events[events.length - 1];
      const nextCursor = lastEvent ? lastEvent.id : null;
      return {
        events,
        nextCursor,
      };
    }),
  getPublications: publicProcedure
    .input(
      z.object({
        take: z.number(),
        cursor: z.number().nullable(),
      })
    )
    .query(async ({ input }) => {
      if (!input.cursor) {
        const firstQueryEvents = await prisma.event.findMany({
          where: {
            type: {
              in: ["PostCreated", "CommentCreated", "MirrorCreated"],
            },
          },
          take: input.take,
          orderBy: { id: "desc" },
        });
        const firstQueryLastEvent =
          firstQueryEvents[firstQueryEvents.length - 1];
        const firstQueryNextCursor = firstQueryLastEvent
          ? firstQueryLastEvent.id
          : null;
        return {
          events: firstQueryEvents,
          nextCursor: firstQueryNextCursor,
        };
      }
      const events = await prisma.event.findMany({
        where: {
          type: {
            in: ["PostCreated", "CommentCreated", "MirrorCreated"],
          },
        },
        take: input.take,
        cursor: {
          id: input.cursor,
        },
        orderBy: {
          id: "desc",
        },
      });
      const lastEvent = events[events.length - 1];
      const nextCursor = lastEvent ? lastEvent.id : null;
      return {
        events,
        nextCursor,
      };
    }),
});
