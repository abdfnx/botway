import { Button } from "@/components/Button";
import LogoSection from "@/components/Logo";
import { fetcher } from "@/lib/fetch";
import Link from "next/link";
import { useCallback, useRef, useState } from "react";
import toast from "react-hot-toast";

const NewPassword = ({ token }: any) => {
  const passwordRef: any = useRef();

  const [status, setStatus]: any = useState();

  const onSubmit = useCallback(
    async (e: any) => {
      e.preventDefault();

      setStatus("loading");

      try {
        await fetcher("/api/user/password/reset", {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            token,
            password: passwordRef.current.value,
          }),
        });

        setStatus("success");
      } catch (e: any) {
        toast.error(e.message);

        setStatus(undefined);
      }
    },
    [token]
  );
  return (
    <main className="flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          <h2 className="text-lg font-medium md:text-2xl text-white">
            Reset Password
          </h2>

          <p className="text-sm pt-1 cursor-pointer">
            <Link href="/sign-in" className="text-blue-700">
              Return to Sign In page.
            </Link>
          </p>

          {status === "success" ? (
            <>
              <p className="text-gray-400 pt-8">
                Your password has been updated successfully.
              </p>
            </>
          ) : (
            <div className="my-2 mb-2">
              <form onSubmit={onSubmit}>
                <p className="text-gray-400 pt-8">
                  Enter a new password for your account
                </p>
                <div className="pt-8">
                  <label
                    htmlFor="new-password"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    New Password
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={passwordRef}
                      type="password"
                      autoComplete="new-password"
                      placeholder="New Password"
                      required
                    />
                  </div>
                </div>

                <div className="mt-4 space-y-2 flex justify-center">
                  <Button htmlType="submit" type="success" size="large">
                    Reset Password
                  </Button>
                </div>
              </form>
            </div>
          )}
        </div>
      </section>
    </main>
  );
};

const BadLink = () => {
  return (
    <main className="flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          <h2 className="text-lg font-medium md:text-2xl text-white">
            Invalid Link
          </h2>

          <p className="text-sm pt-1 cursor-pointer">
            <Link href="/sign-in" className="text-blue-700">
              Return to Sign In page
            </Link>
          </p>
          <p className="text-gray-400 pt-8">
            It looks like you may have clicked on an invalid link. Please close
            this window and try again.
          </p>
        </div>
      </section>
    </main>
  );
};

export const ForgetPasswordToken = ({ valid, token }: any) => {
  return <>{valid ? <NewPassword token={token} /> : <BadLink />}</>;
};
