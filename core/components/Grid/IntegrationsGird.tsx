export const IntegrationsGird = ({
  integrationsByCategory,
  projectId,
}: {
  integrationsByCategory: { [category: string]: any };
  projectId: string;
}) => {
  return (
    <>
      {Object.keys(integrationsByCategory)
        .sort()
        .map((category) => (
          <div
            key={category}
            id={category.toLowerCase()}
            className="space-y-6 my-6"
          >
            {<h2 className="text-xl text-gray-400">{category}</h2>}

            <div className="grid gap-5 grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 lg:max-w-none">
              {integrationsByCategory[category].map((p: any) => (
                <a
                  className="focus:outline-none relative h-44 flex flex-col px-6 pt-6 pb-4 bg-bwdefualt hover:bg-secondary border border-gray-800 transition-all duration-200 rounded-2xl"
                  href={`/project/${projectId}/integrations/${p.slug}`}
                  key={p.slug}
                >
                  <div className="relative">
                    <img
                      width={20}
                      height={20}
                      className="w-10 h-10"
                      src={`https://cdn-botway.deno.dev/icons/${p.slug}.svg`}
                      alt={p.name}
                    />
                  </div>

                  <div className="mt-4 flex-grow flex flex-col space-y-2">
                    <p className="font-semibold text-white line-clamp-1">
                      {p.name}{" "}
                      {p.soon ? (
                        <>
                          - <span className="text-blue-700">Soon</span>
                        </>
                      ) : null}
                    </p>
                    <p className="text-sm line-clamp-2 text-gray-500">
                      {p.desc}
                    </p>
                  </div>
                </a>
              ))}
            </div>
          </div>
        ))}
    </>
  );
};
