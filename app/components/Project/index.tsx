export const Project = ({ project }: any) => {
  return (
    <div className="flex items-center justify-between gap-3 px-5 py-0 rounded-lg border-2 border-dashed border-gray-800 hover:bg-secondary transition-colors duration-200">
      <a
        className="block flex-1 py-5"
        href={`/project/${project._id}`}
        aria-current="page"
      >
        <h2>
          <strong className="text-base leading-tight font-medium align-middle">
            {project.name}
          </strong>
        </h2>
        <br />
        <p className="flex items-center gap-1.5 mt-1.5 text-sm text-gray-500">
          <img
            src={`https://cdn-botway.deno.dev/icons/${project.platform}.svg`}
            alt={`${project.platform} icon`}
            width={16}
          />
          {project.platform}

          <img
            src={`https://cdn-botway.deno.dev/icons/${project.lang}.svg`}
            alt={`${project.lang} icon`}
            width={16}
          />
          {project.lang}
        </p>
      </a>
    </div>
  );
};
