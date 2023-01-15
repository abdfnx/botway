import { LoadingDots } from "@/components/LoadingDots";
import { fetcher } from "@/lib/fetch";
import { useCurrentUser } from "@/lib/user";
import { useRouter } from "next/router";
import { Fragment, useCallback, useEffect } from "react";
import toast from "react-hot-toast";
import { Menu, Transition } from "@headlessui/react";
import { UserAvatar } from "@/components/UserAvatar";
import clsx from "clsx";
import { toastStyle } from "@/tools/toast-style";
import {
  AlertIcon,
  ChevronDownIcon,
  RocketIcon,
  SignOutIcon,
  SlidersIcon,
  VersionsIcon,
} from "@primer/octicons-react";

const EmailNotVerifiedBanner = ({ user }: any) => {
  const verify = useCallback(async () => {
    try {
      await fetcher("/api/user/email/verify", { method: "POST" });

      toast.success(
        "An email has been sent to your mailbox. Follow the instruction to verify your email.",
        toastStyle
      );
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    }
  }, []);

  if (user.emailVerified || process.env.NEXT_PUBLIC_FULL != "true")
    return <></>;

  return (
    <div className="antialiased font-sans font-normal text-sm text-gray-800 leading-6 relative flex items-center w-full whitespace-pre-wrap justify-between px-4 rounded-none border border-solid bg-orange-25 border-orange-600 box-border border-l-0 border-r-0">
      <div className="flex py-2 relative max-w-[1280px] mx-auto px-10 pr-12">
        <div className="flex mr-4 mt-1">
          <AlertIcon className="fill-yellow-400" size={16} aria-hidden="true" />
        </div>
        <div className="flex-col">
          <span>
            Your Email <a className="text-blue-700">{user.email}</a> is not
            verified -{" "}
            <a
              className="text-blue-700 cursor-pointer underline"
              onClick={verify}
            >
              Verify
            </a>
          </span>
        </div>
      </div>
    </div>
  );
};

const Layout = ({ children, title }: any) => {
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

  const onSignOut = useCallback(async () => {
    try {
      await fetcher("/api/auth", {
        method: "DELETE",
      });

      toast.success("You have been signed out", toastStyle);

      mutate({ user: null });
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    }
  }, [mutate]);

  return (
    <>
      {loading ? (
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : data?.user ? (
        <div className="min-h-screen bg">
          <div className="flex flex-col flex-1">
            <EmailNotVerifiedBanner user={data.user} />
            <div className="relative flex-shrink-0 flex h-16">
              <div className="flex-1 px-4 flex justify-between sm:px-6 lg:max-w-6xl lg:mx-auto lg:px-8">
                <button type="button" className="text-gray-400 outline-none">
                  <img
                    className="h-11 w-11 pt-1 mx-auto"
                    src="https://cdn-botway.deno.dev/icon.svg"
                    alt="Botway Logo"
                    onClick={() => router.push("/")}
                  />
                </button>
                <div className="flex-1 flex"></div>
                <div>
                  <div className="ml-4 flex items-center md:ml-6 pt-3">
                    <Menu as="div" className="ml-3 relative">
                      <div>
                        <Menu.Button className="max-w-xs rounded-full flex transition items-center text-sm outline-none focus:ring-gray-800 p-2 lg:rounded-md">
                          <UserAvatar data={data.user.email} size={25} />

                          <ChevronDownIcon
                            className="flex-shrink-0 pl-1 h-5 w-5 text-gray-400"
                            aria-hidden="true"
                          />
                        </Menu.Button>
                      </div>
                      <Transition
                        as={Fragment}
                        enter="transition ease-out duration-100"
                        enterFrom="transform opacity-0 scale-95"
                        enterTo="transform opacity-100 scale-100"
                        leave="transition ease-in duration-75"
                        leaveFrom="transform opacity-100 scale-100"
                        leaveTo="transform opacity-0 scale-95"
                      >
                        <Menu.Items className="origin-top-right bg absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 border border-gray-800 ring-1 ring-gray-800 ring-opacity-5 focus:outline-none z-10">
                          <Menu.Item>
                            {({ active }) => (
                              <a
                                href={"/settings"}
                                className={clsx(
                                  active ? "bg-secondary" : "",
                                  "transition block mx-2 my-1 rounded-md cursor-pointer px-4 py-2 text-sm text-gray-400"
                                )}
                              >
                                <SlidersIcon size={18} className="pr-1" />{" "}
                                Settings
                              </a>
                            )}
                          </Menu.Item>
                          <Menu.Item>
                            {({ active }) => (
                              <a
                                href={"https://botway.deno.dev/docs/ui"}
                                className={clsx(
                                  active ? "bg-secondary" : "",
                                  "transition block mx-2 my-1 rounded-md cursor-pointer px-4 py-2 text-sm text-gray-400"
                                )}
                                target="_blank"
                              >
                                <RocketIcon size={18} className="pr-1" /> Docs
                              </a>
                            )}
                          </Menu.Item>
                          <Menu.Item>
                            {({ active }) => (
                              <a
                                href={"https://botway.deno.dev/changelog"}
                                className={clsx(
                                  active ? "bg-secondary" : "",
                                  "transition block mx-2 my-1 rounded-md cursor-pointer px-4 py-2 text-sm text-gray-400"
                                )}
                                target="_blank"
                              >
                                <VersionsIcon size={18} className="pr-1" />{" "}
                                Changelog
                              </a>
                            )}
                          </Menu.Item>
                          <Menu.Item>
                            {({ active }) => (
                              <a
                                onClick={onSignOut}
                                className={clsx(
                                  active ? "bg-secondary" : "",
                                  "transition text-red-600 block mx-2 my-1 rounded-md cursor-pointer px-4 py-2 text-sm"
                                )}
                              >
                                <SignOutIcon size={18} className="pr-1" /> Sign
                                Out
                              </a>
                            )}
                          </Menu.Item>
                        </Menu.Items>
                      </Transition>
                    </Menu>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <main className="flex-1 pb-8">
            <div>
              <div className="px-4 sm:px-6 lg:max-w-6xl lg:mx-auto lg:px-8">
                <div className="pt-6 pb-2 md:flex md:items-center md:justify-between">
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center">
                      <div>
                        <div className="flex items-center">
                          <h1 className="text-xl font-bold leading-7 text-gray-500 sm:leading-9 sm:truncate">
                            {title}
                          </h1>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div className="mt-2">
              <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
                {children}
              </div>
            </div>
          </main>
        </div>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};

export default Layout;
