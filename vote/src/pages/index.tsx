import { trpc } from "@/utils/trpc";

export default function Home() {
  const mutation = trpc.voteCreate.useMutation();

  return (
    <div>
      <button
        className="p-2 bg-sky-500 hover:bg-sky-600 disabled:bg-sky-800 rounded-full text-white"
        disabled={mutation.isLoading}
        onClick={() => {
          mutation.mutate({ choice: 1 });
        }}
      >
        Vote A
      </button>
      <button
        className="p-2 bg-red-500 hover:bg-red-600 disabled:bg-red-800 rounded-full text-white"
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
