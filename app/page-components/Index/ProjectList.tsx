import { Project } from "@/components/Project";
import { useProjectPages } from "@/lib/project";
import { NewProjectButton } from ".";
import { NewProject } from "../New/NewProject";

export const ProjectList = () => {
  const { data, isReachingEnd } = useProjectPages();
  const projects = data
    ? data.reduce((acc, val) => [...acc, ...val.projects], [])
    : [];

  if (isReachingEnd) {
    return (
      <div className="rounded-lg mt-8 overflow-hidden p-5 cursor-pointer border-2 border-dashed border-gray-800 hover:border-gray-600 shadow-lg transition-shadow duration-500 ease-in-out w-full h-60 flex flex-col items-center justify-center gap-4">
        <h2 className="text-md text-gray-400 text-center">
          Create a New Project
        </h2>
        <NewProject />
      </div>
    );
  }

  return (
    <>
      {projects.map((project: any) => (
        <div className="mt-10 grid lg:grid-cols-3 sm:grid-cols-2 lt-md:!grid-cols-1 gap-3">
          <Project project={project} />
        </div>
      ))}
    </>
  );
};
