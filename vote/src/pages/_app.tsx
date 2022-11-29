import "styles/globals.css";
import type { AppProps } from "next/app";
import Head from "next/head";
import { trpc } from "@/utils/trpc";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import Navbar from "@/components/navbar";

function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>Voting App</title>
      </Head>
      <Navbar />
      <Component {...pageProps} />
      <ReactQueryDevtools />
    </>
  );
}

export default trpc.withTRPC(App);
