import { Button } from "@/components/Button";
import LogoSection from "@/components/Logo";
import { fetcher } from "@/lib/fetch";
import { toastStyle } from "@/tools/toast-style";
import Link from "next/link";
import { useRouter } from "next/router";
import { useCallback, useEffect, useRef, useState } from "react";
import toast from "react-hot-toast";

export const ForgetPassword = () => {
  const router = useRouter();

  useEffect(() => {
    if (process.env.NEXT_PUBLIC_FULL != "true") {
      router.push("/");
    }
  }, []);

  const emailRef: any = useRef();

  const [status, setStatus]: any = useState();
  const [email, setEmail] = useState("");

  const onSubmit = useCallback(async (e: any) => {
    e.preventDefault();

    try {
      setStatus("loading");

      await fetcher("/api/user/password/reset", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          email: emailRef.current.value,
        }),
      });

      setEmail(emailRef.current.value);

      setStatus("success");

      toast.success(
        "Reset Password email has been sent successfuly",
        toastStyle
      );
    } catch (e: any) {
      toast.error(e.message, toastStyle);

      setStatus(undefined);
    }
  }, []);

  return (
    <main className="flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          {status === "success" ? (
            <div className="my-2 mb-2">
              <h2 className="text-lg font-medium md:text-2xl text-white">
                Check your inbox
              </h2>
              <p className="text-gray-400 pt-8">
                An email has been sent to{" "}
                <span className="text-blue-700">{email}</span>. Please follow
                the link in that email to reset your password.
              </p>
            </div>
          ) : (
            <>
              <h2 className="text-lg font-medium md:text-2xl text-white">
                Reset Password
              </h2>

              <p className="text-sm pt-1 cursor-pointer">
                <Link href="/sign-in" className="text-blue-700">
                  Return to Sign In page
                </Link>
              </p>

              <div className="my-2 mb-2">
                <form onSubmit={onSubmit}>
                  <p className="text-gray-400 pt-8">
                    Enter the email address associated with your account, and
                    we&apos;ll send you a link to reset your password.
                  </p>
                  <div className="pt-8">
                    <label
                      htmlFor="email"
                      className="block text-gray-500 text-sm font-semibold"
                    >
                      Email
                    </label>
                    <div className="pt-2">
                      <input
                        className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                        ref={emailRef}
                        autoComplete="email"
                        placeholder="Email Address"
                        aria-label="Email Address"
                        type="email"
                        required
                      />
                    </div>
                  </div>

                  <div className="mt-4 space-y-2 flex justify-center">
                    <Button
                      type="success"
                      htmlType="submit"
                      loading={status === "loading"}
                      className="button w-full p-2"
                    >
                      Continue
                    </Button>
                  </div>
                </form>
              </div>
            </>
          )}
        </div>
      </section>
    </main>
  );
};
