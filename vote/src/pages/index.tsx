import { trpc } from "@/utils/trpc";

export default function Home() {
  const mutation = trpc.voteCreate.useMutation();

  return (
    <div className="flex flex-col justify-center p-2">
      <h1 className="text-4xl py-4 text-center font-bold">
        Choose your &nbsp;
        <span className="font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-amber-400 to-orange-500">
          VOTE
        </span>
      </h1>
      <button
        id="btn-vote-a"
        className="p-2 my-2 bg-sky-500 hover:bg-sky-600 disabled:bg-sky-800 rounded-full"
        disabled={mutation.isLoading}
        onClick={() => {
          mutation.mutate({ choice: 1 });
        }}
      >
        Vote A
      </button>
      <button
        id="btn-vote-b"
        className="p-2 my-2 bg-red-500 hover:bg-red-600 disabled:bg-red-800 rounded-full"
        disabled={mutation.isLoading}
        onClick={() => {
          mutation.mutate({ choice: 2 });
        }}
      >
        Vote B
      </button>
      <div className="flex items-center justify-center px-2 py-4 font-bold">
        {mutation.isLoading && (
          <div id="lbl-loading" className="p-2 rounded bg-sky-500">
            Loading...
          </div>
        )}
        {mutation.isError && (
          <div id="lbl-error" className="p-2 rounded bg-rose-600">
            Error!
          </div>
        )}
        {mutation.isSuccess && (
          <div id="lbl-success" className="p-2 rounded bg-green-500">
            Success!
          </div>
        )}
      </div>
    </div>
  );
}
