import { Head, Html, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html className="bg-black">
      <Head>
        <title>Voting App</title>
      </Head>
      <body className="text-white">
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
