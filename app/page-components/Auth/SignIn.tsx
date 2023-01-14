import { Button } from "@/components/Button";
import LogoSection from "@/components/Logo";
import { fetcher } from "@/lib/fetch";
import { useCurrentUser } from "@/lib/user";
import { toastStyle } from "@/tools/toast-style";
import Link from "next/link";
import { useRouter } from "next/router";
import { useCallback, useEffect, useRef, useState } from "react";
import toast from "react-hot-toast";

const SignIn = () => {
  const emailRef: any = useRef();
  const passwordRef: any = useRef();

  const [isLoading, setIsLoading] = useState(false);

  const { data: { user } = {}, mutate, isValidating } = useCurrentUser();

  const router = useRouter();

  useEffect(() => {
    if (isValidating) return;

    if (user) router.replace("/");
  }, [user, router, isValidating]);

  const onSubmit = useCallback(
    async (e: any) => {
      setIsLoading(true);

      e.preventDefault();

      try {
        const response = await fetcher("/api/auth", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            email: emailRef.current.value,
            password: passwordRef.current.value,
          }),
        });

        mutate({ user: response.user }, false);

        toast.success("You have been logged in.", toastStyle);
      } catch (e) {
        toast.error("Incorrect email or password.", toastStyle);
      } finally {
        setIsLoading(false);
      }
    },
    [mutate]
  );

  return (
    <main className="flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          <h2 className="text-lg font-medium md:text-2xl text-white">
            Sign in
          </h2>

          <p className="text-sm text-gray-500 pt-1 cursor-pointer">
            You don't have an Account?{" "}
            <Link href="/sign-up" className="text-blue-700">
              Sign up for an account
            </Link>
          </p>

          {process.env.NEXT_PUBLIC_FULL == "true" ? (
            <p className="text-sm text-gray-500 pt-1 cursor-pointer">
              Forget Password?{" "}
              <Link href="/forget-password" className="text-blue-700">
                Reset
              </Link>
            </p>
          ) : (
            <></>
          )}

          <div className="my-2 mb-2">
            <form onSubmit={onSubmit}>
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

              <div className="mt-6">
                <label
                  htmlFor="password"
                  className="block text-gray-500 text-sm font-semibold"
                >
                  Password
                </label>
                <div className="pt-2">
                  <input
                    className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                    ref={passwordRef}
                    autoComplete="current-password"
                    aria-label="Password"
                    placeholder="••••••••••••••••"
                    type="password"
                    required
                  />
                </div>
              </div>

              <div className="mt-4 space-y-2 flex justify-center">
                <Button
                  type="success"
                  htmlType="submit"
                  loading={isLoading}
                  className="button w-full p-2"
                >
                  Sign in
                </Button>
              </div>
            </form>
          </div>
        </div>
      </section>
    </main>
  );
};

export default SignIn;
