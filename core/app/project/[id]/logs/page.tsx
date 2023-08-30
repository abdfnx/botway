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
  ApolloClient,
  ApolloLink,
  ApolloProvider,
  HttpLink,
  InMemoryCache,
  concat,
  gql,
  split,
  useSubscription,
} from "@apollo/client";
import { GraphQLWsLink } from "@apollo/client/link/subscriptions";
import { createClient } from "graphql-ws";
import { PropsWithChildren, useState } from "react";
import { getMainDefinition } from "@apollo/client/utilities";

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
      refetchOnReconnect: true,
      refetchOnWindowFocus: true,
      refetchIntervalInBackground: true,
    },
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
    },
  );

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
          projectRWID={project?.zeabur_project_id}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">{project?.name} Deploy Logs</h1>
          </div>

          <div className="mx-6">
            <div className="rounded-md bg-secondary border border-gray-800 overflow-auto p-5 max-h-[400px] mb-6">
              {/* {logsIsLoading ? (
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
                              : "text-blue-700",
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
              )} */}
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
