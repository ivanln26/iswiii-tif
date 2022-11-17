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
    <ul>
      {list.data.votes.map((vote, i) => <li key={i}>{vote.choice}</li>)}
    </ul>
  );
}
