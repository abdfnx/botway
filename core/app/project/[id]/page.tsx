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
import { ContainerIcon } from "@primer/octicons-react";

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
          <div className="mx-6 my-16 flex items-center space-x-6">
            <div className="">
              <img src="https://cdn-botway.deno.dev/icons/bot.svg" width={67} />
            </div>
            <div>
              <h1 className="text-lg text-white">{project?.name}</h1>
              <h1 className="text-base text-gray-400">Bot Project</h1>
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
