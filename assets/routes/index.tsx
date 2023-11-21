import { LogoSection } from "../components/Logo.tsx";
import { Head } from "$fresh/runtime.ts";

const App = () => {
  return (
    <>
      <Head>
        <meta charSet="UTF-8" />

        <meta content="width=device-width, initial-scale=1" name="viewport" />

        <link rel="icon" type="image/svg" href="/simple/icon.svg" />
        <link rel="stylesheet" href="/main.css" />

        <title>Botway CDN 📦</title>
      </Head>

      <body style={{ background: "#13111c" }}>
        <div className="block md:hidden">
          <main className="flex flex-col place-content-center items-center place-items-center h-[90vh]">
            <LogoSection />

            <section className="justify-center md:bg-grid-gray-800/[0.6] px-6 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
              <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
                <div className="place-content-center items-center place-items-center justify-center">
                  <h3
                    className="font-medium text-lg pt-3 text-white"
                    style={{ fontFamily: "Farray" }}
                  >
                    Welcome to Botway CDN 📦
                  </h3>
                </div>
              </div>
            </section>
          </main>
        </div>
        <div className="hidden md:block">
          <main className="flex flex-col md:flex-row-reverse md:h-screen">
            <LogoSection />

            <section className="justify-center md:bg-grid-gray-800/[0.6] px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
              <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
                <h3
                  className="font-medium md:text-xl pt-3 text-white"
                  style={{ fontFamily: "Farray" }}
                >
                  Welcome to Botway CDN 📦
                </h3>
                <br />
                <p
                  className="text-white text-sm"
                  style={{ fontFamily: "JetBrains Mono" }}
                >
                  <span
                    className="text-blue-700"
                    style={{ fontFamily: "Farray" }}
                  >
                    Botway CDN{" "}
                  </span>
                  is a small service that hosts all of botway's assets and
                  integrations data 📡
                </p>
              </div>
            </section>
          </main>
        </div>
      </body>
    </>
  );
};

export default App;
