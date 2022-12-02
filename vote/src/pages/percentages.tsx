import { trpc } from "@/utils/trpc";

export default function Percentages() {
  const percentages = trpc.votePercentages.useQuery(undefined, {
    refetchInterval: 1000,
  });

  if (percentages.isLoading) {
    return <h1>Cargando...</h1>;
  }

  if (percentages.isError) {
    return <h1>Error!</h1>;
  }

  return (
    <>
      <h1 className="p-2 text-4xl font-bold">Percentages</h1>
      <ul className="p-2">
        {percentages.data.percentages.map(({ choice, percentage }, i) => (
          <li key={i} className="flex flex-col">
            <h1 className="text-2xl">Opciones {choice}</h1>
            <div className="flex p-2">
              <h2 className="pr-4">{percentage}%</h2>
              <div className="bg-neutral-800 w-full rounded-md border border-amber-900 overflow-hidden">
                <div
                  className="bg-amber-600 h-full overflow"
                  style={{ width: `${percentage}%` }}
                />
              </div>
            </div>
          </li>
        ))}
      </ul>
    </>
  );
}
