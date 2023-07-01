import {
  ArrowUpRightIcon,
  ContainerIcon,
  GearIcon,
  HomeIcon,
  MarkGithubIcon,
  SparkleFillIcon,
} from "@primer/octicons-react";
import { Toaster } from "react-hot-toast";
import { SignOut } from "@/supabase/auth/container/sign-out";

export const DashLayout = ({ children, name, href }: any) => {
  return (
    <>
      <Toaster />

      <body>
        <div className="h-screen min-h-screen flex flex-col">
          <div className="flex h-full">
            <main className="min-h-screen flex flex-col flex-1 w-full overflow-y-auto">
              <div className="flex min-h-full">
                <div className="hidden md:block h-screen text-white min-h-screen bg-secondary w-64 overflow-auto border-r border-gray-800">
                  <div className="mb-2">
                    <div className="flex h-12 max-h-12 items-center border-b px-6 border-gray-800">
                      <h4
                        className="mb-0 text-lg truncate"
                        style={{ fontFamily: "Farray" }}
                        title="Botway"
                      >
                        Botway
                      </h4>
                    </div>
                  </div>
                  <div className="-mt-1">
                    <nav>
                      <ul>
                        <div className="border-b py-5 px-6 border-gray-800">
                          <div className="flex space-x-3 mb-2 font-normal">
                            <span className="text-sm text-gray-200 w-full">
                              Dashboard
                            </span>
                          </div>
                          <ul className="space-y-1">
                            <a className="block" target="_self" href="/">
                              <button className="group flex max-w-full cursor-pointer items-center space-x-2 border-gray-800 py-1 font-normal outline-none focus-visible:z-10 focus-visible:ring-1">
                                <ContainerIcon />
                                <span
                                  title="My projects"
                                  className="w-full truncate text-sm text-gray-400 hover:text-white transition"
                                >
                                  My projects
                                </span>
                              </button>
                            </a>
                          </ul>
                        </div>
                        <div className="border-b py-5 px-6 border-gray-800">
                          <div className="flex space-x-3 mb-2 font-normal">
                            <span className="text-sm text-gray-200 w-full">
                              Account
                            </span>
                          </div>
                          <ul className="space-y-1">
                            <a
                              className="block"
                              target="_self"
                              href="/settings"
                            >
                              <button className="group flex max-w-full cursor-pointer items-center space-x-2 border-gray-800 py-1 font-normal outline-none focus-visible:z-10 focus-visible:ring-1">
                                <GearIcon />
                                <span
                                  title="Settings"
                                  className="w-full truncate text-sm text-gray-400 hover:text-white transition"
                                >
                                  Settings
                                </span>
                              </button>
                            </a>
                          </ul>
                          <ul className="space-y-1">
                            <a
                              className="block"
                              target="_self"
                              style={{ marginLeft: "0rem" }}
                              href="/settings/tokens"
                            >
                              <a
                                href="https://railway.app/account"
                                target="_blank"
                                className="group flex max-w-full cursor-pointer items-center space-x-2 border-gray-800 py-1 font-normal outline-none focus-visible:z-10 focus-visible:ring-1"
                              >
                                <img
                                  src="https://cdn-botway.deno.dev/icons/railway.svg"
                                  width={16}
                                />
                                <span
                                  title="railway"
                                  className="w-full truncate text-sm text-gray-400 hover:text-white transition"
                                >
                                  My Account on Railway
                                </span>
                              </a>
                            </a>
                          </ul>
                        </div>
                        <div className="border-b py-5 px-6 border-gray-800">
                          <div className="flex space-x-3 mb-2 font-normal">
                            <span className="text-sm text-gray-200 w-full">
                              Resources
                            </span>
                          </div>
                          <ul className="space-y-1">
                            <a
                              className="block"
                              target="_blank"
                              href="https://botway.deno.dev/docs"
                            >
                              <button className="group flex max-w-full cursor-pointer items-center space-x-2 border-gray-800 py-1 font-normal outline-none focus-visible:z-10 focus-visible:ring-1">
                                <ArrowUpRightIcon />
                                <span
                                  title="Docs"
                                  className="w-full truncate text-sm text-gray-400 hover:text-white transition"
                                >
                                  Docs
                                </span>
                              </button>
                            </a>
                          </ul>
                          <ul className="space-y-1">
                            <a
                              className="block"
                              target="_blank"
                              href="https://botway.deno.dev/changelog"
                            >
                              <button className="group flex max-w-full cursor-pointer items-center space-x-2 border-gray-800 py-1 font-normal outline-none focus-visible:z-10 focus-visible:ring-1">
                                <ArrowUpRightIcon />
                                <span
                                  title="Changelog"
                                  className="w-full truncate text-sm text-gray-400 hover:text-white transition"
                                >
                                  Changelog
                                </span>
                              </button>
                            </a>
                          </ul>
                          <ul className="space-y-1">
                            <a
                              className="block"
                              target="_blank"
                              href="https://github.com/abdfnx/botway"
                            >
                              <button className="group flex max-w-full cursor-pointer items-center space-x-2 border-gray-800 py-1 font-normal outline-none focus-visible:z-10 focus-visible:ring-1">
                                <MarkGithubIcon />
                                <span
                                  title="Botway Repo"
                                  className="w-full truncate text-sm text-gray-400 hover:text-white transition"
                                >
                                  Botway Repo
                                </span>
                              </button>
                            </a>
                          </ul>
                        </div>
                        <SignOut />
                      </ul>
                    </nav>
                  </div>
                </div>
                <div className="flex bg-bwdefualt flex-1 flex-col">
                  <div className="flex h-12 max-h-12 items-center justify-between py-2 px-5 border-b border-gray-800">
                    <div className="-ml-2 flex items-center text-sm">
                      <span className="flex border-none rounded p-0 outline-none outline-offset-1 transition-all focus:outline-4">
                        <span className="relative inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-gray-200 shadow-none text-xs px-2.5 py-1">
                          <span className="truncate">
                            <svg
                              width="16"
                              height="16"
                              viewBox="0 0 16 16"
                              fill="none"
                              xmlns="http://www.w3.org/2000/svg"
                            >
                              <mask
                                id="mask0_3107_32"
                                // style="mask-type:luminance"
                                maskUnits="userSpaceOnUse"
                                x="0"
                                y="7"
                                width="16"
                                height="9"
                              >
                                <path
                                  d="M0 7.04004H16V7.36004C16 10.0483 16 11.3924 15.4768 12.4192C15.0166 13.3224 14.2823 14.0567 13.3792 14.5168C12.3524 15.04 11.0083 15.04 8.32 15.04H7.68C4.99174 15.04 3.64762 15.04 2.62085 14.5168C1.71766 14.0567 0.983362 13.3224 0.523168 12.4192C0 11.3924 0 10.0483 0 7.36004V7.04004Z"
                                  fill="white"
                                />
                              </mask>
                              <g mask="url(#mask0_3107_32)">
                                <path
                                  d="M17.6004 10.24C17.6004 13.7746 14.735 16.64 11.2004 16.64H4.80039C1.26577 16.64 -1.59961 13.7746 -1.59961 10.24H1.60039C1.60039 12.0074 3.03308 13.44 4.80039 13.44H11.2004C12.9677 13.44 14.4004 12.0074 14.4004 10.24H17.6004ZM4.80039 16.64C1.26577 16.64 -1.59961 13.7746 -1.59961 10.24V7.04004H1.60039V10.24C1.60039 12.0074 3.03308 13.44 4.80039 13.44V16.64ZM17.6004 7.04004V10.24C17.6004 13.7746 14.735 16.64 11.2004 16.64V13.44C12.9677 13.44 14.4004 12.0074 14.4004 10.24V7.04004H17.6004Z"
                                  fill="white"
                                />
                              </g>
                              <mask
                                id="mask1_3107_32"
                                // style="mask-type:luminance"
                                maskUnits="userSpaceOnUse"
                                x="0"
                                y="0"
                                width="16"
                                height="8"
                              >
                                <path
                                  d="M16 8H0V7.68C0 4.99174 0 3.64762 0.523168 2.62085C0.983362 1.71766 1.71766 0.98336 2.62085 0.523168C3.64762 3.8147e-08 4.99174 0 7.68 0H8.32C11.0083 0 12.3524 3.8147e-08 13.3792 0.523168C14.2823 0.98336 15.0166 1.71766 15.4768 2.62085C16 3.64762 16 4.99174 16 7.68V8Z"
                                  fill="white"
                                />
                              </mask>
                              <g mask="url(#mask1_3107_32)">
                                <path
                                  d="M-1.59961 4.7999C-1.59961 1.26528 1.26577 -1.6001 4.80039 -1.6001H11.2004C14.735 -1.6001 17.6004 1.26528 17.6004 4.7999H14.4004C14.4004 3.03259 12.9677 1.5999 11.2004 1.5999H4.80039C3.03308 1.5999 1.60039 3.03259 1.60039 4.7999H-1.59961ZM11.2004 -1.6001C14.735 -1.6001 17.6004 1.26528 17.6004 4.7999V7.9999H14.4004V4.7999C14.4004 3.03259 12.9677 1.5999 11.2004 1.5999V-1.6001ZM-1.59961 7.9999V4.7999C-1.59961 1.26528 1.26577 -1.6001 4.80039 -1.6001V1.5999C3.03308 1.5999 1.60039 3.03259 1.60039 4.7999V7.9999H-1.59961Z"
                                  fill="white"
                                />
                              </g>
                              <path
                                d="M5.11304 6.08008H4.8C4.35817 6.08008 4 6.43825 4 6.88008V8.17862C4 8.62045 4.35817 8.97862 4.8 8.97862H5.11304C5.55487 8.97862 5.91304 8.62045 5.91304 8.17862V6.88008C5.91304 6.43825 5.55487 6.08008 5.11304 6.08008Z"
                                fill="white"
                              />
                              <path
                                d="M11.199 6.08008H10.8859C10.4441 6.08008 10.0859 6.43825 10.0859 6.88008V8.17862C10.0859 8.62045 10.4441 8.97862 10.8859 8.97862H11.199C11.6408 8.97862 11.999 8.62045 11.999 8.17862V6.88008C11.999 6.43825 11.6408 6.08008 11.199 6.08008Z"
                                fill="white"
                              />
                            </svg>
                          </span>
                        </span>
                      </span>
                      <span className="text-gray-600">
                        <svg
                          viewBox="0 0 24 24"
                          width="16"
                          height="16"
                          stroke="currentColor"
                          strokeWidth="1"
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          fill="none"
                        >
                          <path d="M16 3.549L7.12 20.600"></path>
                        </svg>
                      </span>
                      <span className="flex border-none rounded p-0 outline-none outline-offset-1 transition-all focus:outline-4">
                        <span className="relative inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-gray-200 shadow-none text-xs px-2.5 py-1">
                          <span className="truncate">{name}</span>
                        </span>
                      </span>
                      <span className="text-gray-600">
                        <svg
                          viewBox="0 0 24 24"
                          width="16"
                          height="16"
                          stroke="currentColor"
                          strokeWidth="1"
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          fill="none"
                        >
                          <path d="M16 3.549L7.12 20.600"></path>
                        </svg>
                      </span>
                      <a
                        href={
                          href.toLowerCase() === "projects"
                            ? "/"
                            : `/${href.toLowerCase()}`
                        }
                        className="text-gray-200 block px-2 py-1 text-xs leading-5 focus:outline-none"
                      >
                        {href}
                      </a>
                    </div>
                    <div className="flex items-center space-x-2">
                      <a
                        href="/"
                        className="flex border-gray-800 rounded outline-none outline-offset-1 transition-all focus:outline-4"
                      >
                        <span className="relative cursor-pointer inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-blue-700 shadow-sm text-xs px-2.5 py-1">
                          <HomeIcon />
                        </span>
                      </a>
                      <button
                        type="button"
                        onClick={() => {}}
                        className="flex border-gray-800 hover:bg-secondary border rounded p-0 outline-none outline-offset-1 transition-all focus:outline-4"
                      >
                        <span className="relative cursor-pointer inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-blue-700 shadow-sm text-xs px-2.5 py-1">
                          <SparkleFillIcon />

                          <span className="hidden font-extrabold text-gray-500 md:block">
                            AI
                          </span>
                        </span>
                      </button>
                    </div>
                  </div>
                  {children}
                </div>
              </div>
            </main>
          </div>
        </div>
      </body>
    </>
  );
};
