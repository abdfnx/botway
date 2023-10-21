"use client";

import { redirect, useRouter } from "next/navigation";
import { useAuth } from "@/supabase/auth/provider";
import { LoadingDots } from "@/components/LoadingDots";
import supabase from "@/supabase/browser";
import { ProjectLayout } from "@/components/Layout/project";
import {
  useQuery,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { fetcher } from "@/tools/fetch";
import {
  ArrowRightIcon,
  CheckIcon,
  DatabaseIcon,
  GearIcon,
  MarkGithubIcon,
  XCircleIcon,
} from "@primer/octicons-react";
import Link from "next/link";
import { CheckTokens } from "./settings/page";
import { Tooltip } from "flowbite-react";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { capitalizeFirstLetter } from "@/tools/text";

export const revalidate = 0;

const queryClient = new QueryClient();

const Project = ({ user, projectId }: any) => {
  const router = useRouter();

  const data: { _id: string; name: any; code: any; icon: any }[] = [];

  const fetchLatestDeployment = async () => {
    const dy = await fetcher(`/api/deployments/latest?id=${projectId}`, {
      method: "GET",
    });

    return dy;
  };

  const fetchServices = async () => {
    const services = await fetcher(`/api/projects/services`, {
      method: "POST",
      body: JSON.stringify({
        projectId,
      }),
    });

    return services;
  };

  const { data: services, isLoading: servicesIsLoading } = useQuery({
    queryKey: ["services"],
    queryFn: fetchServices,
    refetchInterval: 36000,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const { data: dy } = useQuery({
    queryKey: ["dy"],
    queryFn: fetchLatestDeployment,
    refetchOnReconnect: true,
  });

  console.log(dy);

  services?.services.map((node: any) => {
    data.push({
      _id: node._id,
      name: node.name,
      code: node?.marketplaceItem?.code || "",
      icon: node?.marketplaceItem?.iconURL || "",
    });
  });

  const fetchProject = async () => {
    const { data: project } = await supabase
      .from("projects")
      .select("*")
      .eq("id", projectId)
      .single();

    return project;
  };

  const { data: project, isLoading: projectIsLoading } = useQuery({
    queryKey: ["project"],
    queryFn: fetchProject,
    refetchInterval: 1,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const openAtZeabur = async (id: any) => {
    const { payload: zeaburProjectId } = await jwtDecrypt(
      project?.zeabur_project_id,
      BW_SECRET_KEY,
    );

    router.push(
      `https://dash.zeabur.com/projects/${zeaburProjectId.data}/services/${id}`,
    );
  };

  return (
    <>
      {projectIsLoading ? (
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : (
        <ProjectLayout
          user={user}
          projectId={projectId}
          projectName={project?.name}
          projectZBID={project?.zeabur_project_id}
          projectServiceID={project?.zeabur_service_id}
          latestDeployment={dy}
          grid={true}
        >
          <h1 className="mx-6 my-16 text-3xl text-white">My Bot</h1>

          <div className="mx-6 mt-16 flex items-center space-x-6">
            <img src="https://cdn-botway.deno.dev/icons/bot.svg" width={55} />

            <div>
              <h1 className="text-base text-white">{project?.name}</h1>
              <h1 className="text-sm text-gray-400">
                {capitalizeFirstLetter(project?.platform)} Bot Project
              </h1>
            </div>
          </div>

          <div className="mx-6 mt-1 bg-secondary justify-between flex border border-gray-800 rounded-lg p-4">
            <div className="flex mt-0.5">
              <MarkGithubIcon size={20} className="fill-gray-400" />

              <Link
                href={`https://github.com/${project?.repo}`}
                target="_blank"
              >
                <h1 className="pl-2 text-sm text-white">{project?.repo}</h1>
              </Link>
            </div>

            <div className="flex">
              <button>
                <Tooltip content="Tokens Status" arrow={false} placement="top">
                  {CheckTokens(project) ? (
                    <>
                      <CheckIcon size={20} className="fill-green-600" />
                    </>
                  ) : (
                    <>
                      <XCircleIcon size={20} className="fill-red-600" />
                    </>
                  )}
                </Tooltip>
              </button>

              <Tooltip content="Bot Settings" arrow={false} placement="top">
                <a href={`/project/${projectId}/settings`}>
                  <GearIcon size={20} className="ml-3 fill-white" />
                </a>
              </Tooltip>
            </div>
          </div>
          <div className="mx-6">
            <div className="my-6">
              <h3 className="text-white text-xl">Containers</h3>
            </div>

            <div className="my-4 max-w-full space-y-8">
              <div className="overflow-x-auto flex-grow rounded-lg border border-gray-800">
                <table className="w-full border-collapse select-auto bg-secondary">
                  <tbody>
                    {servicesIsLoading ? (
                      <tr>
                        <td className="py-3 px-4  place-content-center overflow-hidden items-center justify-center overflow-ellipsis whitespace-nowrap">
                          <div className="flex space-x-2 items-center">
                            <LoadingDots className="flex items-center justify-center" />
                          </div>
                        </td>
                      </tr>
                    ) : (
                      data.map((node: any) => (
                        <tr
                          className={`${
                            node._id % 2 === 0 ? "bg-secondary" : ""
                          }`}
                        >
                          <td
                            className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap border-r border-gray-800"
                            style={{ minWidth: "64px", maxWidth: "100px" }}
                          >
                            <div className="flex space-x-2 items-center">
                              {node?.code ? (
                                <img src={node?.icon} width={20} />
                              ) : (
                                <MarkGithubIcon
                                  className="fill-white"
                                  size={19}
                                />
                              )}
                            </div>
                          </td>

                          <td
                            className="py-3 px-4 overflow-hidden hidden md:table-cell overflow-ellipsis whitespace-nowrap"
                            style={{ minWidth: "64px", maxWidth: "250px" }}
                          >
                            <div className="flex space-x-2 mt-0.5 items-center">
                              <p className="text-sm text-white">{node.name}</p>
                            </div>
                          </td>

                          <td
                            className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500"
                            style={{ minWidth: "64px", maxWidth: "400px" }}
                          >
                            <span className="inline-flex items-center rounded-full bg-blue-50 px-2 py-0.5 text-xs font-medium text-blue-700 ring-1 ring-inset ring-blue-700/10">
                              {node?.code ? "plugin" : "service"}
                            </span>
                          </td>

                          <td
                            className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500"
                            style={{ minWidth: "64px", maxWidth: "400px" }}
                          >
                            <Tooltip
                              content="Open at Zeabur"
                              arrow={false}
                              placement="right"
                            >
                              <img
                                src="https://cdn-botway.deno.dev/icons/zeabur.svg"
                                className="cursor-pointer"
                                onClick={() => openAtZeabur(node._id)}
                                width={20}
                              />
                            </Tooltip>
                          </td>
                        </tr>
                      ))
                    )}
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          <div className="mx-6">
            <div className="my-6">
              <h3 className="text-white text-xl">Infrastructure</h3>
            </div>

            <div className="overflow-hidden shadow pb-12">
              <div className="flex flex-col gap-0">
                <div className="grid lg:grid-cols-2 sm:grid-cols-2 lt-md:!grid-cols-1 gap-3">
                  <a
                    href="https://docker.com"
                    target="_blank"
                    className="border border-gray-800 transition-all bg-[#00084d] hover:bg-[#00124d] duration-200 rounded-2xl p-4 flex flex-col items-center"
                  >
                    <div aria-hidden="true">
                      <img
                        src="https://cdn-botway.deno.dev/icons/docker.svg"
                        width={35}
                      />
                    </div>

                    <div className="space-y-2 mt-3 sm:space-y-4 flex flex-col items-center">
                      <h1 className="text-white text-xs md:text-sm font-bold">
                        Docker is your Container Builder
                      </h1>

                      <p className="text-xs md:text-sm text-gray-400 text-center">
                        Docker is a platform for developing, shipping, and
                        running applications 🐳
                      </p>
                    </div>
                  </a>

                  <a
                    href="https://zeabur.com"
                    target="_blank"
                    className="border border-gray-800 transition-all bg-[#121212] hover:bg-[#141414] duration-200 rounded-2xl p-4 flex flex-col items-center"
                  >
                    <div aria-hidden="true">
                      <img
                        src="https://cdn-botway.deno.dev/icons/zeabur.svg"
                        width={30}
                      />
                    </div>

                    <div className="space-y-2 mt-3 sm:space-y-4 flex flex-col items-center">
                      <h1 className="text-white text-xs md:text-sm font-bold">
                        Zeabur is your Host Service
                      </h1>
                      <p className="text-xs md:text-sm text-gray-400 text-center">
                        Zeabur is a platform that help you deploy your service
                        with one click. 🏗️
                      </p>
                    </div>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </ProjectLayout>
      )}
    </>
  );
};

const ProjectPage = ({ params }: any) => {
  const { initial, user } = useAuth();

  if (initial) {
    return (
      <LoadingDots className="fixed inset-0 flex items-center justify-center" />
    );
  }

  if (user) {
    return (
      <QueryClientProvider client={queryClient}>
        <Project user={user} projectId={params.id} />
      </QueryClientProvider>
    );
  }

  redirect("/");
};

export default ProjectPage;
