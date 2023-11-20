"use client";

import LogoSection from "@/components/Logo";
import { Toaster } from "react-hot-toast";

const Template = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <div className="block md:hidden">
        <main className="flex flex-col place-content-center h-[90vh]">
          <LogoSection />

          <Toaster />

          <section className="justify-center md:bg-grid-gray-800/[0.6] px-6 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
            <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
              {children}
            </div>
          </section>
        </main>
      </div>
      <div className="hidden md:block">
        <main className="flex flex-col md:flex-row-reverse md:h-screen">
          <LogoSection />

          <Toaster />

          <section className="justify-center md:bg-grid-gray-800/[0.6] px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
            <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
              {children}
            </div>
          </section>
        </main>
      </div>
    </>
  );
};

export default Template;
