import redis from "@/lib/redis";
import { z } from "zod";
import { publicProcedure, router } from "../trpc";

export const appRouter = router({
  voteCreate: publicProcedure
    .input(z.object({ choice: z.number().int() }))
    .mutation(async ({ input }) => {
      const vote = input;
      await redis.publish("votes", JSON.stringify(vote));
      return {
        vote,
      };
    }),
});

export type AppRouter = typeof appRouter;
