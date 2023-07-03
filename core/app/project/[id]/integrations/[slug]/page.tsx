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
import {
  ArrowUpRightIcon,
  FileDirectoryIcon,
  MarkGithubIcon,
} from "@primer/octicons-react";
import { marked } from "marked";
import { fetcher } from "@/tools/fetch";
import { Fragment, useState } from "react";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import { Dialog, Transition } from "@headlessui/react";
import { Field, Form, Formik } from "formik";
import { Button } from "@/components/Button";

export const revalidate = 0;

const queryClient = new QueryClient();

const Project = ({ user, projectId, slug }: any) => {
  const [isLoading, setIsLoading] = useState(false);
  const [open, setOpen] = useState(false);

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

  const fetchIntegration = async () => {
    const intx = await fetcher(`/api/integrations/x?slug=${slug}`, {
      method: "GET",
    });

    return intx;
  };

  const { data: int, isLoading: integrationsIsLoading } = useQuery(
    ["integration"],
    fetchIntegration,
    {
      refetchInterval: 1,
      refetchOnReconnect: true,
      refetchOnWindowFocus: true,
      refetchIntervalInBackground: true,
    }
  );

  async function addIntegration(formData: any) {
    try {
      setIsLoading(true);

      let vars = {};

      if (int.variables) {
        if (int.variables.length === 1) {
          vars = { v: formData.v1, k: int.variables[0].name };
        } else {
          vars = {
            v1: formData.v1,
            k1: int.variables[0].name,
            v2: formData.v2,
            k2: int.variables[1].name,
          };
        }
      }

      const body = {
        name: int.name,
        slug: int.slug,
        template_repo: int.template_repo,
        is_plugin: int.is_plugin,
        projectId: project?.railway_project_id,
        vars,
        def_vars: int.def_variables,
        plugin: int.plugin,
        has_volume: int.has_volume,
        volume_path: int.volume_path,
      };

      const newInt = await fetcher("/api/integrations/add", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      if (newInt.message === "Success") {
        toast.success(
          "You have successfully created a new bot integration",
          toastStyle
        );

        setOpen(false);
      } else {
        toast.error(newInt.error, toastStyle);

        setOpen(false);
      }
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);
    }
  }

  const check = async (int: any) => {
    if (!int.is_plugin && int.variables.length != 0) {
      setOpen(true);
    } else if (!int.is_plugin && int.variables.length === 0) {
      try {
        setIsLoading(true);

        const body = {
          name: int.name,
          slug: int.slug,
          template_repo: int.template_repo,
          is_plugin: int.is_plugin,
          projectId: project?.railway_project_id,
          def_vars: int.def_variables,
          plugin: int.plugin,
          has_volume: int.has_volume,
          volume_path: int.volume_path,
        };

        const newInt = await fetcher("/api/integrations/add", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(body),
        });

        if (newInt.message === "Success") {
          toast.success(
            "You have successfully created a new bot integration",
            toastStyle
          );
        } else {
          toast.error(newInt.error, toastStyle);
        }
      } catch (e: any) {
        toast.error(e.message, toastStyle);
      } finally {
        setIsLoading(false);
      }
    } else {
      try {
        setIsLoading(true);

        const newInt = await fetcher("/api/integrations/add", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            name: int.title,
            slug: int.slug,
            template_repo: int.template_repo,
            is_plugin: int.is_plugin,
            projectId: project?.railway_project_id,
          }),
        });

        if (newInt.message === "Success") {
          toast.success(
            "You have successfully created a new bot integration",
            toastStyle
          );

          setOpen(false);
        } else {
          toast.error(newInt.error, toastStyle);

          setOpen(false);
        }
      } catch (e: any) {
        toast.error(e.message, toastStyle);
      } finally {
        setIsLoading(false);
      }
    }
  };

  return (
    <>
      {projectIsLoading || integrationsIsLoading ? (
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : (
        <ProjectLayout
          user={user}
          projectId={projectId}
          projectName={project?.name}
          projectRWID={project?.railway_project_id}
          grid={true}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">{int?.name}</h1>
          </div>
          <div className="mx-6">
            <div className="w-full mx-auto mb-32 px-6 lg:px-0 flex flex-col">
              <div className="relative top-0 mt-4 mb-16 flex flex-col lg:flex-row lg:space-x-12">
                <div className="w-full lg:w-9/12 flex flex-col">
                  <div className="flex flex-col items-center space-y-4 lg:space-y-0 lg:flex-row lg:space-x-8">
                    <img
                      alt={int?.name}
                      src={`https://cdn-botway.deno.dev/icons/${int?.slug}.svg`}
                      className="h-16 w-16 rounded-sm"
                    />

                    <div className="text-center lg:text-left space-y-4 lg:space-y-3 flex flex-col">
                      <p className="text-white text-lg font-semibold">
                        {int?.name}
                      </p>
                      <p className="text-base text-gray-400">{int?.desc}</p>
                    </div>

                    <a
                      className={`flex md:!hidden cursor-pointer items-center justify-center border transition-all duration-200 active:scale-95 outline-none focus:outline-none text-white hover:opacity-90 h-[42px] py-2 px-3 rounded-lg text-base leading-6 space-x-3 ${
                        !int?.soon ? "cursor-pointer" : "cursor-not-allowed"
                      }`}
                    >
                      <Button
                        htmlType="submit"
                        type="success"
                        onClick={!int?.soon ? () => check(int) : null}
                        loading={isLoading}
                        disabled={int?.soon}
                        className={int?.soon ? "cursor-not-allowed" : ""}
                      >
                        {int?.soon ? "Soon" : `Add ${int?.name}`}
                      </Button>
                    </a>
                  </div>

                  <div className="pt-16" />

                  <div className="relative min-h-[420px] lg:min-h-[540px] py-16 w-full rounded-3xl border border-gray-800 bg-secondary flex flex-col justify-center items-center">
                    <div className="grid gap-4 items-center justify-items-center grid-cols-1 px-10">
                      <article className="prose prose-gray prose-headings:text-white prose-p:text-gray-400 prose-a:text-blue-700 prose-strong:text-white prose-ol:text-white prose-li:text-white prose-ul:text-white prose-pre:bg-bwdefualt prose-pre:border prose-pre:rounded-2xl prose-pre:border-gray-800 prose-blockquote:border-l-4 prose-blockquote:border-gray-800 prose-hr:border prose-hr:border-gray-800 prose-code:bg-bwdefualt prose-code:font-mono prose-code:rounded prose-code:text-gray-100">
                        <div
                          dangerouslySetInnerHTML={{
                            __html: marked.parse(int?.overview),
                          }}
                        />
                      </article>
                    </div>
                  </div>
                </div>

                <div className="w-full lg:w-3/12 lg:mt-6 lg:sticky lg:top-[48px] align-self[flex-start] flex flex-col">
                  <a
                    className={`hidden md:flex items-center justify-center transition-all duration-200 active:scale-95 outline-none focus:outline-none lg:!flex text-white hover:opacity-90 h-[42px] py-2 px-3 rounded-lg text-base leading-6 space-x-3 ${
                      !int?.soon ? "cursor-pointer" : "cursor-not-allowed"
                    }`}
                  >
                    <Button
                      htmlType="submit"
                      type="success"
                      onClick={!int?.soon ? () => check(int) : null}
                      loading={isLoading}
                      disabled={int?.soon}
                      className={int?.soon ? "cursor-not-allowed" : ""}
                    >
                      {int?.soon ? "Soon" : `Add ${int?.name}`}
                    </Button>
                  </a>

                  <div className="mt-16 flex flex-col space-y-6">
                    <div className="flex flex-col space-y-4">
                      <div className="ml-2 text-gray-400 flex space-x-6 items-center">
                        <ArrowUpRightIcon size={20} />

                        <a
                          href={int?.website}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-sm"
                        >
                          Website
                        </a>
                      </div>

                      <div className="ml-2 text-gray-400 flex space-x-6 items-center">
                        <svg
                          width="18"
                          height="15"
                          viewBox="0 0 18 15"
                          fill="none"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                            d="M5.66065 14.6275C12.4531 14.6275 16.1683 8.99997 16.1683 4.11986C16.1683 3.95997 16.1683 3.80086 16.1575 3.64174C16.8802 3.11919 17.5041 2.47216 18 1.73093C17.326 2.02959 16.611 2.22546 15.8789 2.31199C16.6498 1.85047 17.2268 1.12458 17.5025 0.269412C16.7775 0.699562 15.9844 1.0027 15.1574 1.16576C14.6006 0.573704 13.8642 0.181668 13.0621 0.0503113C12.2601 -0.0810453 11.4371 0.0556009 10.7205 0.439105C10.0039 0.822609 9.43369 1.43159 9.09808 2.17181C8.76248 2.91204 8.68019 3.74223 8.86395 4.53394C7.39568 4.46032 5.95931 4.07876 4.64808 3.41401C3.33685 2.74927 2.18007 1.81621 1.25283 0.675396C0.780574 1.48838 0.635924 2.45079 0.848329 3.36668C1.06073 4.28257 1.61423 5.08308 2.39611 5.60522C1.80842 5.58779 1.23355 5.42925 0.72 5.14298V5.18974C0.720224 6.04239 1.01538 6.86871 1.5554 7.52854C2.09542 8.18837 2.84706 8.6411 3.68283 8.80992C3.13913 8.9582 2.56868 8.97987 2.0153 8.87327C2.25128 9.60708 2.71071 10.2488 3.32934 10.7086C3.94797 11.1685 4.69485 11.4234 5.46553 11.4379C4.15782 12.4657 2.54236 13.0236 0.879117 13.0219C0.585288 13.0213 0.291746 13.0036 0 12.9686C1.68886 14.0524 3.65394 14.6273 5.66065 14.6246"
                            fill="#9CA3AF"
                          />
                        </svg>

                        <a
                          href={`https://twitter.com/${int?.twitter}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-sm text-[#1DA1F2]"
                        >
                          {int?.twitter}
                        </a>
                      </div>

                      <div className="ml-2 text-gray-400 flex space-x-6 items-center">
                        <MarkGithubIcon size={18} />

                        <a
                          href={int?.repo}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-sm"
                        >
                          Repo
                        </a>
                      </div>

                      <div className="ml-2 text-gray-400 flex space-x-6 items-center">
                        <FileDirectoryIcon size={18} />

                        <p className="text-sm">{int?.category}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <Transition.Root show={open} as={Fragment}>
            <Dialog as="div" className="relative z-10" onClose={setOpen}>
              <Transition.Child
                as={Fragment}
                enter="ease-in-out duration-500"
                enterFrom="opacity-0"
                enterTo="opacity-100"
                leave="ease-in-out duration-500"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
              >
                <div className="fixed inset-0 bg-bwdefualt bg-opacity-50 transition-opacity" />
              </Transition.Child>

              <div className="fixed inset-0 overflow-hidden">
                <div className="absolute inset-0 overflow-hidden">
                  <div className="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10">
                    <Transition.Child
                      as={Fragment}
                      enter="transform transition ease-in-out duration-500 sm:duration-700"
                      enterFrom="translate-x-full"
                      enterTo="translate-x-0"
                      leave="transform transition ease-in-out duration-500 sm:duration-700"
                      leaveFrom="translate-x-0"
                      leaveTo="translate-x-full"
                    >
                      <Dialog.Panel className="pointer-events-auto relative w-screen max-w-md">
                        <Transition.Child
                          as={Fragment}
                          enter="ease-in-out duration-200"
                          enterFrom="opacity-0"
                          enterTo="opacity-100"
                          leave="ease-in-out duration-500"
                          leaveFrom="opacity-100"
                          leaveTo="opacity-0"
                        >
                          <div className="absolute left-0 top-0 -ml-8 flex pr-2 pt-4 sm:-ml-10 sm:pr-4"></div>
                        </Transition.Child>

                        <div className="flex h-full flex-col overflow-y-scroll bg-secondary border-l border-gray-800 py-4 shadow-xl">
                          <div className="px-4 border-b border-gray-800 sm:px-6">
                            <Dialog.Title className="text-lg font-semibold text-white leading-6 pb-4">
                              Add Integration To My Project
                            </Dialog.Title>
                          </div>

                          <div className="relative mt-4 flex-1 px-4 sm:px-6">
                            <div className="my-4 max-w-4xl space-y-8">
                              <Formik
                                initialValues={{
                                  slug: int.slug,
                                  v1: "",
                                  v2: "",
                                }}
                                onSubmit={addIntegration}
                              >
                                {() => (
                                  <>
                                    <Form className="column w-full">
                                      <div>
                                        {int.variables ? (
                                          int.variables.map((varx: any) => (
                                            <>
                                              <label className="text-white col-span-12 text-base lg:col-span-5">
                                                {varx.name}
                                              </label>

                                              <div className="pt-2" />

                                              <Field
                                                className="input"
                                                id={`v${varx.index}`}
                                                name={`v${varx.index}`}
                                                type={
                                                  varx.is_hidden
                                                    ? "password"
                                                    : "text"
                                                }
                                              />

                                              <div className="pb-2" />
                                            </>
                                          ))
                                        ) : (
                                          <></>
                                        )}
                                      </div>

                                      <br />

                                      <Button
                                        htmlType="submit"
                                        type="success"
                                        loading={isLoading}
                                      >
                                        Add
                                      </Button>
                                    </Form>
                                  </>
                                )}
                              </Formik>
                            </div>
                          </div>
                        </div>
                      </Dialog.Panel>
                    </Transition.Child>
                  </div>
                </div>
              </div>
            </Dialog>
          </Transition.Root>
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
        <Project user={user} projectId={params.id} slug={params.slug} />
      </QueryClientProvider>
    );
  }

  redirect("/");
};

export default ProjectPage;
