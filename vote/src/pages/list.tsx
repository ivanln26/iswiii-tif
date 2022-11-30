import { trpc } from "@/utils/trpc";

export default function List() {
  const list = trpc.voteList.useQuery();

  if (list.isLoading) {
    return <h1>Cargando...</h1>;
  }

  if (list.error) {
    return <h1>Error!</h1>;
  }

  return (
    <>
      <h1 className="p-2 text-4xl font-bold">List</h1>
      <div className="p-2">
        <ul className="p-2 list-disc list-inside">
          {list.data.votes.map((vote, i) => (
            <li key={i}>{JSON.stringify(vote)}</li>
          ))}
        </ul>
      </div>
    </>
  );
}
