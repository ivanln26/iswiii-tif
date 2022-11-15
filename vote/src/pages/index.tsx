import { trpc } from "@/utils/trpc";

export default function Home() {
  const mutation = trpc.voteCreate.useMutation();

  return (
    <div>
      <button
        className="p-2 bg-sky-500 rounded-full text-white"
        onClick={() => {
          mutation.mutate({ choice: 1 });
        }}
      >
        Vote A
      </button>
      <button
        className="p-2 bg-red-500 rounded-full text-white"
        onClick={() => {
          mutation.mutate({ choice: 2 });
        }}
      >
        Vote B
      </button>
    </div>
  );
}
