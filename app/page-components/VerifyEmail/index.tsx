import LogoSection from "@/components/Logo";
import Link from "next/link";

export const VerifyEmail = ({ valid }: any) => {
  return (
    <main className="flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          <h2 className="text-lg font-medium md:text-2xl text-white">
            {valid ? "Your Email is verified 🤝" : "Ooops"}
          </h2>

          <p className="text-sm pt-2 cursor-pointer">
            <Link href="/" className="text-blue-700">
              Go back home
            </Link>
          </p>

          <p className="text-gray-400 pt-8">
            {valid
              ? "Thank you for verifying your email address. You can close this page."
              : "It looks like you may have clicked on an invalid link. Please close this window and try again."}
          </p>
        </div>
      </section>
    </main>
  );
};
