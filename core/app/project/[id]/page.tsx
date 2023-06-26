"use client";

import { redirect } from "next/navigation";
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
  CheckIcon,
  GearIcon,
  MarkGithubIcon,
  XCircleIcon,
} from "@primer/octicons-react";
import Link from "next/link";
import { CheckTokens } from "./settings/page";
import { Tooltip } from "flowbite-react";

export const revalidate = 0;

const queryClient = new QueryClient();

const Project = ({ user, projectId }: any) => {
  const fetchServices = async () => {
    const services = await fetcher(`/api/projects/services`, {
      method: "POST",
      body: JSON.stringify({
        projectId,
      }),
    });

    return services;
  };

  const { data: services, isLoading: servicesIsLoading } = useQuery(
    ["services"],
    fetchServices,
    {
      refetchInterval: 1,
      refetchOnReconnect: true,
      refetchOnWindowFocus: true,
      refetchIntervalInBackground: true,
    }
  );

  services?.services.map((node: any, index: any) => {
    // initNodes.push({
    //   id: `"${index + 1}"`,
    //   type: "custom",
    //   data: {
    //     name: node.node.name,
    //   },
    //   position,
    // });
  });

  services?.plugins.map((node: any, index: any) => {});

  const fetchProject = async () => {
    const { data: project } = await supabase
      .from("projects")
      .select("*")
      .eq("id", projectId)
      .single();

    return project;
  };

  const { data: project, isLoading: projectIsLoading } = useQuery(
    ["project"],
    fetchProject,
    {
      refetchInterval: 1,
      refetchOnReconnect: true,
      refetchOnWindowFocus: true,
      refetchIntervalInBackground: true,
    }
  );

  return (
    <>
      {projectIsLoading ? (
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : (
        <ProjectLayout
          user={user}
          projectId={projectId}
          projectName={project?.name}
          projectRWID={project?.railway_project_id}
          grid={true}
        >
          <div className="mx-6 mt-16 flex items-center space-x-6">
            <div className="">
              <img src="https://cdn-botway.deno.dev/icons/bot.svg" width={67} />
            </div>
            <div>
              <h1 className="text-lg text-white">{project?.name}</h1>
              <h1 className="text-base text-gray-400">Bot Project</h1>
            </div>
          </div>
          <div className="mx-6 bg-secondary justify-between flex border border-gray-800 rounded-lg p-4">
            <div className="flex">
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

              <a href={`/project/${projectId}/settings`}>
                <GearIcon size={20} className="ml-3 fill-white" />
              </a>
            </div>
          </div>
          <div className="mx-6"></div>
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
