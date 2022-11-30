import { Button } from "@/components/Button";
import Layout from "@/components/Layout";
import { LoadingDots } from "@/components/LoadingDots";
import { fetcher } from "@/lib/fetch";
import { useCurrentUser } from "@/lib/user";
import { bg } from "@/tools/colors";
import clsx from "clsx";
import Link from "next/link";
import { useRouter } from "next/router";
import { useCallback, useEffect, useRef, useState } from "react";
import toast from "react-hot-toast";

export const AccountInfo = ({ user, mutate }: any) => {
  const usernameRef: any = useRef();
  const nameRef: any = useRef();
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  const currentPath = router.pathname;
  const NAV_ITEMS = [
    {
      title: "Account Information",
      href: "/settings",
    },
    {
      title: "Admin Settings",
      href: "/settings/admin",
    },
    {
      title: "Change Password",
      href: "/settings/change-password",
    },
  ];

  const onSubmit = useCallback(
    async (e: any) => {
      e.preventDefault();

      try {
        setIsLoading(true);

        const formData = new FormData();

        formData.append("username", usernameRef.current.value);
        formData.append("name", nameRef.current.value);

        const response = await fetcher("/api/user", {
          method: "PATCH",
          body: formData,
        });

        mutate({ user: response.user }, false);

        toast.success("Your profile has been updated", {
          style: {
            borderRadius: "10px",
            backgroundColor: bg,
            color: "#fff",
          },
        });
      } catch (e: any) {
        toast.error(e.message, {
          style: {
            borderRadius: "10px",
            backgroundColor: bg,
            color: "#fff",
          },
        });
      } finally {
        setIsLoading(false);
      }
    },
    [mutate]
  );

  useEffect(() => {
    usernameRef.current.value = user.username;
    nameRef.current.value = user.name;
  }, [user]);

  return (
    <>
      <main className="mx-auto w-full flex-1 px-3 py-4 sm:py-6 sm:px-6 lg:px-8">
        <div className="mx-auto h-full max-w-7xl">
          <div id="content"></div>
          <div className="flex h-full flex-col lg:flex-row">
            <div className="w-full shrink-0 lg:w-1/5">
              <div className="-mt-1 hidden flex-col pr-2 lg:flex">
                {NAV_ITEMS.map((item) => (
                  <Link key={item.title} href={item.href}>
                    <p
                      className={clsx(
                        item.href === currentPath
                          ? "text-blue-700"
                          : "text-gray-400 ",
                        "py-1 transition font-semibold"
                      )}
                      aria-current={
                        item.href === currentPath ? "page" : undefined
                      }
                    >
                      {item.title}
                    </p>
                  </Link>
                ))}
              </div>
              <select
                aria-label="Navigation items"
                className="inline-block rounded border border-gray-800 border-secondary bg-secondary pl-1.5 pr-4 h-6 mb-4 font-semibold lg:hidden"
              >
                {NAV_ITEMS.map((item) => (
                  <option value={item.href}>{item.title}</option>
                ))}
              </select>
            </div>
            <div className="flex-1">
              <div className="mb-4">
                <h1 className="text-xl text-gray-400 font-semibold">
                  Account Information
                </h1>
              </div>
              <form onSubmit={onSubmit}>
                <input type="hidden" name="_method" value="PATCH" />

                <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-2">
                  <div className="max-w-md">
                    <label
                      htmlFor="email"
                      className="block text-gray-500 text-sm font-semibold"
                    >
                      Email
                    </label>
                    <div className="pt-2">
                      <input
                        className="w-full border px-1.5 py-sm bg-secondary trsn bg border-gray-800 text-gray-400 cursor-not-allowed sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                        aria-label="Email Address"
                        type="email"
                        value={user.email}
                        disabled
                      />

                      <p className="mb-3 py-2 text-sm text-gray-400">
                        You cannot change your email
                      </p>
                    </div>
                  </div>
                  <div className="max-w-md lg:pl-6">
                    <label
                      htmlFor="name"
                      className="block text-gray-500 text-sm font-semibold"
                    >
                      Name
                    </label>
                    <div className="pt-2">
                      <input
                        className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                        ref={nameRef}
                        placeholder="name"
                        type="text"
                      />
                    </div>
                  </div>
                  <div className="max-w-md pt-4 sm:pt-4">
                    <label
                      htmlFor="username"
                      className="block text-gray-500 text-sm font-semibold"
                    >
                      Username
                    </label>
                    <div className="pt-2 mb-6">
                      <input
                        className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                        ref={usernameRef}
                        placeholder="username"
                        type="text"
                      />
                    </div>
                  </div>
                </div>
                <div className="border-t border-gray-800">
                  <Button htmlType="submit" type="success" loading={isLoading}>
                    Save Changes
                  </Button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </main>
    </>
  );
};

export const Settings = () => {
  const { data, error, mutate } = useCurrentUser();
  const loading = !data && !error;
  const router = useRouter();

  const PushToSignIn = () => {
    useEffect(() => {
      if (!data?.user) {
        router.push("/sign-in");
      }
    }, []);

    return <></>;
  };

  return (
    <>
      {loading ? (
        <LoadingDots>Loading</LoadingDots>
      ) : data?.user ? (
        <Layout title="General Settings">
          <span className="flex items-center">
            {/* <AboutYou user={data.user} mutate={mutate} /> */}
            <AccountInfo user={data.user} mutate={mutate} />
          </span>
        </Layout>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};
