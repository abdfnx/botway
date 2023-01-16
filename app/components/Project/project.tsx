import { fetcher } from "@/lib/fetch";
import { BW_SECRET_KEY, CheckAPITokens } from "@/tools/api-tokens";
import { toastStyle } from "@/tools/toast-style";
import { Dialog, Tab, Transition } from "@headlessui/react";
import {
  AlertIcon,
  ArchiveIcon,
  CheckCircleIcon,
  ClockIcon,
  FileDirectoryIcon,
  GitMergeIcon,
  XCircleIcon,
} from "@primer/octicons-react";
import clsx from "clsx";
import Link from "next/link";
import { Fragment, useCallback, useEffect, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import useSWR from "swr";
import { Button } from "../Button";
import { LoadingDots } from "../LoadingDots";
import { EncryptJWT, jwtDecrypt } from "jose";

const InfoIcon = ({ value }: any) => {
  let iconURL;

  if (value == "render") {
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
    <img
      alt="Icon"
      src={`https://cdn-botway.deno.dev/icons/${iconURL}`}
      width={20}
    />
  );
};

const capitalizeFirstLetter = (text: String) => {
  return text.charAt(0).toUpperCase() + text.slice(1);
};

export const ProjectMain = ({ project, mutate, user }: any) => {
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
              className="rounded-2xl bg-secondary outline-none p-3"
            >
              <Content
                nav={nav}
                project={project}
                mutate={mutate}
                user={user}
              />
            </Tab.Panel>
          ))}
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
};

