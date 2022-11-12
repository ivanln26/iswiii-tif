import Head from "next/head";

export default function Home() {
  return (
    <div>
      <button className="p-2 bg-sky-500 rounded-full text-white">Vote A</button>
      <button className="p-2 bg-red-500 rounded-full text-white">Vote B</button>
    </div>
  );
}
