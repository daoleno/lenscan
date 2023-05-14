import lensClient from "@/lib/lensclient";
import { ProfileFragment } from "@lens-protocol/client";
import { MomokaTx } from "@prisma/client";
import { z } from "zod";
import prisma from "../prisma";
import { publicProcedure, router } from "../trpc";

type MomokaTxs = Array<MomokaTx> & { profile?: ProfileFragment };
export const momokaRouter = router({
  getTxs: publicProcedure
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
        WHERE relname = 'MomokaTx';
      `) as any[];
      const count = Number(res[0]);

      if (!input.cursor) {
        const firstQueryTxs = await prisma.momokaTx.findMany({
          take: input.take,
          orderBy: { id: "desc" },
        });
        const firstQueryLastTx = firstQueryTxs[firstQueryTxs.length - 1];
        const firstQueryNextCursor = firstQueryLastTx
          ? firstQueryLastTx.id
          : null;

        const profiles = await lensClient.profile.fetchAll({
          profileIds: [
            ...new Set(firstQueryTxs.map((tx) => (tx.event as any).profileId)),
          ],
        });

        const txs = firstQueryTxs.map((tx) => {
          const profile = profiles.items.find(
            (profile) => profile.id === (tx.event as any).profileId
          );
          return { ...tx, profile };
        }) as MomokaTxs;

        return {
          count,
          list: txs,
          nextCursor: firstQueryNextCursor,
        };
      }

      const txs = await prisma.momokaTx.findMany({
        take: input.take,
        skip: 1,
        cursor: {
          id: input.cursor,
        },
        orderBy: {
          id: "desc",
        },
      });

      const profiles = await lensClient.profile.fetchAll({
        profileIds: [...new Set(txs.map((tx) => (tx.event as any).profileId))],
      });

      const newTxs = txs.map((tx) => {
        const profile = profiles.items.find(
          (profile) => profile.id === (tx.event as any).profileId
        );
        return { ...tx, profile };
      }) as MomokaTxs;

      const lastTx = newTxs[newTxs.length - 1];
      const nextCursor = lastTx ? lastTx.id : null;
      return {
        count,
        list: newTxs,
        nextCursor,
      };
    }),
  getTx: publicProcedure.input(z.string()).query(async ({ input }) => {
    const tx = await prisma.momokaTx.findUnique({
      where: {
        proofTxId: input,
      },
    });
    return tx;
  }),
  getTxsCountByProfileId: publicProcedure
    .input(z.number())
    .query(async ({ input }) => {
      const res = (await prisma.$queryRawUnsafe(
        `SELECT COUNT(*) FROM "MomokaTx" WHERE tx->'ProfileId' = '${input}'`
      )) as any[];
      const count = Number(res[0].count);
      return count;
    }),
  getTxsByProfileId: publicProcedure
    .input(
      z.object({
        profileId: z.number(),
        take: z.number(),
        cursor: z.number().nullable(),
      })
    )
    .query(async ({ input }) => {
      if (!input.cursor) {
        const firstQuerytxs = (await prisma.$queryRawUnsafe(
          `SELECT * FROM "MomokaTx" WHERE data->'ProfileId' = '${input.profileId}' ORDER BY id DESC LIMIT ${input.take}`
        )) as any[];
        const firstQueryLastTx = firstQuerytxs[firstQuerytxs.length - 1];
        const firstQueryNextCursor = firstQueryLastTx
          ? firstQueryLastTx.id
          : null;
        return {
          list: firstQuerytxs,
          nextCursor: firstQueryNextCursor,
        };
      }

      const txs = (await prisma.$queryRawUnsafe(
        `SELECT * FROM "MomokaTx" WHERE data->'ProfileId' = '${input.profileId}' AND id < ${input.cursor} ORDER BY id DESC LIMIT ${input.take}`
      )) as any[];
      const lastTx = txs[txs.length - 1];
      const nextCursor = lastTx ? lastTx.id : null;
      return {
        list: txs,
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
        const firstQuerytxs = await prisma.momokaTx.findMany({
          where: {
            type: {
              in: ["POST_CREATED", "MIRROR_CREATED", "COMMENT_CREATED"],
            },
          },
          take: input.take,
          orderBy: { id: "desc" },
        });
        const firstQueryLastTx = firstQuerytxs[firstQuerytxs.length - 1];
        const firstQueryNextCursor = firstQueryLastTx
          ? firstQueryLastTx.id
          : null;
        return {
          list: firstQuerytxs,
          nextCursor: firstQueryNextCursor,
        };
      }
      const txs = await prisma.momokaTx.findMany({
        where: {
          type: {
            in: ["POST_CREATED", "MIRROR_CREATED", "COMMENT_CREATED"],
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
      const lastTx = txs[txs.length - 1];
      const nextCursor = lastTx ? lastTx.id : null;
      return {
        list: txs,
        nextCursor,
      };
    }),
});
