import type { AppProps } from "next/app";
import Head from "next/head";
import { Toaster } from "react-hot-toast";
import "@/assets/app.scss";

export default function _App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <meta content="width=device-width, initial-scale=1" name="viewport" />
        <title>Botway</title>
      </Head>
      <Component {...pageProps} />
      <Toaster />
    </>
  );
}
