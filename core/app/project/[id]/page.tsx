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
import ReactFlow, { useNodesState } from "reactflow";
import Node from "./node";

export const revalidate = 0;

const queryClient = new QueryClient();

const nodeTypes = {
  custom: Node,
};

const initNodes = [
  {
    id: "1",
    type: "custom",
    data: { name: "Test1" },
    position: { x: 0, y: 50 },
  },
];

const Flow = () => {
  const [nodes, setNodes, onNodesChange] = useNodesState(initNodes);

  return (
    <ReactFlow
      nodes={nodes}
      onNodesChange={onNodesChange}
      nodeTypes={nodeTypes}
      fitView
      nodesDraggable={false}
      className=""
    />
  );
};

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
          noMargin={true}
        >
          <div className="w-screen h-screen">
            <Flow />
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
