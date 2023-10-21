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
import { Field, Form, Formik } from "formik";
import clsx from "clsx";
import * as Yup from "yup";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import { fetcher } from "@/tools/fetch";
import { Fragment, useState } from "react";
import styles from "@/components/Button/Button.module.scss";
import {
  AlertIcon,
  CheckIcon,
  PencilIcon,
  XCircleIcon,
} from "@primer/octicons-react";
import { capitalizeFirstLetter } from "@/tools/text";
import { Dialog, Transition } from "@headlessui/react";
import { EncryptJWT } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";

export const revalidate = 0;

const UpdateNameSchema = Yup.object().shape({
  name: Yup.string().min(3),
  icon: Yup.string(),
  buildCmd: Yup.string(),
  startCmd: Yup.string(),
  rootDir: Yup.string(),
});

export const CheckTokens = (project: any) => {
  if (project?.platform === "telegram") {
    if (project?.bot_token.length != 0) {
      return true;
    }
  } else if (project?.platform === "discord") {
    if (project?.bot_token.length != 0 && project?.bot_app_token.length != 0) {
      return true;
    }
  } else if (project?.platform === "slack" || project?.platform === "twitch") {
    if (
      project?.bot_token.length != 0 &&
      project?.bot_app_token.length != 0 &&
      project?.bot_secret_token.length != 0
    ) {
      return true;
    }
  }

  return false;
};

const queryClient = new QueryClient();

