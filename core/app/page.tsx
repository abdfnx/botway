"use client";

import Auth from "@/supabase/auth/container";
import { useAuth, VIEWS } from "@/supabase/auth/provider";
import { LoadingDots } from "@/components/LoadingDots";
import { DashLayout, mode } from "@/components/Layout";
import { UserAvatar } from "@/components/UserAvatar";
import {
  CheckIcon,
  ChevronDownIcon,
  CommandPaletteIcon,
  SparkleFillIcon,
  StackIcon,
  ZapIcon,
} from "@primer/octicons-react";
import { Fragment, useRef, useState } from "react";
import { Dialog, Listbox, Menu, Transition } from "@headlessui/react";
import { Field, Form, Formik } from "formik";
import supabase from "@/supabase/browser";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import {
  langs,
  packageManagers,
  platforms,
  PLP,
  visibilityOptions,
} from "@/tools/new/project-options";
import clsx from "clsx";
import * as Yup from "yup";
import { useQuery, QueryClientProvider } from "@tanstack/react-query";
import { fetcher } from "@/tools/fetch";
import { Button } from "@/components/Button";
import { capitalizeFirstLetter } from "@/tools/text";
import { Badge } from "@tremor/react";
import { queryClient } from "@/tools/qc";

export const revalidate = 0;

const AddNewProjectSchema = Yup.object().shape({
  name: Yup.string().min(3),
});

