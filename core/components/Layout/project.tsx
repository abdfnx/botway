import {
  CodespacesIcon,
  GearIcon,
  HomeIcon,
  ListUnorderedIcon,
  PackageIcon,
  ServerIcon,
  SlidersIcon,
  SparkleFillIcon,
} from "@primer/octicons-react";
import { UserAvatar } from "../UserAvatar";
import { Tooltip } from "flowbite-react";
import clsx from "clsx";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { useRouter } from "next/navigation";
import { Toaster } from "react-hot-toast";

export const ProjectLayout = ({
  user,
  projectId,
  projectName,
  children,
  grid,
  projectZBID,
  projectServiceID,
  latestDeployment,
  noMargin,
}: any) => {
  const router = useRouter();

  const openAtZeabur = async () => {
    const { payload: zeaburProjectId } = await jwtDecrypt(
      projectZBID,
      BW_SECRET_KEY,
    );

    router.push(`https://dash.zeabur.com/projects/${zeaburProjectId.data}`);
  };

  const openLogs = async () => {
    const { payload: zeaburProjectId } = await jwtDecrypt(
      projectZBID,
      BW_SECRET_KEY,
    );

    const { payload: zeaburServiceId } = await jwtDecrypt(
      projectServiceID,
      BW_SECRET_KEY,
    );

    router.push(
      `https://dash.zeabur.com/projects/${zeaburProjectId.data}/services/${zeaburServiceId.data}/deployments/${latestDeployment}`,
    );
  };

  return (
    <>
      <Toaster />

      <div className="min-h-full flex flex-col">
        <div className="flex h-full">
          <div className="hidden md:flex w-14 flex-col bg-secondary justify-between overflow-y-hidden p-2 border-r border-gray-800 h-screen">
            <ul className="flex flex-col space-y-2">
              <a
                className="block mt-[0.35rem] mb-2 self-center items-center"
                href="/"
              >
                <svg
                  width="25"
                  height="24"
                  viewBox="0 0 25 24"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <mask
                    id="mask0_3107_32"
                    //   style="mask-type:luminance"
                    maskUnits="userSpaceOnUse"
                    x="0"
                    y="11"
                    width="25"
                    height="13"
                  >
                    <path
                      d="M0 11H25V11.5C25 15.7004 25 17.8006 24.1826 19.4049C23.4635 20.8162 22.3161 21.9635 20.9049 22.6825C19.3006 23.5 17.2004 23.5 13 23.5H12C7.7996 23.5 5.6994 23.5 4.09508 22.6825C2.68385 21.9635 1.5365 20.8162 0.81745 19.4049C0 17.8006 0 15.7004 0 11.5V11Z"
                      fill="white"
                    />
                  </mask>
                  <g mask="url(#mask0_3107_32)">
                    <path
                      d="M27.5 16C27.5 21.5227 23.0229 26 17.5 26H7.5C1.97715 26 -2.5 21.5227 -2.5 16H2.5C2.5 18.7614 4.73857 21 7.5 21H17.5C20.2614 21 22.5 18.7614 22.5 16H27.5ZM7.5 26C1.97715 26 -2.5 21.5227 -2.5 16V11H2.5V16C2.5 18.7614 4.73857 21 7.5 21V26ZM27.5 11V16C27.5 21.5227 23.0229 26 17.5 26V21C20.2614 21 22.5 18.7614 22.5 16V11H27.5Z"
                      fill="white"
                    />
                  </g>
                  <mask
                    id="mask1_3107_32"
                    //   style="mask-type:luminance"
                    maskUnits="userSpaceOnUse"
                    x="0"
                    y="0"
                    width="25"
                    height="13"
                  >
                    <path
                      d="M25 12.5H0V12C0 7.7996 0 5.6994 0.81745 4.09507C1.5365 2.68385 2.68385 1.5365 4.09508 0.81745C5.6994 5.96046e-08 7.7996 0 12 0H13C17.2004 0 19.3006 5.96046e-08 20.9049 0.81745C22.3161 1.5365 23.4635 2.68385 24.1826 4.09507C25 5.6994 25 7.7996 25 12V12.5Z"
                      fill="white"
                    />
                  </mask>
                  <g mask="url(#mask1_3107_32)">
                    <path
                      d="M-2.5 7.49976C-2.5 1.97691 1.97715 -2.50024 7.5 -2.50024H17.5C23.0229 -2.50024 27.5 1.97691 27.5 7.49976H22.5C22.5 4.73833 20.2614 2.49976 17.5 2.49976H7.5C4.73857 2.49976 2.5 4.73833 2.5 7.49976H-2.5ZM17.5 -2.50024C23.0229 -2.50024 27.5 1.97691 27.5 7.49976V12.4998H22.5V7.49976C22.5 4.73833 20.2614 2.49976 17.5 2.49976V-2.50024ZM-2.5 12.4998V7.49976C-2.5 1.97691 1.97715 -2.50024 7.5 -2.50024V2.49976C4.73857 2.49976 2.5 4.73833 2.5 7.49976V12.4998H-2.5Z"
                      fill="white"
                    />
                  </g>
                  <path
                    d="M7.98912 9.5H7.5C6.80964 9.5 6.25 10.0596 6.25 10.75V12.779C6.25 13.4693 6.80964 14.029 7.5 14.029H7.98912C8.67948 14.029 9.23913 13.4693 9.23913 12.779V10.75C9.23913 10.0596 8.67948 9.5 7.98912 9.5Z"
                    fill="white"
                  />
                  <path
                    d="M17.4989 9.5H17.0098C16.3194 9.5 15.7598 10.0596 15.7598 10.75V12.779C15.7598 13.4693 16.3194 14.029 17.0098 14.029H17.4989C18.1892 14.029 18.7489 13.4693 18.7489 12.779V10.75C18.7489 10.0596 18.1892 9.5 17.4989 9.5Z"
                    fill="white"
                  />
                </svg>
              </a>
              <button>
                <Tooltip content="Home" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt bg-secondary shadow-sm"
                    href={`/project/${projectId}`}
                  >
                    <HomeIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
              <div className="border border-gray-800 h-px w-full"></div>
              <button>
                <Tooltip content="Code Editor" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    href={`/project/${projectId}/code-editor`}
                  >
                    <CodespacesIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
              <button>
                <Tooltip content="Integrations" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    href={`/project/${projectId}/integrations`}
                  >
                    <PackageIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
              <div className="border border-gray-800 h-px w-full"></div>
              <button className="place-content-center">
                <Tooltip
                  content="Open at Zeabur"
                  arrow={false}
                  placement="right"
                >
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    onClick={openAtZeabur}
                  >
                    <img
                      src="https://cdn-botway.deno.dev/icons/zeabur.svg"
                      width={17}
                    />
                  </a>
                </Tooltip>
              </button>
              <button>
                <Tooltip content="Deployments" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    href={`/project/${projectId}/deployments`}
                  >
                    <ServerIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
              <div className="border border-gray-800 h-px w-full"></div>
              <button>
                <Tooltip content="Logs" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    onClick={openLogs}
                  >
                    <ListUnorderedIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
              <button>
                <Tooltip content="Environment" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    href={`/project/${projectId}/env`}
                  >
                    <SlidersIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
              <button>
                <Tooltip content="Settings" arrow={false} placement="right">
                  <a
                    className="transition-colors duration-200 flex items-center justify-center h-10 w-10 rounded hover:bg-bwdefualt"
                    href={`/project/${projectId}/settings`}
                  >
                    <GearIcon className="fill-white" size={18} />
                  </a>
                </Tooltip>
              </button>
            </ul>
            <ul className="flex flex-col space-y-2">
              <a
                type="button"
                href="/account"
                className="flex border-none place-content-center self-center rounded bg-transparent p-0 outline-none transition-all focus:outline-4 focus:outline-none"
              >
                <span className="relative place-content-center self-center cursor-pointer inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded-lg outline-none transition-all outline-0 text-gray-400 hover:bg-bwdefualt shadow-none focus-visible:outline-none text-xs px-2.5 py-1">
                  <span className="truncate">
                    <div className="py-1">
                      <UserAvatar data={user.email} size={24} />
                    </div>
                  </span>
                </span>
              </a>
            </ul>
          </div>

          <main className="flex flex-col flex-1 w-full overflow-hidden h-screen">
            <div className="flex h-12 max-h-12 overflow-hidden items-center justify-between py-2 px-5 border-b border-gray-800">
              <div className="-ml-2 flex items-center text-sm">
                <span className="flex border-none rounded p-0 outline-none outline-offset-1 transition-all focus:outline-4">
                  <span className="relative inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-gray-200 shadow-none text-xs px-2.5 py-1">
                    <span className="truncate">
                      {user.user_metadata["name"]}
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
                <a
                  href={`/project/${projectId}`}
                  className="text-gray-200 block px-2 py-1 text-xs leading-5 focus:outline-none"
                >
                  {projectName}
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
                <a
                  // href="/ai"
                  className="flex border-gray-800 hover:bg-secondary border rounded p-0 outline-none outline-offset-1 transition-all focus:outline-4"
                >
                  <span className="relative cursor-pointer inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-blue-700 shadow-sm text-xs px-2.5 py-1">
                    <SparkleFillIcon />

                    <span className="hidden font-extrabold text-gray-500 md:block">
                      AI - Soon
                    </span>
                  </span>
                </a>
              </div>
            </div>
            <main
              className={clsx(
                "flex-1 overflow-y-auto max-h-screen",
                grid ? "bg-grid-gray-800/[0.4]" : "",
              )}
            >
              <div
                className={`${
                  !noMargin ? "mx-auto" : ""
                } w-full max-w-7xl space-y-8`}
              >
                {children}
              </div>
            </main>
          </main>
        </div>
      </div>
    </>
  );
};
