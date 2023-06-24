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
import ReactFlow, { Node, useNodesState, ConnectionLineType } from "reactflow";
import NodeCustom from "./nodes/node";
import NodeLoading from "./nodes/node-loading";
import { fetcher } from "@/tools/fetch";
import dagre from "dagre";

const dagreGraph = new dagre.graphlib.Graph();
dagreGraph.setDefaultEdgeLabel(() => ({}));

export const revalidate = 0;

const queryClient = new QueryClient();

const nodeTypes = {
  custom: NodeCustom,
  loading: NodeLoading,
};

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

  let initNodes: Node<any, string | undefined>[] = [];

  const nodeWidth = 172;
  const nodeHeight = 36;

  const getLayoutedElements = (nodes: any, direction = "TB") => {
    const isHorizontal = direction === "LR";
    dagreGraph.setGraph({ rankdir: direction });

    nodes.forEach((node: any) => {
      dagreGraph.setNode(node.id, { width: nodeWidth, height: nodeHeight });
    });

    dagre.layout(dagreGraph);

    nodes.forEach((node: any) => {
      const nodeWithPosition = dagreGraph.node(node.id);

      node.targetPosition = isHorizontal ? "left" : "top";
      node.sourcePosition = isHorizontal ? "right" : "bottom";

      // We are shifting the dagre node position (anchor=center center) to the top left
      // so it matches the React Flow node anchor point (top left).
      node.position = {
        x: (nodeWithPosition.x * (services?.services.length - 2)) / 2,
        y: 2 * (nodeWithPosition.y * services?.services.length),
      };

      return node;
    });

    return { nodes };
  };

  const position = { x: 0, y: 0 };

  services?.services.map((node: any, index: any) => {
    initNodes.push({
      id: `"${index + 1}"`,
      type: "custom",
      data: {
        name: node.node.name,
      },
      position,
    });
  });

  services?.plugins.map((node: any, index: any) => {
    initNodes.push({
      id: `"${index + services?.services.length + 1}"`,
      type: "custom",
      data: {
        name: node.node.name,
      },
      position,
    });
  });

  const initNodesLoading = [
    {
      id: "1",
      type: "loading",
      data: {},
      position: { x: 0, y: 50 },
    },
  ];

  const { nodes: layoutedNodes } = getLayoutedElements(initNodes);

  const Flow = () => {
    const [nodes, setNodes, onNodesChange] = useNodesState(
      servicesIsLoading ? initNodesLoading : layoutedNodes
    );

    return (
      <ReactFlow
        nodes={nodes}
        connectionLineType={ConnectionLineType.SmoothStep}
        onNodesChange={onNodesChange}
        nodeTypes={nodeTypes}
        fitView
        nodesDraggable={false}
        className=""
      />
    );
  };

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
