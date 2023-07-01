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
import clsx from "clsx";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";

export const revalidate = 0;

const queryClient = new QueryClient();

const Project = ({ user, projectId }: any) => {
  const router = useRouter();

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

  const fetchLogs = async () => {
    const logs = await fetcher(`/api/deployments/logs?id=${projectId}`, {
      method: "GET",
    });

    return logs;
  };

  const { data: logs, isLoading: logsIsLoading } = useQuery(
    ["logs"],
    fetchLogs,
    {
      refetchInterval: 1,
      refetchOnReconnect: true,
      refetchOnWindowFocus: true,
      refetchIntervalInBackground: true,
    }
  );

  const openAtRailway = async () => {
    const { payload: railwayProjectId } = await jwtDecrypt(
      project?.railway_project_id,
      BW_SECRET_KEY
    );

    const { payload: railwayServiceId } = await jwtDecrypt(
      project?.railway_service_id,
      BW_SECRET_KEY
    );

    router.push(
      `https://railway.app/project/${railwayProjectId.data}/service/${railwayServiceId.data}?id=${logs.dyId}`
    );
  };

  const NoLogs = () => {
    return (
      <div className="rounded-2xl overflow-hidden p-5 w-full h-60 flex flex-col items-center justify-center gap-4">
        <h2 className="text-md text-gray-400 text-center">
          There is no logs 📜
        </h2>
      </div>
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
          projectRWID={project?.railway_project_id}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">{project?.name} Deploy Logs</h1>

            <button
              onClick={openAtRailway}
              className="border border-gray-800 transition-all bg-[#181622] hover:bg-[#1f132a] duration-200 rounded-2xl p-3 text-white flex flex-col items-center"
            >
              <span className="flex">
                <img
                  src="https://cdn-botway.deno.dev/icons/railway.svg"
                  width={24}
                />
                <span className="ml-2">Open Logs at Railway</span>
              </span>
            </button>
          </div>

          <div className="mx-6">
            <div className="rounded-md bg-secondary border border-gray-800 overflow-auto p-5 max-h-[400px] mb-6">
              {logsIsLoading ? (
                <LoadingDots />
              ) : logs.message != "No Logs" ? (
                logs.logs.length != 0 ? (
                  logs.logs.map((deploy: any) => (
                    <div>
                      <p className="font-mono text-xs text-white whitespace-nowrap">
                        <span
                          className={clsx(
                            "pr-2",
                            deploy.severity === "err"
                              ? "text-red-600"
                              : "text-blue-700"
                          )}
                        >
                          {deploy.severity}
                        </span>

                        <span className="pr-2 text-gray-400">
                          {deploy.timestamp}
                        </span>

                        {deploy.message}
                      </p>
                    </div>
                  ))
                ) : (
                  <NoLogs />
                )
              ) : (
                <NoLogs />
              )}
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
