import { Button } from "@/components/Button";
import LogoSection from "@/components/Logo";
import { fetcher } from "@/lib/fetch";
import { useCurrentUser } from "@/lib/user";
import { bgSecondary } from "@/tools/colors";
import Link from "next/link";
import { useRouter } from "next/router";
import { useCallback, useRef, useState } from "react";
import toast from "react-hot-toast";

const SignUp = () => {
  const emailRef: any = useRef();
  const passwordRef: any = useRef();
  const usernameRef: any = useRef();
  const nameRef: any = useRef();
  const { mutate } = useCurrentUser();
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  const onSubmit = useCallback(
    async (e: any) => {
      e.preventDefault();

      try {
        setIsLoading(true);

        const response = await fetcher("/api/users", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            email: emailRef.current.value,
            name: nameRef.current.value,
            password: passwordRef.current.value,
            username: usernameRef.current.value,
            githubApiToken: "",
            railwayApiToken: "",
            renderApiToken: "",
          }),
        });

        mutate({ user: response.user }, false);

        toast.success("Your account has been created", {
          style: {
            borderRadius: "10px",
            backgroundColor: bgSecondary,
            color: "#fff",
          },
        });

        router.replace("/");
      } catch (e: any) {
        toast.error(e.message, {
          style: {
            borderRadius: "10px",
            backgroundColor: bgSecondary,
            color: "#fff",
          },
        });
      } finally {
        setIsLoading(false);
      }
    },
    [mutate, router]
  );

  return (
    <main className="flex flex-col md:flex-row-reverse md:h-screen">
      <LogoSection />

      <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
        <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
          <h2 className="text-lg font-medium md:text-2xl text-white">
            Create Admin User
          </h2>

          <p className="text-sm text-gray-500 pt-1 cursor-pointer">
            Already have an account?{" "}
            <Link href="/sign-in" className="text-blue-700">
              Sign in.
            </Link>
          </p>

          <div className="my-2 mb-2">
            <form onSubmit={onSubmit}>
              <div className="pt-8">
                <label
                  htmlFor="name"
                  className="block text-gray-500 text-sm font-semibold"
                >
                  Name
                </label>
                <div className="pt-2">
                  <input
                    className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                    ref={nameRef}
                    autoComplete="name"
                    placeholder="Name"
                    aria-label="Name"
                    required
                  />
                </div>
              </div>

              <div className="pt-8">
                <label
                  htmlFor="username"
                  className="block text-gray-500 text-sm font-semibold"
                >
                  Username
                </label>
                <div className="pt-2">
                  <input
                    className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                    ref={usernameRef}
                    autoComplete="username"
                    placeholder="Username"
                    aria-label="Username"
                    required
                  />
                </div>
              </div>

              <div className="mt-6">
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

              <div className="mt-4">
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
                    autoComplete="new-password"
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
                  Create new Botway Account
                </Button>
              </div>
            </form>
          </div>
        </div>
      </section>
    </main>
  );
};

export default SignUp;
