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
import Link from "next/link";
import {
  ArchiveIcon,
  CheckCircleIcon,
  ClockIcon,
  XCircleIcon,
} from "@primer/octicons-react";
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

  const { data: project, isLoading: projectIsLoading } = useQuery({
    queryKey: ["project"],
    queryFn: fetchProject,
    refetchInterval: 360,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const fetchDeployments = async () => {
    const dys = await fetcher(`/api/deployments?id=${projectId}`, {
      method: "GET",
    });

    return dys;
  };

  const { data: deployments, isLoading: dyIsLoading } = useQuery({
    queryKey: ["dy"],
    queryFn: fetchDeployments,
    refetchInterval: 1,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const status = (deployStatus: any) => {
    switch (deployStatus) {
      case "FAILED":
      case "CRASHED":
        return "text-red-700";

      case "BUILDING":
      case "DEPLOYING":
        return "text-green-700";

      case "RUNNING":
        return "text-blue-700";
    }

    return "text-gray-400";
  };

  const logsURL = async (deploy: any) => {
    const { payload: zeaburProjectId } = await jwtDecrypt(
      project?.zeabur_project_id,
      BW_SECRET_KEY,
    );

    const { payload: zeaburServiceId } = await jwtDecrypt(
      project?.zeabur_service_id,
      BW_SECRET_KEY,
    );

    router.push(
      `https://dash.zeabur.com/projects/${zeaburProjectId.data}/services/${zeaburServiceId.data}/deployments/${deploy.node._id}`,
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
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">{project?.name} Deployments</h1>
          </div>
          <div className="mx-6">
            {dyIsLoading ? (
              <LoadingDots className="fixed inset-0 flex items-center justify-center" />
            ) : deployments ? (
              deployments.length != 0 ? (
                deployments.map((deploy: any) => (
                  <div className="rounded-2xl border border-gray-800 overflow-hidden p-5 min-h-72 mb-6">
                    <header className="flex gap-3 justify-between mb-4">
                      <hgroup>
                        <h2 className="font-medium text-lg !leading-none text-black">
                          <span className={status(deploy.node.status)}>
                            {deploy.node.status}
                          </span>
                        </h2>

                        <h3 className="text-gray-500 mt-1 !leading-tight">
                          {deploy.node.status === "RUNNUNG"
                            ? "The deployment that is live on your production domains."
                            : deploy.node.status === "FAILED"
                              ? "The deployment is failed."
                              : deploy.node.status === "REMOVED"
                                ? "The deployment is removed."
                                : deploy.node.status === "CRASHED"
                                  ? "The deployment is crashed"
                                  : "Waiting..."}
                        </h3>
                      </hgroup>

                      <a
                        className="h-8 px-3.5 text-white rounded-md inline-flex flex-shrink-0 bg-secondary whitespace-nowrap items-center gap-2 transition-colors duration-150 ease-in-out leading-none border border-gray-800 hover:border-gray-700 cursor-pointer"
                        onClick={() => logsURL(deploy)}
                      >
                        Logs
                      </a>
                    </header>

                    <label className="flex items-center mt-5 mb-1 text-sm text-gray-400">
                      Deployment Details
                    </label>

                    <div className="flex items-center gap-3 mt-2">
                      <span className="w-5 h-5 inline-flex items-center justify-center rounded-full flex-shrink-0 bg-fresh/15">
                        {deploy.node.status === "RUNNING" ? (
                          <CheckCircleIcon
                            className="fill-green-700"
                            size={16}
                          />
                        ) : deploy.node.status != "RUNNING" ? (
                          deploy.node.status === "REMOVED" ? (
                            <ArchiveIcon className="fill-gray-400" size={16} />
                          ) : deploy.node.status === "FAILED" ||
                            deploy.node.status === "CRASHED" ? (
                            <XCircleIcon className="fill-red-700" size={16} />
                          ) : (
                            <ClockIcon className="fill-gray-400" size={16} />
                          )
                        ) : (
                          <ArchiveIcon className="fill-gray-400" size={16} />
                        )}
                      </span>

                      <span className="flex items-center gap-1">
                        <img
                          src="https://cdn-botway.deno.dev/icons/docker.svg"
                          width={18}
                          className="mr-1"
                        />
                      </span>

                      <span className="inline-flex items-center gap-2 max-w-100">
                        <Link
                          className="text-gray-400 text-sm hover:text-gray-500 transition-all duration-200 hover:underline truncate"
                          href={`https://github.com/${deploy.node.repoOwner}/${deploy.node.repoName}/commit/${deploy.node.commitSHA}`}
                          target="_blank"
                          title={deploy.node.commitMessage}
                        >
                          {deploy.node.commitMessage}
                        </Link>
                      </span>
                    </div>
                  </div>
                ))
              ) : (
                <div className="rounded-2xl overflow-hidden p-5 border-2 border-dashed border-gray-800 shadow-lg w-full h-60 flex flex-col items-center justify-center gap-4">
                  <h2 className="text-md text-gray-400 text-center">
                    Your project has no deploys
                  </h2>
                </div>
              )
            ) : (
              <div className="rounded-2xl overflow-hidden p-5 border-2 border-dashed border-gray-800 shadow-lg w-full h-60 flex flex-col items-center justify-center gap-4">
                <h2 className="text-md text-gray-400 text-center">
                  Your project has no deploys
                </h2>
              </div>
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
