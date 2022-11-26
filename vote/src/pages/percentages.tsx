import { trpc } from "@/utils/trpc";

export default function Percentages() {
  const percentages = trpc.votePercentages.useQuery();

  if (percentages.isLoading) {
    return <h1>Cargando...</h1>;
  }

  if (percentages.isError) {
    return <h1>Error!</h1>;
  }

  return (
    <ul>
      {percentages.data.percentages.map(({ choice, percentage }) => (
        <li>{choice} {percentage}</li>
      ))}
    </ul>
  );
}
