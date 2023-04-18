"use client";

import LogoSection from "@/components/Logo";
import { Toaster } from "react-hot-toast";

const Template = ({ children }: { children: React.ReactNode }) => {
  return (
    <main className="bg-bwdefualt flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <Toaster />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          {children}
        </div>
      </section>
    </main>
  );
};

export default Template;
