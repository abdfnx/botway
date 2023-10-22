"use client";

import { redirect, useRouter } from "next/navigation";
import { useAuth } from "@/supabase/auth/provider";
import { LoadingDots } from "@/components/LoadingDots";
import supabase from "@/supabase/browser";
import { ProjectLayout } from "@/components/Layout/project";
import {
  useQuery,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { fetcher } from "@/tools/fetch";
import Link from "next/link";
import { EncryptJWT } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { Button } from "@/components/Button";
import { Fragment, useState } from "react";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import { Dialog, Transition } from "@headlessui/react";
import { Field, Form, Formik } from "formik";
import styles from "@/components/Button/Button.module.scss";
import slug from "slug";

export const revalidate = 0;

const queryClient = new QueryClient();

const CE = ({ user, projectId }: any) => {
  const [isLoading, setIsLoading] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const [ce, setCE] = useState("");

  const closeModal = () => {
    setIsOpen(false);
  };

  const openModal = () => {
    setIsOpen(true);
  };

  const router = useRouter();

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
    refetchInterval: 1,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const enableCE = async (formData: any) => {
    try {
      setIsLoading(true);

      let body = {};

      let ce: any;

      if (project?.enable_ce) {
        body = {
          projectId,
        };

        ce = await fetcher("/api/ce", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(body),
        });

        if (ce.message != "Success") {
          toast.error(ce.error, toastStyle);
        }

        router.push(`https://${ce.domain}`);
      } else {
        body = {
          password: await new EncryptJWT({
            data: formData.password,
          })
            .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
            .encrypt(BW_SECRET_KEY),
          slug: slug(project?.name, "-"),
          projectId,
        };

        ce = await fetcher("/api/ce/enable", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(body),
        });

        closeModal();
      }

      if (ce.message != "Success") {
        toast.error(ce.error, toastStyle);
      }
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);
    }
  };

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
          noMargin={ce != ""}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">Code Editor</h1>
          </div>
          <div className="mx-6">
            <div className="w-full h-60 grid lg:grid-cols-2 lt-md:!grid-cols-1 items-center justify-center gap-4">
              <img
                src="https://cdn-botway.deno.dev/images/coder.svg"
                alt="Botway CE"
              />

              <div className="pb-4">
                <h2 className="text-md text-gray-400 text-center">
                  {project?.enable_ce
                    ? "Botway CE is enabled 👍"
                    : "Your project needs to enable Botway CE"}
                </h2>

                <h2 className="text-sm text-gray-500 text-center">
                  Botway CE is a code editor that built on top of{" "}
                  <Link
                    href="https://coder.com"
                    target="_blank"
                    className="text-blue-700"
                  >
                    Coder
                  </Link>
                  , with a lot of useful packages 🖥️
                </h2>

                {project?.enable_ce ? (
                  <Button
                    htmlType="submit"
                    onClick={enableCE}
                    type="success"
                    loading={isLoading}
                  >
                    Open My Editor
                  </Button>
                ) : (
                  <Button htmlType="submit" onClick={openModal} type="success">
                    Enable & Deploy
                  </Button>
                )}
              </div>
            </div>
          </div>

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
                    <Dialog.Panel className="w-full bg-bwdefualt bg-grid-gray-800/[0.4] max-w-md transform overflow-hidden rounded-2xl p-6 text-left align-middle shadow-xl transition-all border border-gray-800">
                      <Dialog.Title
                        as="h3"
                        className="text-lg font-medium pb-2 leading-6 text-white"
                      >
                        Add Password to your Code Editor
                      </Dialog.Title>
                      <div className="mt-2">
                        <Formik
                          initialValues={{
                            password: "",
                          }}
                          onSubmit={enableCE}
                        >
                          {({ errors }) => (
                            <Form className="column w-full">
                              <Field
                                className="input"
                                id="password"
                                name="password"
                                type="password"
                              />

                              {errors.password ? (
                                <div className="text-red-600 text-sm font-semibold pt-2">
                                  {errors.password}
                                </div>
                              ) : null}

                              <br />

                              <div className="flex w-full items-center gap-2 justify-end">
                                <div className="flex items-center gap-2">
                                  <button
                                    className="relative text-white cursor-pointer bg-blue-700 inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 shadow-sm text-xs px-2.5 py-1"
                                    type="submit"
                                  >
                                    {isLoading && (
                                      <LoadingDots
                                        className={styles.loading}
                                        children
                                      />
                                    )}
                                    <span className="truncate">Done</span>
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
        <CE user={user} projectId={params.id} />
      </QueryClientProvider>
    );
  }

  redirect("/");
};

export default ProjectPage;