const Project = ({ user, projectId }: any) => {
  const [isLoading, setIsLoading] = useState(false);
  const [isLoadingDelete, setIsLoadingDelete] = useState(false);
  const [isLoadingTokens, setIsLoadingTokens] = useState(false);
  const [isOpen, setIsOpen] = useState(false);

  const closeModal = () => {
    setIsOpen(false);
  };

  const openModal = () => {
    setIsOpen(true);
  };

  const fetchProject = async () => {
    const { data: project } = await supabase
      .from("projects")
      .select("*")
      .eq("id", projectId)
      .single();

    return project;
  };

  const { data: project, isLoading: projectIsLoading } = useQuery({
    queryKey: ["project"],
    queryFn: fetchProject,
    refetchInterval: 360,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const updateProjectName = async (formData: any) => {
    try {
      setIsLoading(true);

      const body = {
        name: formData.name,
        icon: formData.icon,
        buildCmd: formData.buildCmd,
        startCmd: formData.startCmd,
        rootDir: formData.rootDir,
        projectId,
      };

      const settings = await fetcher("/api/projects/settings", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      if (settings.message === "Success") {
        toast.success("Done", toastStyle);
      } else {
        toast.error(settings.error, toastStyle);
      }
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);
    }
  };

  const editTokens = async (formData: any) => {
    try {
      setIsLoadingTokens(true);

      const botToken = await new EncryptJWT({
        data: formData.botToken,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      let body = {
        botToken,
        botAppToken: "",
        botSecretToken: "",
        projectId,
        platform: project?.platform,
        zeaburEnvId: project?.zeabur_env_id,
        zeaburServiceId: project?.zeabur_service_id,
      };

      if (project?.platform != "telegram") {
        const botAppToken = await new EncryptJWT({
          data: formData.botAppToken,
        })
          .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
          .encrypt(BW_SECRET_KEY);

        body["botAppToken"] = botAppToken;
      }

      if (project?.platform === "slack" || project?.platform === "twitch") {
        const botSecretToken = await new EncryptJWT({
          data: formData.botSecretToken,
        })
          .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
          .encrypt(BW_SECRET_KEY);

        body["botSecretToken"] = botSecretToken;
      }

      await fetcher("/api/projects/config", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      closeModal();
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoadingTokens(false);
    }
  };

  const deleteProject = async () => {
    try {
      setIsLoadingDelete(true);

      await fetcher("/api/projects/settings/delete", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          projectId,
          zeaburEnvId: project?.zeabur_env_id,
        }),
      });
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoadingDelete(false);
    }
  };

  if (!project && !projectIsLoading) {
    redirect("/");
  }

  return (
    <>
      {projectIsLoading ? (
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : (
        <ProjectLayout
          user={user}
          projectId={projectId}
          projectName={project?.name}
          projectZBID={project?.zeabur_project_id}
        >
          <main className="flex-1 max-h-screen">
            <div className="overflow-y-auto mx-auto flex flex-col px-5 py-6 lg:px-16">
              <Formik
                initialValues={{
                  name: project?.name,
                  icon: project?.icon,
                  buildCmd: project?.build_command,
                  startCmd: project?.start_command,
                  rootDir: project?.root_directory,
                }}
                validationSchema={UpdateNameSchema}
                onSubmit={updateProjectName}
              >
                {({ errors, values }) => (
                  <Form className="column w-full">
                    <div className="mb-10">
                      <h3 className="text-white mb-2 text-xl">
                        Project Settings
                      </h3>
                    </div>

                    <div className="bg-secondary border border-gray-800 overflow-hidden rounded-2xl shadow">
                      <div className="px-8 py-8 opacity-100">
                        <div className="relative col-span-12 flex flex-col gap-6 lg:col-span-7">
                          <div className="grid lg:grid-cols-2 sm:grid-cols-2 lt-md:!grid-cols-1 gap-5">
                            <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                              <div className="flex flex-row space-x-2 justify-between col-span-12">
                                <label className="block text-gray-400 text-sm leading-4">
                                  Project name
                                </label>
                              </div>

                              <div className="col-span-12">
                                <div>
                                  <div className="relative">
                                    <Field
                                      className="input"
                                      id="name"
                                      name="name"
                                      type="text"
                                    />

                                    {errors.name ? (
                                      <div className="text-red-600 text-sm font-semibold pt-2">
                                        {errors.name.toString()}
                                      </div>
                                    ) : null}
                                  </div>
                                </div>
                              </div>
                            </div>

                            <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                              <div className="flex flex-row space-x-2 justify-between col-span-12">
                                <label className="block text-gray-400 text-sm leading-4">
                                  Project Icon
                                </label>
                              </div>

                              <div className="col-span-12">
                                <div>
                                  <div className="relative">
                                    <Field
                                      className="input"
                                      id="icon"
                                      name="icon"
                                      type="text"
                                    />

                                    {errors.icon ? (
                                      <div className="text-red-600 text-sm font-semibold pt-2">
                                        {errors.icon.toString()}
                                      </div>
                                    ) : null}
                                  </div>
                                </div>
                              </div>
                            </div>

                            <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                              <div className="flex flex-row space-x-2 justify-between col-span-12">
                                <label className="block text-gray-400 text-sm leading-4">
                                  Build Command
                                </label>
                              </div>

                              <div className="col-span-12">
                                <div>
                                  <div className="relative">
                                    <Field
                                      className="input"
                                      id="buildCmd"
                                      name="buildCmd"
                                      type="text"
                                      placeholder={
                                        project?.build_command != ""
                                          ? null
                                          : "default"
                                      }
                                    />

                                    {errors.buildCmd ? (
                                      <div className="text-red-600 text-sm font-semibold pt-2">
                                        {errors.buildCmd.toString()}
                                      </div>
                                    ) : null}
                                  </div>
                                </div>
                              </div>
                            </div>

                            <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                              <div className="flex flex-row space-x-2 justify-between col-span-12">
                                <label className="block text-gray-400 text-sm leading-4">
                                  Start Command
                                </label>
                              </div>

                              <div className="col-span-12">
                                <div>
                                  <div className="relative">
                                    <Field
                                      className="input"
                                      id="startCmd"
                                      name="startCmd"
                                      type="text"
                                      placeholder={
                                        project?.start_command != ""
                                          ? null
                                          : "default"
                                      }
                                    />

                                    {errors.startCmd ? (
                                      <div className="text-red-600 text-sm font-semibold pt-2">
                                        {errors.startCmd.toString()}
                                      </div>
                                    ) : null}
                                  </div>
                                </div>
                              </div>
                            </div>

                            <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                              <div className="flex flex-row space-x-2 justify-between col-span-12">
                                <label className="block text-gray-400 text-sm leading-4">
                                  Root Directory
                                </label>
                              </div>

                              <div className="col-span-12">
                                <div>
                                  <div className="relative">
                                    <Field
                                      className="input"
                                      id="rootDir"
                                      name="rootDir"
                                      type="text"
                                      placeholder={
                                        project?.root_directory != ""
                                          ? null
                                          : "./"
                                      }
                                    />

                                    {errors.rootDir ? (
                                      <div className="text-red-600 text-sm font-semibold pt-2">
                                        {errors.rootDir.toString()}
                                      </div>
                                    ) : null}
                                  </div>
                                </div>
                              </div>
                            </div>

                            <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                              <div className="flex flex-row space-x-2 justify-between col-span-12">
                                <label className="block text-gray-400 text-sm leading-4">
                                  Project ID
                                </label>
                              </div>

                              <div className="col-span-12">
                                <div>
                                  <div className="relative">
                                    <input
                                      id="id"
                                      name="id"
                                      disabled={true}
                                      type="text"
                                      className="input"
                                      value={project?.id}
                                    />
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>

                      <div className="border-gray-800 border-t" />

                      <div className="flex bg-bwdefualt py-4 px-8">
                        <div className="flex w-full items-center gap-2 justify-end">
                          <div className="flex items-center gap-2">
                            <button
                              className={clsx(
                                "relative text-white bg-blue-700 inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 shadow-sm text-xs px-2.5 py-1",
                                project?.name === values.name &&
                                  project?.icon === values.icon &&
                                  project?.build_command === values.buildCmd &&
                                  project?.start_command === values.startCmd &&
                                  project?.root_directory === values.rootDir
                                  ? "opacity-50 cursor-not-allowed pointer-events-none"
                                  : "cursor-pointer",
                              )}
                              type="submit"
                            >
                              {isLoading && (
                                <LoadingDots
                                  className={styles.loading}
                                  children
                                />
                              )}

                              <span className="truncate">Save</span>
                            </button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </Form>
                )}
              </Formik>

              <div>
                <div className="my-6">
                  <h3 className="text-white text-xl">Bot Configuration</h3>
                </div>

                <div className="my-4 max-w-full space-y-8">
                  <div className="overflow-x-auto flex-grow rounded-2xl border border-gray-800">
                    <table className="w-full border-collapse select-auto">
                      <thead>
                        <tr className="bg-secondary border-b border-gray-800">
                          <th className="py-3 px-4 text-left font-semibold text-xs text-gray-400">
                            Name
                          </th>
                          <th className="py-3 px-4 text-left hidden md:block font-semibold text-xs text-gray-400">
                            Platform
                          </th>
                          <th className="py-3 px-4 text-left font-semibold text-xs text-gray-400">
                            Tokens Status
                          </th>
                          <th className="py-3 px-4 text-left font-semibold text-xs text-gray-400" />
                        </tr>
                      </thead>
                      <tbody>
                        <tr>
                          <td
                            className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap"
                            style={{ minWidth: "64px", maxWidth: "400px" }}
                          >
                            <div className="flex space-x-2 items-center">
                              <p className="text-sm pt-0.5 text-white">
                                {capitalizeFirstLetter(project?.platform)}{" "}
                                Config
                              </p>
                            </div>
                          </td>
                          <td
                            className="py-3 px-4 overflow-hidden hidden md:table-cell overflow-ellipsis whitespace-nowrap"
                            style={{ minWidth: "64px", maxWidth: "400px" }}
                          >
                            <div className="flex space-x-2 mt-0.5 items-center">
                              <div className="text-gray-500" aria-hidden="true">
                                <img
                                  src={`https://cdn-botway.deno.dev/icons/${project?.platform}.svg`}
                                  width={20}
                                />
                              </div>

                              <p className="text-sm text-white">
                                {capitalizeFirstLetter(project?.platform)}
                              </p>
                            </div>
                          </td>
                          <td
                            className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500"
                            style={{ minWidth: "64px", maxWidth: "400px" }}
                          >
                            {CheckTokens(project) ? (
                              <>
                                <CheckIcon
                                  size={18}
                                  className="fill-green-600"
                                />
                              </>
                            ) : (
                              <>
                                <XCircleIcon
                                  size={18}
                                  className="fill-red-600"
                                />
                              </>
                            )}
                          </td>
                          <td
                            className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500 w-[44px]"
                            style={{ minWidth: "64px", maxWidth: "400px" }}
                          >
                            <button
                              onClick={openModal}
                              className="flex items-center justify-center transform transition-transform duration-50 active:scale-95 focus:outline-none focus-visible:ring-2 focus-visible:outline-none h-[34px] py-1.5 rounded-md text-sm leading-5 space-x-2 w-[34px] px-0"
                            >
                              <div className="!w-4 !h-4" aria-hidden="true">
                                <PencilIcon
                                  size={18}
                                  className="hover:fill-gray-400 transition-all duration-200"
                                />
                              </div>
                            </button>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>

              <section>
                <div className="my-1">
                  <h3 className="text-red-500 mb-6 text-xl">Danger Zone</h3>
                </div>

                <div className="relative">
                  <div className="transition-opacity duration-300">
                    <div className="overflow-hidden rounded-2xl border bg-red-200 border-red-700 shadow-sm mb-8">
                      <div>
                        <div className="px-6 py-4">
                          <div className="relative py-4 px-4 flex space-x-4 items-start ">
                            <div className="text-red-900">
                              <AlertIcon size={18} />
                            </div>

                            <div className="flex flex-1 items-center justify-between">
                              <div>
                                <h3 className="block text-base font-normal mb-1">
                                  Delete Project
                                </h3>
                                <div className="text-sm text-red-900">
                                  <div>
                                    <p className="mb-4 block">
                                      Delete {project?.name} and delete it on
                                      Zeabur. This action is not reversible, so
                                      continue with extreme caution.
                                    </p>

                                    <button
                                      className="relative cursor-pointer text-white bg-red-500 hover:bg-red-600 inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 shadow-sm text-xs px-2.5 py-1"
                                      onClick={deleteProject}
                                      type="submit"
                                    >
                                      {isLoadingDelete && (
                                        <LoadingDots
                                          className={styles.loading}
                                          children
                                        />
                                      )}
                                      <span className="truncate">
                                        Delete Bot Project
                                      </span>
                                    </button>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </section>
            </div>
          </main>

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
                <div className="fixed inset-0 bg-bwdefualt bg-opacity-50 transition-opacity" />
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
                    <Dialog.Panel className="w-full bg-secondary max-w-md transform overflow-hidden rounded-2xl p-6 text-left align-middle shadow-xl transition-all border border-gray-800">
                      <Dialog.Title
                        as="h3"
                        className="text-lg font-medium pb-2 leading-6 text-white"
                      >
                        Edit {capitalizeFirstLetter(project?.platform)} Config
                        Tokens
                      </Dialog.Title>
                      <div className="mt-2">
                        <Formik
                          initialValues={{
                            botToken: "",
                            botAppToken: "",
                            botSecretToken: "",
                          }}
                          onSubmit={editTokens}
                        >
                          {({ errors }) => (
                            <Form className="column w-full">
                              <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                                <div className="flex flex-row space-x-2 justify-between col-span-12">
                                  <label className="block text-gray-400 text-sm leading-4">
                                    {capitalizeFirstLetter(project?.platform)}{" "}
                                    Bot Token
                                  </label>
                                </div>

                                <div className="col-span-12">
                                  <div>
                                    <div className="relative">
                                      <Field
                                        className="input"
                                        id="botToken"
                                        name="botToken"
                                        type="password"
                                      />
                                    </div>
                                  </div>
                                </div>
                              </div>

                              {project?.platform != "telegram" ? (
                                <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                                  <br />
                                  <div className="flex flex-row space-x-2 justify-between col-span-12">
                                    <label className="block text-gray-400 text-sm leading-4">
                                      {capitalizeFirstLetter(project?.platform)}{" "}
                                      {project?.platform != "twitch"
                                        ? `Bot App ${
                                            project?.platform === "discord"
                                              ? "ID"
                                              : "Token"
                                          }`
                                        : "Bot Client ID"}
                                    </label>
                                  </div>

                                  <div className="col-span-12">
                                    <div>
                                      <div className="relative">
                                        <Field
                                          className="input"
                                          id="botAppToken"
                                          name="botAppToken"
                                          type="password"
                                        />
                                      </div>
                                    </div>
                                  </div>
                                </div>
                              ) : (
                                <></>
                              )}

                              {project?.platform === "slack" ||
                              project?.platform === "twitch" ? (
                                <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                                  <br />
                                  <div className="flex flex-row space-x-2 justify-between col-span-12">
                                    <label className="block text-gray-400 text-sm leading-4">
                                      {capitalizeFirstLetter(project?.platform)}{" "}
                                      {project?.platform === "twitch"
                                        ? "Bot Client Secret"
                                        : "Bot Signing Secret"}
                                    </label>
                                  </div>

                                  <div className="col-span-12">
                                    <div>
                                      <div className="relative">
                                        <Field
                                          className="input"
                                          id="botSecretToken"
                                          name="botSecretToken"
                                          type="password"
                                        />
                                      </div>
                                    </div>
                                  </div>
                                </div>
                              ) : (
                                <></>
                              )}

                              <br />

                              <div className="flex w-full items-center gap-2 justify-end">
                                <div className="flex items-center gap-2">
                                  <button
                                    className="relative text-white cursor-pointer bg-blue-700 inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 shadow-sm text-xs px-2.5 py-1"
                                    type="submit"
                                  >
                                    {isLoadingTokens && (
                                      <LoadingDots
                                        className={styles.loading}
                                        children
                                      />
                                    )}
                                    <span className="truncate">Save</span>
                                  </button>
                                </div>
                              </div>
                            </Form>
                          )}
                        </Formik>
                      </div>
                    </Dialog.Panel>
                  </Transition.Child>
                </div>
              </div>
            </Dialog>
          </Transition>
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
