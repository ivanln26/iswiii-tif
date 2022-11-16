import "styles/globals.css";
import type { AppProps } from "next/app";
import { trpc } from "../utils/trpc";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Component {...pageProps} />
      <ReactQueryDevtools />
    </>
  );
}

export default trpc.withTRPC(App);
