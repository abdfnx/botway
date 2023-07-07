"use client";

import { redirect } from "next/navigation";
import { DashLayout } from "@/components/Layout";
import { useAuth } from "@/supabase/auth/provider";
import { LoadingDots } from "@/components/LoadingDots";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import supabase from "@/supabase/browser";
import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import clsx from "clsx";
import { EncryptJWT } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { Dialog, Transition } from "@headlessui/react";
import {
  CheckIcon,
  MarkGithubIcon,
  PencilIcon,
  XCircleIcon,
} from "@primer/octicons-react";
import { Fragment, useState } from "react";

export const revalidate = 0;

const UpdateNameSchema = Yup.object().shape({
  name: Yup.string(),
});

const UpdateApiTokensSchema = Yup.object().shape({
  githubApiToken: Yup.string().min(40).max(40),
  railwayApiToken: Yup.string().min(36).max(36),
});

const Settings = () => {
  const { initial, user } = useAuth();
  const [isOpenGH, setIsOpenGH] = useState(false);
  const [isOpenRW, setIsOpenRW] = useState(false);

  const closeModalGH = () => {
    setIsOpenGH(false);
  };

  const openModalGH = () => {
    setIsOpenGH(true);
  };

  const closeModalRW = () => {
    setIsOpenRW(false);
  };

  const openModalRW = () => {
    setIsOpenRW(true);
  };

  if (initial) {
    return (
      <LoadingDots className="fixed inset-0 flex items-center justify-center" />
    );
  }

  async function updateName(formData: any) {
    const { error } = await supabase.auth.updateUser({
      data: { name: formData.name },
    });

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    } else {
      toast.success("Your info has been updated", toastStyle);
    }

    await supabase.auth.refreshSession();
  }

  async function updateGHApiTokens(formData: any) {
    const githubApiToken = await new EncryptJWT({
      data: formData.githubApiToken,
    })
      .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
      .encrypt(BW_SECRET_KEY);

    const { error } = await supabase.auth.updateUser({
      data: {
        githubApiToken,
      },
    });

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    } else {
      toast.success(
        "GitHub API Token has been updated successfully",
        toastStyle,
      );
    }

    await supabase.auth.refreshSession();

    closeModalGH();
  }

  async function updateRWApiTokens(formData: any) {
    const railwayApiToken = await new EncryptJWT({
      data: formData.railwayApiToken,
    })
      .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
      .encrypt(BW_SECRET_KEY);

    const { error } = await supabase.auth.updateUser({
      data: {
        railwayApiToken,
      },
    });

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    } else {
      toast.success(
        "Railway API Token has been updated successfully",
        toastStyle,
      );
    }

    await supabase.auth.refreshSession();

    closeModalRW();
  }

  if (user) {
    return (
      <DashLayout name={user.user_metadata["name"]} href="Settings">
        <div className="flex-1 flex-grow overflow-auto">
          <div className="p-4">
            <div className="my-2 pl-1">
              <div className="flex">
                <h3 className="text-xl text-white">
                  {user.user_metadata["name"]} Settings
                </h3>
              </div>
            </div>

            <div className="mb-9 mt-10 ml-1">
              <div className="my-4 max-w-4xl space-y-8">
                <Formik
                  initialValues={{
                    name: user.user_metadata["name"],
                  }}
                  validationSchema={UpdateNameSchema}
                  onSubmit={updateName}
                >
                  {({ values }) => (
                    <Form className="column w-full">
                      <div className="border-gray-800 overflow-hidden rounded-2xl border shadow">
                        <div className="flex flex-col bg-secondary gap-0">
                          <div className="grid grid-cols-12 gap-6 px-8 py-9 opacity-100">
                            <label className="text-white col-span-12 text-base lg:col-span-5">
                              General settings
                            </label>

                            <div className="relative col-span-12 flex flex-col gap-6 lg:col-span-7">
                              <div className="text-sm leading-4 grid gap-2 md:grid md:grid-cols-12">
                                <div className="flex flex-row space-x-2 justify-between col-span-12">
                                  <label
                                    className="block text-gray-400 pt-1 text-sm leading-4"
                                    htmlFor="name"
                                  >
                                    User Name
                                  </label>
                                </div>

                                <div className="col-span-12">
                                  <div className="relative">
                                    <Field
                                      className="input"
                                      id="name"
                                      name="name"
                                      type="text"
                                    />
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>

                        <div className="border-gray-800 border-t" />

                        <div className="flex py-4 px-8">
                          <div className="flex w-full items-center gap-2 justify-end">
                            <div className="flex items-center gap-2">
                              <button
                                className={clsx(
                                  "relative text-white bg-blue-700 inline-flex items-center space-x-2 text-center font-regular ease-out duration-200 rounded outline-none transition-all outline-0 focus-visible:outline-4 focus-visible:outline-offset-1 shadow-sm text-xs px-2.5 py-1",
                                  user.user_metadata["name"] === values.name
                                    ? "opacity-50 cursor-not-allowed pointer-events-none"
                                    : "cursor-pointer",
                                )}
                                type="submit"
                              >
                                <span className="truncate">Save</span>
                              </button>
                            </div>
                          </div>
                        </div>
                      </div>
                    </Form>
                  )}
                </Formik>
              </div>
            </div>

            <div className="mb-8 ml-1">
              <div className="my-4 max-w-4xl space-y-8">
                <div className="overflow-x-auto flex-grow border border-gray-800 rounded-lg">
                  <table className="w-full border-collapse select-auto">
                    <thead>
                      <tr className="bg-secondary border-b border-gray-800">
                        <th className="py-3 px-4 text-left font-semibold text-xs text-gray-400">
                          Platform
                        </th>
                        <th className="py-3 px-4 text-left font-semibold text-xs text-gray-400">
                          Token Status
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
                            <div className="text-gray-500" aria-hidden="true">
                              <MarkGithubIcon
                                className="fill-white"
                                size={18}
                              />
                            </div>

                            <p className="text-sm pt-0.5 text-white">
                              GitHub API Token
                            </p>
                          </div>
                        </td>

                        <td
                          className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500"
                          style={{ minWidth: "64px", maxWidth: "400px" }}
                        >
                          {user.user_metadata["githubApiToken"].length != 0 ? (
                            <>
                              <CheckIcon size={18} className="fill-green-600" />

                              <span className="pl-2 text-sm hidden md:inline">
                                Thank you for adding this token
                              </span>
                            </>
                          ) : (
                            <>
                              <XCircleIcon size={18} className="fill-red-600" />

                              <span className="pl-2 text-sm hidden md:inline">
                                You need to add this token
                              </span>
                            </>
                          )}
                        </td>

                        <td
                          className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500 w-[44px]"
                          style={{ minWidth: "64px", maxWidth: "400px" }}
                        >
                          <button
                            onClick={openModalGH}
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

                      <tr>
                        <td
                          className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap"
                          style={{ minWidth: "64px", maxWidth: "400px" }}
                        >
                          <div className="flex space-x-2 items-center">
                            <div className="text-gray-500" aria-hidden="true">
                              <img
                                src="https://cdn-botway.deno.dev/icons/railway.svg"
                                width={18}
                              />
                            </div>

                            <p className="text-sm text-white">
                              Railway API Token
                            </p>
                          </div>
                        </td>

                        <td
                          className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500"
                          style={{ minWidth: "64px", maxWidth: "400px" }}
                        >
                          {user.user_metadata["railwayApiToken"].length != 0 ? (
                            <>
                              <CheckIcon size={18} className="fill-green-600" />

                              <span className="pl-2 text-sm hidden md:inline">
                                Thank you for adding this token
                              </span>
                            </>
                          ) : (
                            <>
                              <XCircleIcon size={18} className="fill-red-600" />

                              <span className="pl-2 text-sm hidden md:inline">
                                You need to add this token
                              </span>
                            </>
                          )}
                        </td>

                        <td
                          className="py-3 px-4 overflow-hidden overflow-ellipsis whitespace-nowrap text-gray-500 w-[44px]"
                          style={{ minWidth: "64px", maxWidth: "400px" }}
                        >
                          <button
                            onClick={openModalRW}
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
          </div>
        </div>

        <Transition appear show={isOpenGH} as={Fragment}>
          <Dialog as="div" className="relative z-10" onClose={closeModalGH}>
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
                  <Dialog.Panel className="w-full bg-[#161b22] max-w-md transform overflow-hidden rounded-2xl p-6 text-left align-middle shadow-xl transition-all border border-gray-800">
                    <Dialog.Title
                      as="h3"
                      className="text-lg font-medium pb-2 leading-6 text-white"
                    >
                      Edit GitHub API Token
                    </Dialog.Title>
                    <div className="mt-2">
                      <Formik
                        initialValues={{
                          githubApiToken: "",
                        }}
                        validationSchema={UpdateApiTokensSchema}
                        onSubmit={updateGHApiTokens}
                      >
                        {({ errors }) => (
                          <Form className="column w-full">
                            <Field
                              className="transition-all bg-[#0d1117] border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-[#7d8590] block w-full p-2"
                              id="githubApiToken"
                              name="githubApiToken"
                              type="password"
                            />

                            {errors.githubApiToken ? (
                              <div className="text-red-600 text-sm font-semibold pt-2">
                                {errors.githubApiToken}
                              </div>
                            ) : null}
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

        <Transition appear show={isOpenRW} as={Fragment}>
          <Dialog as="div" className="relative z-10" onClose={closeModalRW}>
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
                  <Dialog.Panel className="w-full bg-[#1f132a] max-w-md transform overflow-hidden rounded-2xl p-6 text-left align-middle shadow-xl transition-all border border-gray-800">
                    <Dialog.Title
                      as="h3"
                      className="text-lg font-medium pb-2 leading-6 text-white"
                    >
                      Edit Railway API Token
                    </Dialog.Title>

                    <div className="mt-2">
                      <Formik
                        initialValues={{
                          railwayApiToken: "",
                        }}
                        validationSchema={UpdateApiTokensSchema}
                        onSubmit={updateRWApiTokens}
                      >
                        {({ errors }) => (
                          <Form className="column w-full">
                            <Field
                              className="transition-all bg-[#13111c] border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-[#853bce] block w-full p-2"
                              id="railwayApiToken"
                              name="railwayApiToken"
                              type="password"
                            />

                            {errors.railwayApiToken ? (
                              <div className="text-red-600 text-sm font-semibold pt-2">
                                {errors.railwayApiToken}
                              </div>
                            ) : null}
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
      </DashLayout>
    );
  }

  redirect("/");
};

export default Settings;
