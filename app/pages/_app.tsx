import type { AppProps } from "next/app";
import Head from "next/head";
import "@/assets/app.scss";
import { Toaster } from "react-hot-toast";

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <meta content="width=device-width, initial-scale=1" name="viewport" />
        <title>🤖 Botway</title>
      </Head>
      <Component {...pageProps} />
      <Toaster />
    </>
  );
}
