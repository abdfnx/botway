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
import { IntegrationsGird } from "@/components/Grid/IntegrationsGird";
import { fetcher } from "@/tools/fetch";

export const revalidate = 0;

const queryClient = new QueryClient();

const Project = ({ user, projectId }: any) => {
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
    refetchInterval: 360,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const integrationsByCategory: { [category: string]: any } = {};

  const fetchIntegrations = async () => {
    const integrations = await fetcher(`/api/integrations`, {
      method: "GET",
    });

    return integrations;
  };

  const { data: integrations, isLoading: integrationsIsLoading } = useQuery({
    queryKey: ["integrations"],
    queryFn: fetchIntegrations,
    refetchInterval: 1,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  integrations?.forEach(
    (i: any) =>
      (integrationsByCategory[i.category] = [
        ...(integrationsByCategory[i.category] ?? []),
        i,
      ]),
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
          projectZBID={project?.zeabur_project_id}
          grid={true}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">Integrations</h1>
          </div>

          <div className="mx-6">
            {integrationsIsLoading ? (
              <LoadingDots className="fixed inset-0 flex items-center justify-center" />
            ) : (
              <IntegrationsGird
                integrationsByCategory={integrationsByCategory}
                projectId={projectId}
              />
            )}
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
