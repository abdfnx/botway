import React from "react";
import Head from "next/head";
import { LogoSection } from "../components/logo";

function MainPage() {
  return (
    <>
      <main className="flex flex-col md:flex-row-reverse md:h-screen sm:pt-20">
        <Head>
          <title>abdfnx/botway 🤖</title>
          <link rel="icon" href="/icon.svg" />
        </Head>
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <p className="text-lg pt-2 text-gray-400">abdfnx/botway 🤖</p>
          </div>
        </section>
      </main>
    </>
  );
}

export default MainPage;
