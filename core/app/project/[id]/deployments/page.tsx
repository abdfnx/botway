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
import useSWR from "swr";
import { fetcher } from "@/tools/fetch";
import Link from "next/link";
import {
  ArchiveIcon,
  CheckCircleIcon,
  ClockIcon,
  FileDirectoryIcon,
  GitMergeIcon,
  XCircleIcon,
} from "@primer/octicons-react";

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

  const { data: project, isLoading: projectIsLoading } = useQuery(
    ["project"],
    fetchProject,
    {
      refetchInterval: 1,
    }
  );

  const deploymentsFetcher = (url: any) =>
    fetcher(url, {
      method: "GET",
    });

  const { data, error } = useSWR(
    `/api/deployments?id=${projectId}`,
    deploymentsFetcher,
    {
      refreshWhenOffline: true,
      refreshWhenHidden: false,
      refreshInterval: 0,
    }
  );

  if (!data && !error)
    return (
      <LoadingDots className="fixed inset-0 flex items-center justify-center" />
    );

  console.log(data);

  const status = (deployStatus: any) => {
    switch (deployStatus) {
      case "FAILED":
        return "text-red-700";

      case "SUCCESS":
        return "text-green-700";
    }

    return "text-gray-400";
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
          // grid={true}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">{project?.name} Deployments</h1>
          </div>
          <div className="mx-6">
            {data.length != 0 ? (
              data.map((deploy: any) => (
                <div className="rounded-2xl border border-gray-800 overflow-hidden p-5 bg-ultralight min-h-72 mb-6">
                  <header className="flex gap-3 justify-between mb-4">
                    <hgroup>
                      <h2 className="font-medium text-lg !leading-none text-black">
                        {deploy.node.url ? (
                          <Link href={deploy.node.url} target="_blank">
                            {deploy.node.url}
                          </Link>
                        ) : (
                          <span className={status(deploy.node.status)}>
                            {deploy.node.status}
                          </span>
                        )}
                      </h2>
                      <h3 className="text-gray-500 mt-1 !leading-tight">
                        {deploy.node.status == "SUCCESS"
                          ? "The deployment that is live on your production domains."
                          : deploy.node.status == "FAILED"
                          ? "The deployment is failed."
                          : deploy.node.status == "REMOVED"
                          ? "The deployment is removed."
                          : "Waiting..."}
                      </h3>
                    </hgroup>
                    <Link
                      target="_blank"
                      className="h-8 px-3.5 text-white rounded-md inline-flex flex-shrink-0 bg-secondary whitespace-nowrap items-center gap-2 transition-colors duration-150 ease-in-out leading-none border border-gray-800 hover:border-gray-700 cursor-pointer"
                      href={`/projects/${projectId}/logs`}
                      aria-current="page"
                    >
                      Logs
                    </Link>
                  </header>

                  <label className="flex items-center mt-5 mb-1 text-sm text-gray-400">
                    Deployment Details
                  </label>
                  <div className="flex items-center gap-3 mt-2">
                    <span className="w-5 h-5 inline-flex items-center justify-center rounded-full flex-shrink-0 bg-fresh/15">
                      {deploy.node.status === "SUCCESS" ? (
                        <CheckCircleIcon className="fill-green-700" size={16} />
                      ) : deploy.node.status != "SUCCESS" ? (
                        deploy.node.status == "REMOVED" ? (
                          <ArchiveIcon className="fill-red-700" size={16} />
                        ) : deploy.node.status == "FAILED" ? (
                          <XCircleIcon className="fill-red-700" size={16} />
                        ) : (
                          <ClockIcon className="fill-gray-400" size={16} />
                        )
                      ) : (
                        <ArchiveIcon className="fill-red-700" size={16} />
                      )}
                    </span>
                    <span className="flex items-center gap-1">
                      <img
                        src="https://cdn-botway.deno.dev/icons/docker.svg"
                        width={18}
                        className="mr-1"
                      />
                    </span>
                    <span className="flex items-center gap-1">
                      <FileDirectoryIcon
                        size={16}
                        className="fill-gray-600 mr-1 font-mono"
                      />
                      {deploy.node.meta.rootDirectory}
                    </span>
                    <span className="flex items-center gap-1">
                      <GitMergeIcon size={16} className="fill-gray-600 mr-1" />
                      <span className="text-gray-400">
                        {deploy.node.meta.branch}
                      </span>
                    </span>
                    <span className="inline-flex items-center gap-2 max-w-100">
                      <Link
                        className="text-gray-400 text-sm hover:text-gray-500 transition-all duration-200 hover:underline truncate"
                        href={`https://github.com/${deploy.node.meta.repo}/commit/${deploy.node.meta.commitHash}`}
                        target="_blank"
                        title={deploy.node.meta.commitMessage}
                      >
                        {deploy.node.meta.commitMessage}
                      </Link>
                    </span>
                  </div>
                </div>
              ))
            ) : (
              <div className="rounded-2xl overflow-hidden p-5 cursor-pointer border-2 border-dashed border-gray-800 hover:border-gray-600 shadow-lg transition duration-300 ease-in-out w-full h-60 flex flex-col items-center justify-center gap-4">
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
