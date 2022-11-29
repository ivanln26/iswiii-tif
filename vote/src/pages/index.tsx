import { trpc } from "@/utils/trpc";

export default function Home() {
  const mutation = trpc.voteCreate.useMutation();

  return (
    <div className="flex flex-col justify-center h-screen p-2">
      <h1 className="text-4xl py-4 text-center font-bold">
        Choose your &nbsp;
        <span className="font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-amber-400 to-orange-500">
          VOTE
        </span>
      </h1>
      <button
        className="p-2 my-2 bg-sky-500 hover:bg-sky-600 disabled:bg-sky-800 rounded-full"
        disabled={mutation.isLoading}
        onClick={() => {
          mutation.mutate({ choice: 1 });
        }}
      >
        Vote A
      </button>
      <button
        className="p-2 my-2 bg-red-500 hover:bg-red-600 disabled:bg-red-800 rounded-full"
        disabled={mutation.isLoading}
        onClick={() => {
          mutation.mutate({ choice: 2 });
        }}
      >
        Vote B
      </button>
    </div>
  );
}