const Home = ({ user }: any) => {
  const [open, setOpen] = useState(false);
  const [promptOpen, setPromptOpen] = useState(false);

  const platformRef: any = useRef();
  const langRef: any = useRef();
  const packageManagerRef: any = useRef();
  const visibilityRef: any = useRef();
  const [isLoading, setIsLoading] = useState(false);

  const fetchProjects = async () => {
    const { data: projects } = await supabase
      .from("projects")
      .select("*")
      .order("created_at");

    return projects;
  };

  const { data: projects, isLoading: projectIsLoading } = useQuery({
    queryKey: ["project"],
    queryFn: fetchProjects,
    refetchInterval: 1,
    refetchOnReconnect: true,
    refetchOnWindowFocus: true,
    refetchIntervalInBackground: true,
  });

  const [visibilitySelected, setvisibilitySelected]: any = useState(
    visibilityOptions[0],
  );

  const [platformSelected, setPlatformSelected]: any = useState(platforms[0]);

  const [langSelected, setLangSelected] = useState(
    langs(platformSelected.name)[0],
  );

  const [pmSelected, setPMSelected] = useState(
    packageManagers(langSelected.name)[0],
  );

  const addNewProject = async (formData: any) => {
    try {
      setIsLoading(true);

      if (
        visibilityRef.current.value.toLowerCase() != "choose" &&
        platformRef.current.value.toLowerCase() != "choose" &&
        langRef.current.value.toLowerCase() != "choose" &&
        packageManagerRef.current.value.toLowerCase() != "choose" &&
        PLP[platformRef.current.value][langRef.current.value] != null &&
        PLP[platformRef.current.value][langRef.current.value].pm.includes(
          packageManagerRef.current.value,
        )
      ) {
        const body = {
          name: formData.name,
          visibility: visibilityRef.current.value,
          platform: platformRef.current.value,
          lang: langRef.current.value,
          package_manager: packageManagerRef.current.value,
        };

        const newBot = await fetcher("/api/projects", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(body),
        });

        if (newBot.message === "Success") {
          toast.success(
            "You have successfully created a new bot project",
            toastStyle,
          );

          setOpen(false);
        } else {
          toast.error(newBot.error, toastStyle);

          setOpen(false);
        }
      } else {
        toast.error("Choose the right choice(s)", toastStyle);
      }
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);
    }
  };

  const submitPrompt = async (formData: any) => {
    try {
      setIsLoading(true);

      const newBot = await fetcher("/api/ai", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ prompt: formData.prompt }),
      });

      if (newBot.message === "Success") {
        toast.success(
          "You have successfully created a new bot project",
          toastStyle,
        );

        setPromptOpen(false);
      } else {
        toast.error(newBot.error, toastStyle);

        setPromptOpen(false);
      }
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <DashLayout user={user} href="Projects">
      <div className="flex-1 flex-grow overflow-auto">
        <div className="py-4 px-5">
          <div className="my-2">
            <div className="border bg-grid-gray-800 flex items-center text-white place-content-center justify-center border-gray-800 p-12 rounded-xl">
              <div className="items-center">
                <h2 className="text-base md:text-xl font-farray mr-3">
                  <span className="font-bold">Welcome To Botway 👋</span>
                </h2>
              </div>
            </div>
          </div>
          <div className="my-5 space-y-8">
            <div className="space-y-3">
              <div className="flex items-center justify-between pb-2 gap-4">
                <div className="flex-1 gap-2 justify-end flex-shrink-0">
                  <a className="h-9 mt-1 px-4.5 inline-flex flex-shrink-0 whitespace-nowrap items-center gap-2">
                    <UserAvatar data={user.email} size={32} />
                    <span className="text-white font-semibold font-mono text-xl md:text-2xl">
                      {user.user_metadata["name"]}
                    </span>
                    <Badge color="cyan" className="hidden md:block">
                      <span className="font-mono">{mode()} Mode</span>
                    </Badge>
                  </a>
                </div>
                <div className="flex gap-2 justify-end flex-shrink-0">
                  <Menu as="div" className="relative inline-block text-left">
                    <div>
                      <Menu.Button className="h-9 px-2 py-3.5 rounded-lg border border-gray-800 inline-flex flex-shrink-0 whitespace-nowrap items-center gap-2 transition-colors duration-200 ease-in-out leading-none cursor-pointer text-white hover:bg-secondary focus:outline-none outline-none">
                        <ZapIcon size={16} className="fill-blue-700" />
                        New Project
                      </Menu.Button>
                    </div>

                    <Transition
                      as={Fragment}
                      enter="transition ease-out duration-100"
                      enterFrom="transform opacity-0 scale-95"
                      enterTo="transform opacity-100 scale-100"
                      leave="transition ease-in duration-75"
                      leaveFrom="transform opacity-100 scale-100"
                      leaveTo="transform opacity-0 scale-95"
                    >
                      <Menu.Items className="absolute right-0 mt-2 w-56 origin-top-right border border-gray-800 rounded-lg bg-secondary shadow-lg focus:outline-none">
                        <div className="px-1 py-1">
                          <Menu.Item>
                            {({ active }) => (
                              <button
                                onClick={() => setOpen(true)}
                                className={`${
                                  active ? "bg-bwdefualt" : ""
                                } group flex w-full text-white items-center rounded-lg px-2 py-2 text-sm`}
                              >
                                <StackIcon
                                  size={18}
                                  className="fill-blue-700 mr-2"
                                />
                                <h1 className="font-mono">Botway Templates</h1>
                              </button>
                            )}
                          </Menu.Item>
                          <Menu.Item>
                            {() => (
                              <button
                                className={`cursor-context-menu group flex w-full text-white items-center rounded-lg px-2 py-2 text-sm`}
                                // onClick={() => setPromptOpen(true)}
                              >
                                <SparkleFillIcon
                                  size={18}
                                  className="fill-blue-700 mr-2"
                                />
                                <h1 className="font-mono">Botway AI - Soon</h1>
                              </button>
                            )}
                          </Menu.Item>
                        </div>
                      </Menu.Items>
                    </Transition>
                  </Menu>
                </div>
              </div>

              {projectIsLoading ? (
                <LoadingDots className="fixed inset-0 flex items-center justify-center" />
              ) : projects?.length != 0 ? (
                <div className="mt-10 grid lg:grid-cols-3 sm:grid-cols-2 lt-md:!grid-cols-1 gap-3">
                  {projects?.map((project) => (
                    <div className="col-span-1">
                      <a href={`/project/${project.id}`}>
                        <div className="group text-left border-2 border-dashed border-gray-800 rounded-xl py-4 px-6 flex flex-row transition duration-150 h-32 cursor-pointer hover:bg-secondary">
                          <div className="flex h-full w-full flex-col space-y-2">
                            <h5 className="text-white">
                              <div className="flex w-full items-center flex-row justify-between gap-1">
                                <span className="flex-shrink truncate">
                                  {project.name}
                                </span>

                                <CommandPaletteIcon
                                  size={18}
                                  className="fill-blue-700"
                                />
                              </div>
                            </h5>
                            <br />
                            <div className="w-full">
                              <p className="flex items-center gap-1.5 mt-3 text-sm text-gray-500">
                                <img
                                  src={`https://cdn-botway.deno.dev/icons/${project.platform}.svg`}
                                  alt={`${project.platform} icon`}
                                  width={16}
                                />
                                {capitalizeFirstLetter(project.platform)}

                                <img
                                  src={`https://cdn-botway.deno.dev/icons/${project.lang}.svg`}
                                  alt={`${project.lang} icon`}
                                  width={16}
                                />
                                {capitalizeFirstLetter(project.lang)}
                              </p>
                            </div>
                          </div>
                        </div>
                      </a>
                    </div>
                  ))}
                </div>
              ) : (
                <div className="rounded-xl mt-8 overflow-hidden p-5 cursor-pointer border-2 border-dashed border-gray-800 hover:border-gray-700 transition duration-300 ease-in-out w-full h-80 flex flex-col items-center justify-center gap-4">
                  <h2 className="text-md text-gray-400 text-center">
                    Create a New Project
                  </h2>
                  <button
                    onClick={() => setOpen(true)}
                    className="h-9 px-2 py-3.5 rounded-lg border border-gray-800 inline-flex flex-shrink-0 whitespace-nowrap items-center gap-2 transition-colors duration-200 ease-in-out leading-none cursor-pointer text-white hover:bg-secondary focus:outline-none outline-none"
                  >
                    <ZapIcon size={16} className="fill-blue-700" />
                    New Project
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>

      <Transition.Root show={open} as={Fragment}>
        <Dialog
          as="div"
          className="relative z-10"
          onClose={() => setOpen(false)}
        >
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
                          Create a new Bot Project
                        </Dialog.Title>
                      </div>
                      <div className="relative mt-4 flex-1 px-4 sm:px-6">
                        <div className="my-4 max-w-4xl space-y-8">
                          <Formik
                            initialValues={{
                              name: "",
                            }}
                            validationSchema={AddNewProjectSchema}
                            onSubmit={addNewProject}
                          >
                            {({ errors }) => (
                              <>
                                <Form className="column w-full">
                                  <div>
                                    <label className="text-white col-span-12 text-base lg:col-span-5">
                                      Bot Name
                                    </label>

                                    <div className="pt-2" />

                                    <Field
                                      className="input"
                                      id="name"
                                      name="name"
                                      type="text"
                                    />

                                    {errors.name ? (
                                      <div className="text-red-600 text-sm font-semibold pt-2">
                                        {errors.name}
                                      </div>
                                    ) : null}
                                  </div>
                                  <br />
                                  <div>
                                    <label className="text-white col-span-12 text-base lg:col-span-5">
                                      Platform
                                    </label>

                                    <div className="pt-2" />

                                    <Listbox
                                      value={platformSelected}
                                      onChange={setPlatformSelected}
                                      name="platform"
                                    >
                                      {({ open }) => (
                                        <>
                                          <div className="relative">
                                            <Listbox.Button className="relative w-full cursor-pointer rounded-lg border border-gray-800 bg-bwdefualt py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                                              <span className="flex items-center">
                                                <img
                                                  src={`https://cdn-botway.deno.dev/icons/${platformSelected.slug}.svg`}
                                                  alt={`${platformSelected.slug} icon`}
                                                  className="h-6 w-6 flex-shrink-0"
                                                />
                                                <span className="ml-3 block text-white truncate">
                                                  {platformSelected.name}
                                                </span>
                                              </span>
                                              <span className="pointer-events-none absolute inset-y-0 right-0 ml-3 flex items-center pr-2">
                                                <ChevronDownIcon
                                                  className="h-5 w-5 text-gray-400"
                                                  aria-hidden="true"
                                                />
                                              </span>
                                            </Listbox.Button>

                                            <Transition
                                              show={open}
                                              as={Fragment}
                                              leave="transition ease-in duration-100"
                                              leaveFrom="opacity-100"
                                              leaveTo="opacity-0"
                                            >
                                              <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-lg bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                                                {platforms.map((platform) => (
                                                  <Listbox.Option
                                                    key={platform.name}
                                                    onChange={() =>
                                                      setPlatformSelected(
                                                        platform,
                                                      )
                                                    }
                                                    className={({ active }) =>
                                                      clsx(
                                                        active
                                                          ? "text-white bg-secondary"
                                                          : "text-gray-500",
                                                        "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-lg mx-2 my-1",
                                                      )
                                                    }
                                                    value={platform}
                                                  >
                                                    {({ selected, active }) => (
                                                      <>
                                                        <div className="flex items-center">
                                                          <img
                                                            src={`https://cdn-botway.deno.dev/icons/${platform.slug}.svg`}
                                                            alt={`${platform.slug} icon`}
                                                            className="h-6 w-6 flex-shrink-0"
                                                            width={16}
                                                          />
                                                          <span
                                                            className={clsx(
                                                              selected
                                                                ? "font-semibold"
                                                                : "font-normal",
                                                              "ml-3 block truncate",
                                                            )}
                                                          >
                                                            {platform.name}
                                                          </span>
                                                        </div>

                                                        {selected ? (
                                                          <span
                                                            className={clsx(
                                                              active
                                                                ? "text-white"
                                                                : "text-blue-700",
                                                              "absolute inset-y-0 right-0 flex items-center pr-4",
                                                            )}
                                                          >
                                                            <CheckIcon
                                                              className="h-5 w-5"
                                                              aria-hidden="true"
                                                            />
                                                          </span>
                                                        ) : null}
                                                      </>
                                                    )}
                                                  </Listbox.Option>
                                                ))}
                                              </Listbox.Options>
                                            </Transition>
                                          </div>
                                        </>
                                      )}
                                    </Listbox>

                                    <input
                                      type="hidden"
                                      name="platform[name]"
                                      value={platformSelected.slug}
                                      ref={platformRef}
                                    />
                                  </div>
                                  <br />
                                  <div>
                                    <label className="text-white col-span-12 text-base lg:col-span-5">
                                      Programming Language
                                    </label>

                                    <div className="pt-2" />

                                    <Listbox
                                      value={langSelected}
                                      onChange={setLangSelected}
                                    >
                                      {({ open }) => (
                                        <>
                                          <div className="relative">
                                            <Listbox.Button className="relative w-full cursor-pointer rounded-lg border border-gray-800 bg-bwdefualt py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                                              <span className="flex items-center">
                                                <img
                                                  src={`https://cdn-botway.deno.dev/icons/${langSelected.slug}.svg`}
                                                  alt={`${langSelected.slug} icon`}
                                                  className="h-6 w-6 flex-shrink-0"
                                                />
                                                <span className="ml-3 block text-white truncate">
                                                  {langSelected.name}
                                                </span>
                                              </span>
                                              <span className="pointer-events-none absolute inset-y-0 right-0 ml-3 flex items-center pr-2">
                                                <ChevronDownIcon
                                                  className="h-5 w-5 text-gray-400"
                                                  aria-hidden="true"
                                                />
                                              </span>
                                            </Listbox.Button>

                                            <Transition
                                              show={open}
                                              as={Fragment}
                                              leave="transition ease-in duration-100"
                                              leaveFrom="opacity-100"
                                              leaveTo="opacity-0"
                                            >
                                              <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-lg bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                                                {langs(
                                                  platformSelected.name,
                                                ).map((lang) => (
                                                  <Listbox.Option
                                                    key={lang.name}
                                                    className={({ active }) =>
                                                      clsx(
                                                        active
                                                          ? "text-white bg-secondary"
                                                          : "text-gray-500",
                                                        "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-lg mx-2 my-1",
                                                      )
                                                    }
                                                    value={lang}
                                                  >
                                                    {({ selected, active }) => (
                                                      <>
                                                        <div className="flex items-center">
                                                          <img
                                                            src={`https://cdn-botway.deno.dev/icons/${lang.slug}.svg`}
                                                            alt={`${lang.slug} icon`}
                                                            className="h-6 w-6 flex-shrink-0"
                                                            width={16}
                                                          />
                                                          <span
                                                            className={clsx(
                                                              selected
                                                                ? "font-semibold"
                                                                : "font-normal",
                                                              "ml-3 block truncate",
                                                            )}
                                                          >
                                                            {lang.name}
                                                          </span>
                                                        </div>

                                                        {selected ? (
                                                          <span
                                                            className={clsx(
                                                              active
                                                                ? "text-white"
                                                                : "text-blue-700",
                                                              "absolute inset-y-0 right-0 flex items-center pr-4",
                                                            )}
                                                          >
                                                            <CheckIcon
                                                              className="h-5 w-5"
                                                              aria-hidden="true"
                                                            />
                                                          </span>
                                                        ) : null}
                                                      </>
                                                    )}
                                                  </Listbox.Option>
                                                ))}
                                              </Listbox.Options>
                                            </Transition>
                                          </div>
                                        </>
                                      )}
                                    </Listbox>

                                    <input
                                      type="hidden"
                                      name="lang[name]"
                                      value={langSelected.slug}
                                      ref={langRef}
                                    />
                                  </div>
                                  <br />
                                  <div>
                                    <label className="text-white col-span-12 text-base lg:col-span-5">
                                      Package Manager
                                    </label>

                                    <div className="pt-2" />

                                    <Listbox
                                      value={pmSelected}
                                      refName={packageManagerRef}
                                      onChange={setPMSelected}
                                    >
                                      {({ open }) => (
                                        <>
                                          <div className="relative">
                                            <Listbox.Button className="relative w-full cursor-pointer rounded-lg border border-gray-800 bg-bwdefualt py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                                              <span className="flex items-center">
                                                <img
                                                  src={`https://cdn-botway.deno.dev/icons/${pmSelected.logo}`}
                                                  alt={`${pmSelected.logo} icon`}
                                                  className="h-6 w-6 flex-shrink-0"
                                                />
                                                <span className="ml-3 block text-white truncate">
                                                  {pmSelected.name}
                                                </span>
                                              </span>
                                              <span className="pointer-events-none absolute inset-y-0 right-0 ml-3 flex items-center pr-2">
                                                <ChevronDownIcon
                                                  className="h-5 w-5 text-gray-400"
                                                  aria-hidden="true"
                                                />
                                              </span>
                                            </Listbox.Button>

                                            <Transition
                                              show={open}
                                              as={Fragment}
                                              leave="transition ease-in duration-100"
                                              leaveFrom="opacity-100"
                                              leaveTo="opacity-0"
                                            >
                                              <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-lg bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                                                {packageManagers(
                                                  langSelected.name,
                                                ).map((pm) => (
                                                  <Listbox.Option
                                                    key={pm.name}
                                                    className={({ active }) =>
                                                      clsx(
                                                        active
                                                          ? "text-white bg-secondary"
                                                          : "text-gray-500",
                                                        "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-lg mx-2 my-1",
                                                      )
                                                    }
                                                    value={pm}
                                                  >
                                                    {({ selected, active }) => (
                                                      <>
                                                        <div className="flex items-center">
                                                          <img
                                                            src={`https://cdn-botway.deno.dev/icons/${pm.logo}`}
                                                            alt={`${pm.name} icon`}
                                                            className="h-6 w-6 flex-shrink-0"
                                                            width={16}
                                                          />
                                                          <span
                                                            className={clsx(
                                                              selected
                                                                ? "font-semibold"
                                                                : "font-normal",
                                                              "ml-3 block truncate",
                                                            )}
                                                          >
                                                            {pm.name}
                                                          </span>
                                                        </div>

                                                        {selected ? (
                                                          <span
                                                            className={clsx(
                                                              active
                                                                ? "text-white"
                                                                : "text-blue-700",
                                                              "absolute inset-y-0 right-0 flex items-center pr-4",
                                                            )}
                                                          >
                                                            <CheckIcon
                                                              className="h-5 w-5"
                                                              aria-hidden="true"
                                                            />
                                                          </span>
                                                        ) : null}
                                                      </>
                                                    )}
                                                  </Listbox.Option>
                                                ))}
                                              </Listbox.Options>
                                            </Transition>
                                          </div>
                                        </>
                                      )}
                                    </Listbox>

                                    <input
                                      type="hidden"
                                      name="pm[name]"
                                      value={pmSelected.name}
                                      ref={packageManagerRef}
                                    />
                                  </div>
                                  <br />
                                  <div>
                                    <label className="text-white col-span-12 text-base lg:col-span-5">
                                      Visibility On GitHub
                                    </label>

                                    <div className="pt-2" />

                                    <Listbox
                                      value={visibilitySelected}
                                      refName={visibilityRef}
                                      onChange={setvisibilitySelected}
                                    >
                                      {({ open }) => (
                                        <>
                                          <div className="relative">
                                            <Listbox.Button className="relative w-full cursor-pointer rounded-lg border border-gray-800 bg-bwdefualt py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                                              <span className="flex items-center">
                                                <span className="ml-2 text-white block truncate">
                                                  {visibilitySelected.typeName}
                                                </span>
                                              </span>
                                              <span className="pointer-events-none absolute inset-y-0 right-0 ml-3 flex items-center pr-2">
                                                <ChevronDownIcon
                                                  className="h-5 w-5 text-gray-400"
                                                  aria-hidden="true"
                                                />
                                              </span>
                                            </Listbox.Button>

                                            <Transition
                                              show={open}
                                              as={Fragment}
                                              leave="transition ease-in duration-100"
                                              leaveFrom="opacity-100"
                                              leaveTo="opacity-0"
                                            >
                                              <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-lg bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                                                {visibilityOptions.map(
                                                  (visibility) => (
                                                    <Listbox.Option
                                                      key={visibility.typeName}
                                                      className={({ active }) =>
                                                        clsx(
                                                          active
                                                            ? "text-white bg-secondary"
                                                            : "text-gray-500",
                                                          "relative transition cursor-pointer select-none py-2 pl-2 pr-9 rounded-lg mx-2 my-1",
                                                        )
                                                      }
                                                      value={visibility}
                                                    >
                                                      {({
                                                        selected,
                                                        active,
                                                      }) => (
                                                        <>
                                                          <div className="flex items-center">
                                                            <span
                                                              className={clsx(
                                                                selected
                                                                  ? "font-semibold"
                                                                  : "font-normal",
                                                                "ml-3 block truncate",
                                                              )}
                                                            >
                                                              {
                                                                visibility.typeName
                                                              }
                                                            </span>
                                                          </div>

                                                          {selected ? (
                                                            <span
                                                              className={clsx(
                                                                active
                                                                  ? "text-white"
                                                                  : "text-blue-700",
                                                                "absolute inset-y-0 right-0 flex items-center pr-4",
                                                              )}
                                                            >
                                                              <CheckIcon
                                                                className="h-5 w-5"
                                                                aria-hidden="true"
                                                              />
                                                            </span>
                                                          ) : null}
                                                        </>
                                                      )}
                                                    </Listbox.Option>
                                                  ),
                                                )}
                                              </Listbox.Options>
                                            </Transition>
                                          </div>
                                        </>
                                      )}
                                    </Listbox>

                                    <input
                                      type="hidden"
                                      name="visibility[typeName]"
                                      value={visibilitySelected.type}
                                      ref={visibilityRef}
                                    />
                                  </div>

                                  <Button
                                    htmlType="submit"
                                    type="success"
                                    loading={isLoading}
                                  >
                                    Create Bot Project
                                  </Button>
                                </Form>

                                <a
                                  href="https://zeabur.com"
                                  target="_blank"
                                  className="mt-4 border border-gray-800 transition-all bg-[#121212] hover:bg-[#141414] duration-200 rounded-2xl p-4 sm:mt-8 flex flex-col items-center"
                                >
                                  <div aria-hidden="true">
                                    <img
                                      src="https://cdn-botway.deno.dev/icons/zeabur.svg"
                                      width={30}
                                    />
                                  </div>
                                  <div className="space-y-2 mt-3 sm:space-y-4 flex flex-col items-center">
                                    <h1 className="text-white text-xs md:text-sm font-bold">
                                      Your Bot Project will be hosted on Zeabur
                                    </h1>
                                    <p className="text-xs md:text-sm text-gray-400 text-center">
                                      Zeabur is a platform that help you deploy
                                      your service with one click. 🏗️
                                    </p>
                                  </div>
                                </a>
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

      <Transition appear show={promptOpen} as={Fragment}>
        <Dialog
          as="div"
          className="relative z-10"
          onClose={() => setPromptOpen(false)}
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
            <div className="flex min-h-full mx-4 md:mx-32 items-center justify-center p-4 text-center">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 scale-95"
                enterTo="opacity-100 scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 scale-100"
                leaveTo="opacity-0 scale-95"
              >
                <Dialog.Panel className="min-w-full bg-secondary transform overflow-scroll rounded-2xl p-0 text-left align-middle shadow-xl transition-all border border-gray-800">
                  <div>
                    {promptOpen ? (
                      isLoading ? (
                        <Transition.Child
                          as={Fragment}
                          enter="ease-out duration-300"
                          enterFrom="opacity-0"
                          enterTo="opacity-100"
                          leave="ease-in duration-200"
                          leaveFrom="opacity-100"
                          leaveTo="opacity-0"
                        >
                          <div
                            className="gap-4 flex justify-between items-center"
                            style={{ margin: "22px" }}
                          >
                            <h1 className="text-sm md:text-xl text-white">
                              Creating
                            </h1>
                            <LoadingDots />
                          </div>
                        </Transition.Child>
                      ) : (
                        <Formik
                          initialValues={{
                            prompt: "",
                          }}
                          onSubmit={submitPrompt}
                        >
                          {() => (
                            <Form className="mt-2 column overflow-scroll min-h-full">
                              <Field
                                className="border-none outline-none focus:outline-none focus-within:outline-none bg-secondary text-sm md:text-xl placeholder:text-sm md:placeholder:text-xl placeholder:text-gray-400 text-white rounded-sm block w-full p-4"
                                id="prompt"
                                name="prompt"
                                type="text"
                                placeholder="Build your own bot using AI, you need to specify bot platform, language (default: python), and package manager (optional)"
                              />
                            </Form>
                          )}
                        </Formik>
                      )
                    ) : (
                      <div className="items-center" style={{ margin: "22px" }}>
                        <h1 className="text-sm md:text-xl text-white">
                          Done 🤝
                        </h1>
                      </div>
                    )}
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition>
    </DashLayout>
  );
};

const App = () => {
  const { initial, user, view } = useAuth();

  if (initial) {
    return (
      <LoadingDots className="fixed inset-0 flex items-center justify-center" />
    );
  }

  if (view === VIEWS.UPDATE_PASSWORD) {
    return <Auth view={view} />;
  }

  if (user) {
    return (
      <QueryClientProvider client={queryClient}>
        <Home user={user} />
      </QueryClientProvider>
    );
  }

  return <Auth view={view} />;
};

export default App;
