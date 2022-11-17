import redis from "@/lib/redis";
import { z } from "zod";
import { publicProcedure, router } from "../trpc";

const Vote = z.object({
  choice: z.number().int(),
});

const Votes = z.array(Vote);

export const appRouter = router({
  voteCreate: publicProcedure
    .input(Vote)
    .mutation(async ({ input }) => {
      const vote = input;
      await redis.publish("votes", JSON.stringify(vote));
      return {
        vote,
      };
    }),
  voteList: publicProcedure.query(async () => {
    const res = await fetch(`${process.env.BACKEND_URI}/votes`);
    const json = await res.json();
    const votes = Votes.parse(json);
    return {
      votes: votes,
    };
  }),
});

export type AppRouter = typeof appRouter;
