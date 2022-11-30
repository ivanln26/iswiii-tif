import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="px-2 lg:px-4 bg-[#000000]">
      <ul className="flex">
        <li className="px-3 py-4">
          <Link href="/">Home</Link>
        </li>
        <li className="px-3 py-4">
          <Link href="/list">List</Link>
        </li>
        <li className="px-3 py-4">
          <Link href="/percentages">Percentages</Link>
        </li>
      </ul>
    </nav>
  );
}
