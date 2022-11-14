import redis from "@/lib/redis";
import { z } from "zod";
import { publicProcedure, router } from "../trpc";

export const appRouter = router({
  vote: publicProcedure
    .input(z.object({ choice: z.number().int() }))
    .mutation(async ({ input }) => {
      await redis.publish("votes", `{"vote": ${input.choice}}`);
      return {
        detail: "voted correctly",
      };
    }),
});

export type AppRouter = typeof appRouter;
