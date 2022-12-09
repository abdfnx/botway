import { Tab } from "@headlessui/react";
import clsx from "clsx";
import { useState } from "react";

const InfoIcon = ({ value }: any) => {
  let iconURL;

  if (value == "Render") {
    iconURL = "render.png";
  } else if (value == "default") {
    iconURL = "c.svg";
  } else if (value == "shards") {
    iconURL = "crystal.svg";
  } else if (value == "pub") {
    iconURL = "dart.svg";
  } else if (value == "deno package manager") {
    iconURL = "deno.svg";
  } else if (value == "go package manager") {
    iconURL = "go.svg";
  } else if (value == "pipenv") {
    iconURL = "pipenv.png";
  } else if (value == "bundler") {
    iconURL = "bundler.png";
  } else if (value == "cargo") {
    iconURL = "cargo.png";
  } else if (value == "fleet") {
    iconURL = "rust.svg";
  } else {
    iconURL = value.toLowerCase() + ".svg";
  }

  return (
    <img src={`https://cdn-botway.deno.dev/icons/${iconURL}`} width={20} />
  );
};

export const ProjectMain = ({ project }: any) => {
  let [navs] = useState([
    "Overview",
    "Config",
    "Vars",
    "Deployments",
    "Settings",
  ]);

  return (
    <div className="w-full max-w-full max-h-md px-1 py-1 sm:py-1">
      <Tab.Group>
        <Tab.List className="flex space-x-1 border border-gray-800 rounded-md p-2">
          {navs.map((nav) => (
            <Tab
              key={nav}
              className={({ selected }) =>
                clsx(
                  "w-full rounded-lg transition py-2.5 text-sm font-medium outline-none leading-5 text-blue-700",
                  selected
                    ? "bg-secondary shadow"
                    : "text-blue-100 hover:bg-secondary/70 hover:text-white"
                )
              }
            >
              {nav}
            </Tab>
          ))}
        </Tab.List>
        <Tab.Panels className="mt-2">
          {navs.map((nav) => (
            <Tab.Panel
              key={nav}
              className="rounded-xl bg-secondary outline-none p-3"
            >
              <Content nav={nav} project={project} />
            </Tab.Panel>
          ))}
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
};

const Content = ({ nav, project }: any) => {
  if (nav == "Overview") {
    const elements = [
      {
        title: "Bot Name",
        value: project.name,
        icon: false,
      },
      {
        title: "Platform",
        value: project.platform,
        icon: true,
      },
      {
        title: "Bot Programming Language",
        value: project.lang,
        icon: true,
      },
      {
        title: "Package Manager",
        value: project.packageManager,
        icon: true,
      },
      {
        title: "Host Service",
        value: project.hostService,
        icon: true,
      },
    ];

    return (
      <>
        <div className="overflow-hidden sm:rounded-lg">
          <div className="px-4 py-5 sm:px-6">
            <h3 className="text-lg font-medium leading-6 text-gray-900">
              Bot Information
            </h3>
          </div>
          <div>
            <dl>
              {elements.map((e) => (
                <>
                  <div className="px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-400">
                      {e.title}
                    </dt>
                    <dd className="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                      {e.icon ? (
                        <>
                          <div className="flex items-center">
                            <InfoIcon
                              className="h-6 w-6 flex-shrink-0"
                              value={e.value}
                            />
                            <span className="font-semibold ml-2 block truncate">
                              {e.value}
                            </span>
                          </div>
                        </>
                      ) : (
                        e.value
                      )}
                    </dd>
                  </div>
                </>
              ))}
            </dl>
          </div>
        </div>
      </>
    );
  }

  return <></>;
};