const Content = ({ nav, project, mutate, user }: any) => {
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
      {
        title: "Builder",
        value: "docker",
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
                      {capitalizeFirstLetter(e.title)}
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
    const [isLoading, setIsLoading] = useState(false);

    const botTokenRef: any = useRef();
    const botAppTokenRef: any = useRef();
    const botSecretTokenRef: any = useRef();
    const renderServiceIdRef: any = useRef();

    let [isOpen, setIsOpen] = useState(false);
    let [pluginx, setPluginx] = useState("PostgreSQL");

    const closeModal = () => {
      setIsOpen(false);
    };

    const openModal = (plugin: any) => {
      setPluginx(plugin);
      setIsOpen(true);
    };

    const plugins = [
      {
        name: "PostgreSQL",
      },
      {
        name: "Redis",
      },
    ];

    if (project.hostService == "railway") {
      plugins.push(
        {
          name: "MongoDB",
        },
        {
          name: "MySQL",
        }
      );
    }

    const TokensOnSubmit = useCallback(
      async (e: any) => {
        e.preventDefault();

        try {
          setIsLoading(true);

          const formData = new FormData();

          const botToken = await new EncryptJWT({
            data: botTokenRef.current.value,
          })
            .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
            .encrypt(BW_SECRET_KEY);

          if (project.platform != "telegram") {
            const botAppToken = await new EncryptJWT({
              data: botAppTokenRef.current.value,
            })
              .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
              .encrypt(BW_SECRET_KEY);

            formData.append("botAppToken", botAppToken);
          }

          if (project.platform == "slack" || project.platform == "twitch") {
            const botSecretToken = await new EncryptJWT({
              data: botSecretTokenRef.current.value,
            })
              .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
              .encrypt(BW_SECRET_KEY);

            formData.append("botSecretToken", botSecretToken);
          }

          if (project.hostService == "render") {
            const renderServiceId = await new EncryptJWT({
              data: renderServiceIdRef.current.value,
            })
              .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
              .encrypt(BW_SECRET_KEY);

            formData.append("renderServiceId", renderServiceId);
          }

          formData.append("id", project.id);
          formData.append("name", project.name);
          formData.append("userId", user._id);
          formData.append("repo", project.repo);
          formData.append("visibility", project.visibility);
          formData.append("platform", project.platform);
          formData.append("lang", project.lang);
          formData.append("packageManager", project.packageManager);
          formData.append("hostService", project.hostService);
          formData.append("botToken", botToken);
          formData.append("ghToken", user.githubApiToken);
          formData.append("railwayApiToken", user.railwayApiToken);
          formData.append("railwayProjectId", project.railwayProjectId);
          formData.append("railwayServiceId", project.railwayServiceId);
          formData.append("renderApiToken", user.renderApiToken);
          formData.append("repoBranch", project.repoBranch);
          formData.append(
            "pullRequestPreviewsEnabled",
            project.pullRequestPreviewsEnabled
          );

          await fetcher("/api/projects/update", {
            method: "PATCH",
            body: formData,
          });

          toast.success("Your project config has been updated", toastStyle);
        } catch (e: any) {
          toast.error(e.message, toastStyle);
        } finally {
          setIsLoading(false);
        }
      },
      [mutate]
    );

    useEffect(() => {
      const set = async () => {
        const { payload: botToken } = await jwtDecrypt(
          project.botToken,
          BW_SECRET_KEY
        );

        let botAppToken: any = "",
          botSecretToken: any = "",
          renderServiceId: any = "";

        if (project.botAppToken) {
          const { payload } = await jwtDecrypt(
            project.botAppToken,
            BW_SECRET_KEY
          );

          botAppToken = payload.data;
        }

        if (project.botSecretToken) {
          const { payload } = await jwtDecrypt(
            project.botSecretToken,
            BW_SECRET_KEY
          );

          botSecretToken = payload.data;
        }

        if (project.renderServiceId) {
          const { payload } = await jwtDecrypt(
            project.renderServiceId,
            BW_SECRET_KEY
          );

          renderServiceId = payload.data;
        }

        botTokenRef.current.value = botToken.data;
        botAppTokenRef.current.value = botAppToken;
        botSecretTokenRef.current.value = botSecretToken;
        renderServiceIdRef.current.value = renderServiceId;
      };

      set().catch(console.error);
    }, [project]);

    return (
      <div className="overflow-hidden sm:rounded-lg">
        <div>
          <div className="px-4 py-5 sm:px-6">
            <h3 className="text-lg font-medium leading-6 text-gray-400">
              Bot Configuration
            </h3>
          </div>
          <form onSubmit={TokensOnSubmit}>
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

              {project.hostService == "render" ? (
                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="render-service-id"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Render Service ID
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={renderServiceIdRef}
                      type="password"
                      required
                    />
                  </div>
                </div>
              ) : (
                <input ref={renderServiceIdRef} hidden />
              )}
            </div>

            <div className="mb-2 space-y-2 flex justify-center border-b border-gray-800">
              <Button
                type="success"
                htmlType="submit"
                loading={isLoading}
                className="button w-full p-2 mb-6"
              >
                Update Configuration
              </Button>
            </div>
          </form>
        </div>
        <div className="overflow-hidden sm:rounded-lg">
          <div className="px-4 py-5 sm:px-6">
            <h3 className="text-lg font-medium leading-6 text-gray-400">
              Bot Plugins (Databases)
            </h3>
          </div>

          <div className="px-4 py-5 sm:px-6">
            <p>
              {project.hostService == "railway"
                ? `Railway has a are built in Database Management Interface, this
                allows you to perform common actions on your Database such as
                viewing and editing the contents of your database services in
                Railway. The interface is available for all database services
                deployed within a project.`
                : `Stateless services are simple and scalable, but most complex services eventually end up needing persistent state. Fortunately, you can store state in a fully managed PostgreSQL or Redis instance on Render.`}
            </p>
            <div className="mt-10 grid lg:grid-cols-4 sm:grid-cols-2 lt-md:!grid-cols-1 gap-4">
              {plugins.map((plugin) => (
                <div
                  onClick={() => openModal(plugin.name)}
                  className="flex items-center justify-between gap-4 px-5 py-0 rounded-lg border-2 border-dashed border-gray-800 hover:bg-secondary transition-colors duration-200"
                >
                  <div className="block flex-1 py-5 cursor-pointer">
                    <p className="flex items-center gap-1.5 text-sm text-gray-500">
                      <h2>
                        <strong className="text-base text-white leading-tight font-medium align-middle">
                          {plugin.name}
                        </strong>
                      </h2>
                      <img
                        src={`https://cdn-botway.deno.dev/icons/${plugin.name.toLowerCase()}.svg`}
                        alt={`${plugin.name.toLowerCase()} icon`}
                        className="h-6 w-h-6 max-h-6 max-w-h-6 ml-1 mt-1"
                      />
                    </p>
                  </div>
                </div>
              ))}

              <Transition appear show={isOpen} as={Fragment}>
                <Dialog as="div" className="relative z-10" onClose={closeModal}>
                  <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                  >
                    <div className="fixed inset-0 bg-black bg-opacity-25" />
                  </Transition.Child>

                  <div className="fixed inset-0 overflow-y-auto">
                    <div className="flex min-h-full items-center justify-center p-4 text-center">
                      <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0 scale-95"
                        enterTo="opacity-100 scale-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100 scale-100"
                        leaveTo="opacity-0 scale-95"
                      >
                        <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg p-6 text-left align-middle shadow-xl transition-all border border-gray-800">
                          <Dialog.Title
                            as="h3"
                            className="text-lg font-medium leading-6 text-gray-400"
                          >
                            How to Create {pluginx} Database Plugin
                          </Dialog.Title>
                          <div className="mt-2">
                            {project.hostService == "railway" ? (
                              <p className="text-sm text-gray-500">
                                1. Press <a className="font-mono">New</a> button
                                then choose database choice
                                <br />
                                <img
                                  src={`https://cdn-botway.deno.dev/screenshots/db/railway/db.svg`}
                                  alt={`db icon`}
                                />
                                <br />
                                2. Choose {pluginx} choice
                                <br />
                                <img
                                  src={`https://cdn-botway.deno.dev/screenshots/db/railway/${pluginx.toLowerCase()}.svg`}
                                  alt={`${pluginx.toLowerCase()} icon`}
                                />
                              </p>
                            ) : (
                              <p className="text-sm text-gray-500">
                                1. Press <a className="font-mono">New</a> button
                                then choose {pluginx} choice
                                <br />
                                <img
                                  src={`https://cdn-botway.deno.dev/screenshots/db/render/${pluginx.toLowerCase()}.svg`}
                                  alt={`${pluginx} icon`}
                                />
                              </p>
                            )}
                          </div>

                          <div className="mt-4">
                            <Button onClick={closeModal}>Got it</Button>
                          </div>
                        </Dialog.Panel>
                      </Transition.Child>
                    </div>
                  </div>
                </Dialog>
              </Transition>
            </div>
          </div>
        </div>
      </div>
    );
  } else if (nav == "Deployments") {
    const NoDeploys = () => {
      return (
        <div className="rounded-2xl overflow-hidden p-5 cursor-pointer border-2 border-dashed border-gray-800 hover:border-gray-600 shadow-lg transition duration-300 ease-in-out w-full h-60 flex flex-col items-center justify-center gap-4">
          <h2 className="text-md text-gray-400 text-center">
            Your project has no deploys
          </h2>
        </div>
      );
    };

    if (project.hostService == "railway" || project.renderServiceId) {
      const formData = new FormData();

      formData.append("hostService", project.hostService);
      formData.append("railwayApiToken", user.railwayApiToken);
      formData.append("railwayProjectId", project.railwayProjectId);
      formData.append("renderApiToken", user.renderApiToken);
      formData.append("renderServiceId", project.renderServiceId);

      const deploymentsFetcher = (url: any) =>
        fetch(url, {
          method: "PATCH",
          body: formData,
        }).then((res) => res.json());

      const { data, error } = useSWR(
        `/api/projects/deployments/${project.hostService}`,
        deploymentsFetcher,
        {
          refreshWhenOffline: false,
          refreshWhenHidden: false,
          refreshInterval: 0,
        }
      );

      if (!data && !error) return <LoadingDots className="mb-3" />;

      const status = (deployStatus: any) => {
        switch (deployStatus) {
          case "FAILED":
          case "update_failed":
            return "text-red-700";

          case "SUCCESS":
          case "live":
            return "text-green-700";
        }

        return "text-gray-400";
      };

      const logsURL = (hostService: any, deploy: any) => {
        if (hostService == "railway") {
          return `https://railway.app/project/${project.railwayProjectId}/service/${project.railwayServiceId}?id=${deploy.node.id}`;
        } else if (hostService == "render") {
          return `https://dashboard.render.com/web/${project.renderServiceId}/logs`;
        }

        return "";
      };

      return (
        <div className="overflow-hidden sm:rounded-lg">
          <div>
            <div className="px-4 py-5 sm:px-6">
              <h3 className="text-lg font-medium leading-6 text-gray-400">
                Bot Deployments
              </h3>
              <div className="py-11">
                {data.length != 0 ? (
                  data.map((deploy: any) => (
                    <div className="rounded-2xl border border-gray-800 overflow-hidden p-5 bg-ultralight min-h-72 mb-6">
                      <header className="flex gap-3 justify-between mb-4">
                        <hgroup>
                          <h2 className="font-medium text-lg !leading-none text-black">
                            {project.hostService == "railway" ? (
                              deploy.node.url ? (
                                <Link href={deploy.node.url} target="_blank">
                                  {deploy.node.url}
                                </Link>
                              ) : (
                                <span className={status(deploy.node.status)}>
                                  {deploy.node.status}
                                </span>
                              )
                            ) : (
                              <span className={status(deploy.deploy.status)}>
                                {deploy.deploy.status == "update_failed"
                                  ? "FAILED"
                                  : deploy.deploy.status.toUpperCase()}
                              </span>
                            )}
                          </h2>
                          <h3 className="text-gray-500 mt-1 !leading-tight">
                            {project.hostService == "railway"
                              ? deploy.node.status == "SUCCESS"
                                ? "The deployment that is live on your production domains."
                                : deploy.node.status == "FAILED"
                                ? "The deployment is failed."
                                : deploy.node.status == "REMOVED"
                                ? "The deployment is removed."
                                : "Waiting..."
                              : deploy.deploy.status == "live"
                              ? "The deployment that is live on your production domains."
                              : deploy.deploy.status == "update_failed"
                              ? "The deployment is failed."
                              : "The deployment is deactivated."}
                          </h3>
                        </hgroup>
                        <Link
                          target="_blank"
                          className="h-8 px-3.5 rounded-md inline-flex flex-shrink-0 bg-secondary whitespace-nowrap items-center gap-2 transition-colors duration-150 ease-in-out leading-none border border-gray-800 hover:border-gray-700 cursor-pointer"
                          href={logsURL(project.hostService, deploy)}
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
                          {(project.hostService == "railway" &&
                            deploy.node.status != "SUCCESS") ||
                          (project.hostService == "render" &&
                            deploy.deploy.status != "live") ? (
                            project.hostService == "railway" ||
                            deploy.deploy.status == "update_failed" ? (
                              project.hostService == "railway" &&
                              deploy.node.status == "REMOVED" ? (
                                <ArchiveIcon
                                  className="fill-red-700"
                                  size={16}
                                />
                              ) : project.hostService == "railway" &&
                                deploy.node.status == "FAILED" ? (
                                <XCircleIcon
                                  className="fill-red-700"
                                  size={16}
                                />
                              ) : project.hostService == "render" &&
                                deploy.deploy.status == "update_failed" ? (
                                <XCircleIcon
                                  className="fill-red-700"
                                  size={16}
                                />
                              ) : (
                                <ClockIcon
                                  className="fill-gray-400"
                                  size={16}
                                />
                              )
                            ) : (
                              <ArchiveIcon className="fill-red-700" size={16} />
                            )
                          ) : (
                            <CheckCircleIcon
                              className="fill-green-700"
                              size={16}
                            />
                          )}
                        </span>
                        <span className="flex items-center gap-1">
                          <img
                            src="https://cdn-botway.deno.dev/icons/docker.svg"
                            width={18}
                            className="mr-1"
                          />
                        </span>
                        {project.hostService == "railway" ? (
                          <span className="flex items-center gap-1">
                            <FileDirectoryIcon
                              size={16}
                              className="fill-gray-600 mr-1 font-mono"
                            />
                            {deploy.node.meta.rootDirectory}
                          </span>
                        ) : (
                          <></>
                        )}
                        <span className="flex items-center gap-1">
                          <GitMergeIcon
                            size={16}
                            className="fill-gray-600 mr-1"
                          />
                          {project.hostService == "railway"
                            ? deploy.node.meta.branch
                            : project.repoBranch}
                        </span>
                        <span className="inline-flex items-center gap-2 max-w-100">
                          {project.hostService == "railway" ? (
                            <Link
                              className="text-gray-400 text-sm hover:text-gray-900 hover:underline truncate"
                              href={`https://github.com/${deploy.node.meta.repo}/commit/${deploy.node.meta.commitHash}`}
                              target="_blank"
                              title={deploy.node.meta.commitMessage}
                            >
                              {deploy.node.meta.commitMessage}
                            </Link>
                          ) : (
                            <Link
                              className="text-gray-400 text-sm hover:text-gray-900 hover:underline truncate"
                              href={`https://github.com/${project.repo}/commit/${deploy.deploy.commit.id}`}
                              target="_blank"
                              title={deploy.deploy.commit.message}
                            >
                              {deploy.deploy.commit.message}
                            </Link>
                          )}
                        </span>
                      </div>
                    </div>
                  ))
                ) : (
                  <NoDeploys />
                )}
              </div>
            </div>
          </div>
        </div>
      );
    } else {
      return (
        <div className="overflow-hidden sm:rounded-lg">
          <div>
            <div className="px-4 py-5 sm:px-6">
              <h3 className="text-lg font-medium leading-6 text-gray-400">
                Bot Deployments
              </h3>
              <div className="py-11">
                <NoDeploys />
              </div>
            </div>
          </div>
        </div>
      );
    }
  } else if (nav == "Settings") {
    const [isLoadingUpdate, setIsLoadingUpdate] = useState(false);
    const [isLoadingDelete, setIsLoadingDelete] = useState(false);

    const nameRef: any = useRef();
    const iconRef: any = useRef();
    const repoRef: any = useRef();
    const buildCommandRef: any = useRef();
    const startCommandRef: any = useRef();
    const rootDirectoryRef: any = useRef();
    const repoBranchRef: any = useRef();
    const pullRequestPreviewsEnabledRef: any = useRef();
    const projectNameRef: any = useRef();
    const deleteMyBotRef: any = useRef();

    let [isOpen, setIsOpen] = useState(false);

    const closeModal = () => {
      setIsOpen(false);
    };

    const openModal = () => {
      setIsOpen(true);
    };

    const SettingsOnSubmit = useCallback(
      async (e: any) => {
        e.preventDefault();

        try {
          setIsLoadingUpdate(true);

          CheckAPITokens(user, project.hostService);

          const formData = new FormData();

          formData.append("id", project.id);
          formData.append("userId", user._id);
          formData.append("visibility", project.visibility);
          formData.append("platform", project.platform);
          formData.append("lang", project.lang);
          formData.append("packageManager", project.packageManager);
          formData.append("hostService", project.hostService);
          formData.append("botToken", project.botToken);
          formData.append("railwayApiToken", user.railwayApiToken);
          formData.append("railwayProjectId", project.railwayProjectId);
          formData.append("railwayEnvId", project.railwayEnvId);
          formData.append("railwayServiceId", project.railwayServiceId);
          formData.append("renderApiToken", user.renderApiToken);
          formData.append("renderServiceId", project.renderServiceId);
          formData.append("name", nameRef.current.value);
          formData.append("repo", repoRef.current.value);
          formData.append("icon", iconRef.current.value);
          formData.append("buildCommand", buildCommandRef.current.value);
          formData.append("startCommand", startCommandRef.current.value);
          formData.append("rootDirectory", rootDirectoryRef.current.value);
          formData.append("repoBranch", repoBranchRef.current.value);
          formData.append(
            "pullRequestPreviewsEnabled",
            pullRequestPreviewsEnabledRef.current.value
          );

          if (project.platform != "telegram") {
            formData.append("botAppToken", project.botAppToken);
          }

          if (project.platform == "slack" || project.platform == "twitch") {
            formData.append("botSecretToken", project.botSecretToken);
          }

          await fetcher("/api/projects/settings", {
            method: "PATCH",
            body: formData,
          });

          mutate();

          toast.success("Your project settings has been updated", toastStyle);
        } catch (e: any) {
          toast.error(e.message, toastStyle);
        } finally {
          setIsLoadingUpdate(false);
        }
      },
      [mutate]
    );

    const DeleteProject = useCallback(
      async (e: any) => {
        e.preventDefault();

        if (
          projectNameRef.current.value == project.name &&
          deleteMyBotRef.current.value == "delete my bot"
        ) {
          try {
            setIsLoadingDelete(true);

            CheckAPITokens(user, project.hostService);

            const formData = new FormData();

            formData.append("id", project.id);
            formData.append("userId", user._id);
            formData.append("name", project.name);
            formData.append("hostService", project.hostService);
            formData.append("railwayApiToken", user.railwayApiToken);
            formData.append("railwayProjectId", project.railwayProjectId);
            formData.append("renderServiceId", project.renderServiceId);
            formData.append("renderApiToken", user.renderApiToken);

            await fetcher("/api/projects/delete", {
              method: "PATCH",
              body: formData,
            });

            toast.success("Your project has been deleted", toastStyle);

            mutate();

            closeModal();
          } catch (e: any) {
            toast.error(e.message, toastStyle);
          } finally {
            setIsLoadingDelete(false);
          }
        }
      },
      [mutate]
    );

    useEffect(() => {
      if (project.hostService == "railway" || project.renderServiceId) {
        nameRef.current.value = project.name;
        iconRef.current.value = project.icon;
        repoRef.current.value = project.repo;
        buildCommandRef.current.value = project.buildCommand;
        startCommandRef.current.value = project.startCommand;
        rootDirectoryRef.current.value = project.rootDirectory;
        repoBranchRef.current.value = project.repoBranch;
        pullRequestPreviewsEnabledRef.current.value =
          project.pullRequestPreviewsEnabled;
      }

      // projectNameRef.current.value = "";
      // deleteMyBotRef.current.value = "";
    }, [project]);

    const SettingsForm = () => {
      return (
        <form onSubmit={SettingsOnSubmit}>
          <div className="grid lg:grid-cols-2 sm:grid-cols-1 lt-md:!grid-cols-1 gap-3">
            <div className="px-4 py-5 sm:px-6">
              <label
                htmlFor="bot-name"
                className="block text-gray-500 text-sm font-semibold"
              >
                Bot Name
              </label>
              <div className="pt-2">
                <input
                  className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                  ref={nameRef}
                  type="text"
                />
              </div>
            </div>

            {project.hostService == "railway" ? (
              <>
                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="bot-icon"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Bot Icon
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={iconRef}
                      type="text"
                    />
                  </div>
                </div>

                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="github-repo"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    GitHub Repo
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={repoRef}
                      placeholder={`user/repoName`}
                      type="text"
                    />
                  </div>
                </div>

                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="build-command"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Build Command
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={buildCommandRef}
                      placeholder="default"
                      type="text"
                    />
                  </div>
                </div>

                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="start-command"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Start Command
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={startCommandRef}
                      placeholder="default"
                      type="text"
                    />
                  </div>
                </div>

                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="root-directory"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Root Directory
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={rootDirectoryRef}
                      placeholder="./"
                      type="text"
                    />
                  </div>
                </div>

                <input ref={repoBranchRef} hidden />
                <input ref={pullRequestPreviewsEnabledRef} hidden />
              </>
            ) : project.renderServiceId ? (
              <>
                <input ref={iconRef} hidden />
                <input ref={repoRef} hidden />
                <input ref={buildCommandRef} hidden />
                <input ref={startCommandRef} hidden />
                <input ref={rootDirectoryRef} hidden />

                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="repo-branch"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Repo Branch
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={repoBranchRef}
                      type="text"
                    />
                  </div>
                </div>

                <div className="px-4 py-5 sm:px-6">
                  <label
                    htmlFor="pull-request-previews-enabled"
                    className="block text-gray-500 text-sm font-semibold"
                  >
                    Pull Request Previews Enabled
                  </label>
                  <div className="pt-2">
                    <input
                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                      ref={pullRequestPreviewsEnabledRef}
                      placeholder="yes | no"
                      type="text"
                    />
                  </div>
                </div>
              </>
            ) : (
              <>
                <input ref={iconRef} hidden />
                <input ref={repoRef} hidden />
                <input ref={buildCommandRef} hidden />
                <input ref={startCommandRef} hidden />
                <input ref={rootDirectoryRef} hidden />
                <input ref={repoBranchRef} hidden />
                <input ref={pullRequestPreviewsEnabledRef} hidden />
              </>
            )}
          </div>

          <div className="mb-2 space-y-2 flex justify-center">
            <Button
              type="success"
              htmlType="submit"
              loading={isLoadingUpdate}
              className="p-2"
            >
              Update Settings
            </Button>
          </div>

          <div className="px-4 py-5 mb-2 sm:px-6">
            <label
              htmlFor="danger-zone"
              className="block text-red-500 text-sm font-semibold"
            >
              <AlertIcon size={16} className="mr-1" /> DANGER ZONE
            </label>
            <div className="rounded-2xl overflow-hidden p-5 bg-ultralight mt-5 border border-gray-800 bg-bwdefualt">
              <header className="flex gap-3 justify-between my-2">
                <hgroup>
                  <h2 className="font-medium text-lg !leading-none text-black">
                    Delete Project
                  </h2>
                  <br />
                  <h3 className="text-gray-500 mt-1 !leading-tight">
                    Delete {project.name} and delete it on {project.hostService}
                    . This action is not reversible, so continue with extreme
                    caution.
                  </h3>
                </hgroup>
                <div></div>
              </header>

              <Button type="delete" onClick={openModal} className="button p-2">
                Delete Project
              </Button>

              <Transition appear show={isOpen} as={Fragment}>
                <Dialog as="div" className="relative z-10" onClose={closeModal}>
                  <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                  >
                    <div className="fixed inset-0 bg-black bg-opacity-25" />
                  </Transition.Child>

                  <div className="fixed inset-0 overflow-y-auto">
                    <div className="flex min-h-full items-center justify-center p-4 text-center">
                      <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0 scale-95"
                        enterTo="opacity-100 scale-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100 scale-100"
                        leaveTo="opacity-0 scale-95"
                      >
                        <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg p-6 text-left align-middle shadow-xl transition-all border border-gray-800">
                          <Dialog.Title
                            as="h3"
                            className="text-lg font-medium leading-6 text-white"
                          >
                            Delete Bot Project
                          </Dialog.Title>

                          <div className="mt-2 text-gray-400">
                            Delete{" "}
                            <a className="text-gray-500">{project.name}</a> and
                            delete it on{" "}
                            <a className="text-gray-500">
                              {project.hostService}
                            </a>
                            . This action is not reversible, so continue with
                            extreme caution.
                          </div>

                          <form onSubmit={DeleteProject}>
                            {!isLoadingDelete ? (
                              <>
                                <div className="mt-4">
                                  <label
                                    htmlFor="project-name"
                                    className="block text-gray-500 text-sm font-semibold"
                                  >
                                    Enter the project name{" "}
                                    <a className="text-white font-bold">
                                      {project.name}
                                    </a>{" "}
                                    to continue:
                                  </label>
                                  <div className="pt-2">
                                    <input
                                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                                      ref={projectNameRef}
                                      type="text"
                                      required
                                    />
                                  </div>
                                </div>

                                <div className="mt-4">
                                  <label
                                    htmlFor="project-name"
                                    className="block text-gray-500 text-sm font-semibold"
                                  >
                                    To verify, type{" "}
                                    <a className="text-white font-bold">
                                      delete my bot
                                    </a>{" "}
                                    below:
                                  </label>
                                  <div className="pt-2">
                                    <input
                                      className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                                      ref={deleteMyBotRef}
                                      type="text"
                                      required
                                    />
                                  </div>
                                </div>
                              </>
                            ) : (
                              <></>
                            )}

                            <div className="mt-4">
                              <Button
                                type="delete"
                                loading={isLoadingDelete}
                                onClick={DeleteProject}
                                className="button p-2"
                              >
                                Continue
                              </Button>
                            </div>
                          </form>
                        </Dialog.Panel>
                      </Transition.Child>
                    </div>
                  </div>
                </Dialog>
              </Transition>
            </div>
          </div>
        </form>
      );
    };

    return (
      <div className="overflow-hidden sm:rounded-lg">
        <div>
          <div className="px-4 py-5 sm:px-6">
            <h3 className="text-lg font-medium leading-6 text-gray-400">
              Bot Settings
            </h3>
          </div>
        </div>
        {project.hostService == "render" ? (
          project.renderServiceId ? (
            <SettingsForm />
          ) : (
            <div className="py-11">
              <div className="rounded-2xl overflow-hidden p-5 cursor-pointer border-2 border-dashed border-gray-800 hover:border-gray-600 shadow-lg transition duration-300 ease-in-out w-full h-60 flex flex-col items-center justify-center gap-4">
                <h2 className="text-md text-yellow-500 text-center">
                  You need to set Render Service Id for this project
                </h2>
              </div>
            </div>
          )
        ) : (
          <SettingsForm />
        )}
      </div>
    );
  }

  return <></>;
};
