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
import { fetcher } from "@/tools/fetch";
import { EncryptJWT, jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import {
  EyeClosedIcon,
  EyeIcon,
  PencilIcon,
  PlusIcon,
} from "@primer/octicons-react";
import { Fragment, useState } from "react";
import { Dialog, Transition } from "@headlessui/react";
import { Field, Form, Formik } from "formik";
import styles from "@/components/Button/Button.module.scss";
import * as Yup from "yup";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";

export const revalidate = 0;

const queryClient = new QueryClient();

const UpdateVarSchema = Yup.object().shape({
  value: Yup.string().min(1),
});

const AddVarSchema = Yup.object().shape({
  key: Yup.string().min(1),
  value: Yup.string().min(1),
});

const Env = ({ user, projectId }: any) => {
  const [show, setShow] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [isOpenAdd, setIsOpenAdd] = useState(false);
  const [isOpenUpdate, setIsOpenUpdate] = useState(false);
  const [currentVar, setCurrentVar] = useState("");

  const closeModalUpdate = () => {
    setIsOpenUpdate(false);
  };

  const openModalUpdate = (varx: any) => {
    setIsOpenUpdate(true);
    setCurrentVar(varx);
  };

  const closeModalAdd = () => {
    setIsOpenAdd(false);
  };

  const showEvent = (show: any, varx: any) => {
    setShow(show);
    setCurrentVar(varx);
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
    refetchInterval: 1,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const fetchVariables = async () => {
    const encVars = await fetcher(`/api/projects/env`, {
      method: "POST",
      body: JSON.stringify({
        projectId,
      }),
    });

    const { payload: vars } = await jwtDecrypt(encVars.vars, BW_SECRET_KEY);

    return vars.data;
  };

  const {
    data: vars,
    isLoading: varsIsLoading,
    refetch,
  }: any = useQuery({
    queryKey: ["services"],
    queryFn: fetchVariables,
    refetchInterval: 60000,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const updateVar = async (formData: any) => {
    try {
      setIsLoading(true);

      const value = await new EncryptJWT({
        data: formData.value,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      console.log(vars);

      const vx = await new EncryptJWT({
        data: vars,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      const body = {
        value,
        vars: vx,
        key: currentVar,
        projectId,
      };

      await fetcher("/api/projects/env/update", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      closeModalUpdate();

      refetch();
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);
    }
  };

  const addVar = async (formData: any) => {
    try {
      setIsLoading(true);

      const value = await new EncryptJWT({
        data: formData.value,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      const body = {
        value,
        key: formData.key,
        projectId,
      };

      await fetcher("/api/projects/env/add", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      closeModalAdd();

      refetch();
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
          grid={true}
        >
          <div className="mx-6 my-16 flex items-center space-x-6">
            <h1 className="text-3xl text-white">
              {project?.name} Environment Variables
            </h1>
          </div>

          {varsIsLoading ? (
            <LoadingDots className="fixed inset-0 flex items-center justify-center" />
          ) : (
            <div className="mx-6 my-4 max-w-full space-y-8">
              <div className="flex justify-between items-center">
                <h1 className="text-lg text-white">
                  {Object.keys(vars?.vars).length} Service Variables
                </h1>

                <button
                  type="button"
                  onClick={() => setIsOpenAdd(true)}
                  className="flex border-gray-800 bg-bwdefualt hover:bg-secondary border p-0.5 outline-none ease-out duration-200 rounded-md outline-offset-1 transition-all focus:outline-4"
                >
                  <span className="relative cursor-pointer inline-flex items-center place-content-center space-x-2 text-center font-regular  outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 text-blue-700 shadow-sm px-2.5 py-1">
                    <PlusIcon />

                    <span className="hidden text-sm text-gray-400 md:block">
                      New Variable
                    </span>
                  </span>
                </button>
              </div>

              <div className="overflow-x-auto flex-grow rounded-lg border border-gray-800">
                <table className="w-full border-collapse select-auto bg-secondary">
                  <tbody>
                    {vars?.vars.map((key: any, index: any) => (
                      <tr className="justify-between">
                        <td
                          className={`py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap border-r border-gray-800 ${
                            index != vars?.vars.length - 1 ? "border-b" : ""
                          }`}
                          style={{ minWidth: "64px", maxWidth: "100px" }}
                        >
                          <div className="flex text-sm font-semibold leading-6 text-white space-x-2 items-center font-mono">
                            {key.key}
                          </div>
                        </td>

                        <td
                          className={`py-3 px-4 overflow-hidden hidden md:table-cell overflow-ellipsis whitespace-nowrap ${
                            index != vars?.vars.length - 1
                              ? "border-b border-gray-800"
                              : ""
                          } `}
                          style={{ minWidth: "64px", maxWidth: "250px" }}
                        >
                          <div className="flex justify-between">
                            <div className="flex text-sm font-medium leading-6 text-white space-x-2 items-center font-mono">
                              {show && currentVar === key
                                ? vars?.vars[index].value
                                : "•".repeat(vars?.vars[index].value.length)}

                              {show && currentVar === key ? (
                                <div
                                  className="cursor-pointer"
                                  onClick={() => showEvent(!show, key)}
                                >
                                  <EyeClosedIcon
                                    className="ml-4 fill-white"
                                    size={18}
                                  />
                                </div>
                              ) : (
                                <div
                                  className="cursor-pointer"
                                  onClick={() => showEvent(!show, key)}
                                >
                                  <EyeIcon
                                    className="ml-4 fill-white"
                                    size={18}
                                  />
                                </div>
                              )}
                            </div>

                            <button
                              onClick={() => openModalUpdate(key)}
                              className="flex items-center justify-center transform transition-transform duration-50 active:scale-95 focus:outline-none focus-visible:ring-2 focus-visible:outline-none h-[34px] py-1.5 rounded-md text-sm leading-5 space-x-2 w-[34px] px-0"
                            >
                              <div className="!w-4 !h-4" aria-hidden="true">
                                <PencilIcon
                                  size={18}
                                  className="hover:fill-gray-400 fill-gray-500 transition-all duration-200"
                                />
                              </div>
                            </button>
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          )}

          <Transition appear show={isOpenUpdate} as={Fragment}>
            <Dialog
              as="div"
              className="relative z-10"
              onClose={closeModalUpdate}
            >
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
                        Update
                      </Dialog.Title>
                      <div className="mt-2">
                        <Formik
                          initialValues={{
                            value: vars?.vars[currentVar],
                          }}
                          validationSchema={UpdateVarSchema}
                          onSubmit={updateVar}
                        >
                          {({ errors }) => (
                            <Form className="column w-full">
                              <Field
                                className="input font-mono"
                                id="value"
                                name="value"
                                type="text"
                              />

                              {errors.value ? (
                                <div className="text-red-600 text-sm font-semibold pt-2">
                                  {errors.value.toString()}
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

          <Transition appear show={isOpenAdd} as={Fragment}>
            <Dialog as="div" className="relative z-10" onClose={closeModalAdd}>
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
                        New Variable
                      </Dialog.Title>
                      <div className="mt-2">
                        <Formik
                          initialValues={{
                            key: "",
                            value: "",
                          }}
                          validationSchema={AddVarSchema}
                          onSubmit={addVar}
                        >
                          {({ errors }) => (
                            <Form className="column w-full">
                              <div>
                                <label className="text-white col-span-12 text-base lg:col-span-5">
                                  Variable Name
                                </label>

                                <div className="pt-2" />

                                <Field
                                  className="input font-mono"
                                  id="key"
                                  name="key"
                                  type="text"
                                />

                                {errors.key ? (
                                  <div className="text-red-600 text-sm font-semibold pt-2">
                                    {errors.key.toString()}
                                  </div>
                                ) : null}
                              </div>

                              <br />

                              <div>
                                <label className="text-white col-span-12 text-base lg:col-span-5">
                                  Value
                                </label>

                                <div className="pt-2" />

                                <Field
                                  className="input font-mono"
                                  id="value"
                                  name="value"
                                  type="text"
                                />

                                {errors.value ? (
                                  <div className="text-red-600 text-sm font-semibold pt-2">
                                    {errors.value.toString()}
                                  </div>
                                ) : null}
                              </div>

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
                                    <span className="truncate">Add</span>
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
        <Env user={user} projectId={params.id} />
      </QueryClientProvider>
    );
  }

  redirect("/");
};

export default ProjectPage;
