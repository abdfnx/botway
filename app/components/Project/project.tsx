import { fetcher } from "@/lib/fetch";
import { bgSecondary } from "@/tools/colors";
import { Tab } from "@headlessui/react";
import clsx from "clsx";
import { useCallback, useEffect, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { Button } from "../Button";

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

export const ProjectMain = ({ project, mutate }: any) => {
  let [navs] = useState(["Overview", "Config", "Deployments", "Settings"]);

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
              <Content nav={nav} project={project} mutate={mutate} />
            </Tab.Panel>
          ))}
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
};

const Content = ({ nav, project, mutate }: any) => {
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
            <h3 className="text-lg font-medium leading-6 text-gray-400">
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
                    <dd className="mt-1 text-sm text-gray-400 sm:col-span-2 sm:mt-0">
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
  } else if (nav == "Config") {
    const capitalizeFirstLetter = (text: String) => {
      return text.charAt(0).toUpperCase() + text.slice(1);
    };

    const [isLoading, setIsLoading] = useState(false);

    const botTokenRef: any = useRef();
    const botAppTokenRef: any = useRef();
    const botSecretTokenRef: any = useRef();

    const onSubmit = useCallback(
      async (e: any) => {
        e.preventDefault();

        try {
          setIsLoading(true);

          const formData = new FormData();

          formData.append("id", project._id);
          formData.append("name", project.name);
          formData.append("platform", project.platform);
          formData.append("botToken", botTokenRef.current.value);

          if (project.platform != "telegram") {
            formData.append("botAppToken", botAppTokenRef.current.value);
          }

          if (project.platform == "slack" || project.platform == "twitch") {
            formData.append("botSecretToken", botSecretTokenRef.current.value);
          }

          await fetcher("/api/projects/update", {
            method: "PATCH",
            body: formData,
          });

          toast.success("Your project has been updated", {
            style: {
              borderRadius: "10px",
              backgroundColor: bgSecondary,
              color: "#fff",
            },
          });
        } catch (e: any) {
          toast.error(e.message, {
            style: {
              borderRadius: "10px",
              backgroundColor: bgSecondary,
              color: "#fff",
            },
          });
        } finally {
          setIsLoading(false);
        }
      },
      [mutate]
    );

    useEffect(() => {
      botTokenRef.current.value = project.botToken;
      botAppTokenRef.current.value = project.botAppToken;
      botSecretTokenRef.current.value = project.botSecretToken;
    }, [project]);

    return (
      <div className="overflow-hidden sm:rounded-lg">
        <div className="px-4 py-5 sm:px-6">
          <h3 className="text-lg font-medium leading-6 text-gray-400">
            Bot Configuration
          </h3>
        </div>
        <form onSubmit={onSubmit}>
          <div className="grid lg:grid-cols-2 sm:grid-cols-1 lt-md:!grid-cols-1 gap-3">
            <div className="px-4 py-5 sm:px-6">
              <label
                htmlFor={`${project.platform}-bot-token`}
                className="block text-gray-500 text-sm font-semibold"
              >
                {capitalizeFirstLetter(project.platform)} Bot Token
              </label>
              <div className="pt-2">
                <input
                  className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                  ref={botTokenRef}
                  type="password"
                  required
                />
              </div>
            </div>

            {project.platform != "telegram" ? (
              <div className="px-4 py-5 sm:px-6">
                <label
                  htmlFor={`${project.platform}-app-id`}
                  className="block text-gray-500 text-sm font-semibold"
                >
                  {capitalizeFirstLetter(project.platform)}{" "}
                  {project.platform != "twitch"
                    ? `Bot App ${
                        project.platform == "discord" ? "ID" : "Token"
                      }`
                    : "Bot Client ID"}
                </label>
                <div className="pt-2">
                  <input
                    className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                    ref={botAppTokenRef}
                    type="password"
                    required
                  />
                </div>
              </div>
            ) : (
              <input ref={botAppTokenRef} hidden />
            )}

            {project.platform == "slack" || project.platform == "twitch" ? (
              <div className="px-4 py-5 sm:px-6">
                <label
                  htmlFor={`${project.platform}-app-id`}
                  className="block text-gray-500 text-sm font-semibold"
                >
                  {capitalizeFirstLetter(project.platform)}{" "}
                  {project.platform == "twitch"
                    ? "Bot Client Secret"
                    : "Bot Signing Secret"}
                </label>
                <div className="pt-2">
                  <input
                    className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                    ref={botSecretTokenRef}
                    type="password"
                    required
                  />
                </div>
              </div>
            ) : (
              <input ref={botSecretTokenRef} hidden />
            )}
          </div>

          <div className="mb-6 space-y-2 flex justify-center">
            <Button
              type="success"
              htmlType="submit"
              loading={isLoading}
              className="button w-full p-2"
            >
              Update
            </Button>
          </div>
        </form>
      </div>
    );
  }

  return <></>;
};
