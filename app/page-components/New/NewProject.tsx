import { fetcher } from "@/lib/fetch";
import { Transition, Listbox } from "@headlessui/react";
import { Fragment, useCallback, useRef, useState } from "react";
import toast from "react-hot-toast";
import clsx from "clsx";
import { Button } from "@/components/Button";
import {
  platforms,
  langs,
  hostServices,
  packageManagers,
  visibilityOptions,
} from "./Options";
import { NewProjectModal } from "./NewProjectModal";
import { CheckIcon, ChevronDownIcon } from "@primer/octicons-react";
import { useCurrentUser } from "@/lib/user";
import { Octokit } from "octokit";
import { BW_SECRET_KEY, CheckAPITokens } from "@/tools/api-tokens";
import { toastStyle } from "@/tools/toast-style";
import { jwtDecrypt } from "jose";

const NewProjectHandler = () => {
  const { data: { user } = {}, mutate } = useCurrentUser();

  const nameRef: any = useRef();
  const platformRef: any = useRef();
  const langRef: any = useRef();
  const packageManagerRef: any = useRef();
  const hostServiceRef: any = useRef();
  const visibilityRef: any = useRef();

  const [isLoading, setIsLoading] = useState(false);

  let [visibilitySelected, setvisibilitySelected]: any = useState(
    visibilityOptions[0]
  );
  let [platformSelected, setPlatformSelected]: any = useState(platforms[0]);
  let [langSelected, setLangSelected] = useState(
    langs(platformSelected.name)[0]
  );
  let [hostServiceSelected, setHostServiceSelected] = useState(hostServices[0]);
  let [pmSelected, setPMSelected] = useState(
    packageManagers(langSelected.name)[0]
  );

  const onSubmit = useCallback(
    async (e: any) => {
      e.preventDefault();

      try {
        setIsLoading(true);

        CheckAPITokens(user, "");

        const { payload } = await jwtDecrypt(
          user.githubApiToken,
          BW_SECRET_KEY
        );

        const octokit = new Octokit({
          auth: payload.data,
        });

        const ghu = await (await octokit.request("GET /user", {})).data;

        let body = {
          railwayApiToken: user.railwayApiToken,
          ghToken: user.githubApiToken,
          userId: user._id,
          name: nameRef.current.value,
          botToken: "",
          botAppToken: "",
          botSecretToken: "",
          visibility: visibilityRef.current.value,
          platform: platformRef.current.value,
          lang: langRef.current.value,
          packageManager: packageManagerRef.current.value,
          hostService: hostServiceRef.current.value,
          repo: `${ghu.login}/${nameRef.current.value}`,
          repoBranch: "",
          pullRequestPreviewsEnabled: "",
        };

        if (hostServiceRef.current.value == "render") {
          body["repoBranch"] = "main";
          body["pullRequestPreviewsEnabled"] = "no";
        }

        await fetcher("/api/projects", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(body),
        }).then(async () => {
          toast.success(
            "You have successfully created a new bot project",
            toastStyle
          );
        });

        // refresh projects lists
        mutate();
      } catch (e: any) {
        toast.error(e.message);
      } finally {
        setIsLoading(false);
      }
    },
    [mutate]
  );

  return (
    <form onSubmit={onSubmit}>
      <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-2 p-6">
        <div className="max-w-md">
          <label
            htmlFor="platform"
            className="block text-gray-500 text-sm font-semibold"
          >
            Platform
          </label>
          <div className="pt-2 mb-6">
            <Listbox
              value={platformSelected}
              onChange={setPlatformSelected}
              name="platform"
            >
              {({ open }) => (
                <>
                  <div className="relative">
                    <Listbox.Button className="relative w-full cursor-pointer rounded-md border border-gray-800 bg-secondary py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                      <span className="flex items-center">
                        <img
                          src={`https://cdn-botway.deno.dev/icons/${platformSelected.slug}.svg`}
                          alt={`${platformSelected.slug} icon`}
                          className="h-6 w-6 flex-shrink-0"
                        />
                        <span className="ml-3 block truncate">
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
                      <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                        {platforms.map((platform) => (
                          <Listbox.Option
                            key={platform.name}
                            onChange={() => setPlatformSelected(platform)}
                            className={({ active }) =>
                              clsx(
                                active
                                  ? "text-white bg-secondary"
                                  : "text-gray-500",
                                "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-md mx-2 my-1"
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
                                      "ml-3 block truncate"
                                    )}
                                  >
                                    {platform.name}
                                  </span>
                                </div>

                                {selected ? (
                                  <span
                                    className={clsx(
                                      active ? "text-white" : "text-blue-700",
                                      "absolute inset-y-0 right-0 flex items-center pr-4"
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
        </div>
        <div className="max-w-md">
          <label
            htmlFor="language"
            className="block text-gray-500 text-sm font-semibold"
          >
            Programming Language
          </label>
          <div className="pt-2 mb-6">
            <Listbox value={langSelected} onChange={setLangSelected}>
              {({ open }) => (
                <>
                  <div className="relative">
                    <Listbox.Button className="relative w-full cursor-pointer rounded-md border border-gray-800 bg-secondary py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                      <span className="flex items-center">
                        <img
                          src={`https://cdn-botway.deno.dev/icons/${langSelected.slug}.svg`}
                          alt={`${langSelected.slug} icon`}
                          className="h-6 w-6 flex-shrink-0"
                        />
                        <span className="ml-3 block truncate">
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
                      <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                        {langs(platformSelected.name).map((lang) => (
                          <Listbox.Option
                            key={lang.name}
                            className={({ active }) =>
                              clsx(
                                active
                                  ? "text-white bg-secondary"
                                  : "text-gray-500",
                                "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-md mx-2 my-1"
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
                                      "ml-3 block truncate"
                                    )}
                                  >
                                    {lang.name}
                                  </span>
                                </div>

                                {selected ? (
                                  <span
                                    className={clsx(
                                      active ? "text-white" : "text-blue-700",
                                      "absolute inset-y-0 right-0 flex items-center pr-4"
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
        </div>

        <div className="max-w-md pt-4 sm:pt-4">
          <label
            htmlFor="package-manager"
            className="block text-gray-500 text-sm font-semibold"
          >
            Package Manager
          </label>
          <div className="pt-2 mb-4">
            <Listbox
              value={pmSelected}
              refName={packageManagerRef}
              onChange={setPMSelected}
            >
              {({ open }) => (
                <>
                  <div className="relative">
                    <Listbox.Button className="relative w-full cursor-pointer rounded-md border border-gray-800 bg-secondary py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                      <span className="flex items-center">
                        <img
                          src={`https://cdn-botway.deno.dev/icons/${pmSelected.logo}`}
                          alt={`${pmSelected.logo} icon`}
                          className="h-6 w-6 flex-shrink-0"
                        />
                        <span className="ml-3 block truncate">
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
                      <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                        {packageManagers(langSelected.name).map((pm) => (
                          <Listbox.Option
                            key={pm.name}
                            className={({ active }) =>
                              clsx(
                                active
                                  ? "text-white bg-secondary"
                                  : "text-gray-500",
                                "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-md mx-2 my-1"
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
                                      "ml-3 block truncate"
                                    )}
                                  >
                                    {pm.name}
                                  </span>
                                </div>

                                {selected ? (
                                  <span
                                    className={clsx(
                                      active ? "text-white" : "text-blue-700",
                                      "absolute inset-y-0 right-0 flex items-center pr-4"
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
        </div>

        <div className="max-w-md pt-4 sm:pt-4">
          <label
            htmlFor="host-service"
            className="block text-gray-500 text-sm font-semibold"
          >
            Host Service
          </label>
          <div className="pt-2 mb-4">
            <Listbox
              value={hostServiceSelected}
              refName={hostServiceRef}
              onChange={setHostServiceSelected}
            >
              {({ open }) => (
                <>
                  <div className="relative">
                    <Listbox.Button className="relative w-full cursor-pointer rounded-md border border-gray-800 bg-secondary py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                      <span className="flex items-center">
                        <img
                          src={`https://cdn-botway.deno.dev/icons/${hostServiceSelected.logo}`}
                          alt={`${hostServiceSelected.logo} icon`}
                          className="h-6 w-6 flex-shrink-0"
                        />
                        <span className="ml-3 block truncate">
                          {hostServiceSelected.name}
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
                      <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                        {hostServices.map((hostService) => (
                          <Listbox.Option
                            key={hostService.name}
                            className={({ active }) =>
                              clsx(
                                active
                                  ? "text-white bg-secondary"
                                  : "text-gray-500",
                                "relative transition cursor-pointer select-none py-2 pl-3 pr-9 rounded-md mx-2 my-1"
                              )
                            }
                            value={hostService}
                          >
                            {({ selected, active }) => (
                              <>
                                <div className="flex items-center">
                                  <img
                                    src={`https://cdn-botway.deno.dev/icons/${hostService.logo}`}
                                    alt={`${hostService.logo} icon`}
                                    className="h-6 w-6 flex-shrink-0"
                                    width={16}
                                  />
                                  <span
                                    className={clsx(
                                      selected
                                        ? "font-semibold"
                                        : "font-normal",
                                      "ml-3 block truncate"
                                    )}
                                  >
                                    {hostService.name}
                                  </span>
                                </div>

                                {selected ? (
                                  <span
                                    className={clsx(
                                      active ? "text-white" : "text-blue-700",
                                      "absolute inset-y-0 right-0 flex items-center pr-4"
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
              name="hostService[name]"
              value={hostServiceSelected.name.toLowerCase()}
              ref={hostServiceRef}
            />
          </div>
        </div>

        <div className="max-w-md pt-4 sm:pt-4">
          <label
            htmlFor="visibility"
            className="block text-gray-500 text-sm font-semibold"
          >
            Visibility On GitHub
          </label>
          <div className="pt-2 mb-4">
            <Listbox
              value={visibilitySelected}
              refName={visibilityRef}
              onChange={setvisibilitySelected}
            >
              {({ open }) => (
                <>
                  <div className="relative">
                    <Listbox.Button className="relative w-full cursor-pointer rounded-md border border-gray-800 bg-secondary py-2 pl-3 pr-10 text-left shadow-sm outline-none sm:text-sm">
                      <span className="flex items-center">
                        <span className="ml-2 block truncate">
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
                      <Listbox.Options className="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                        {visibilityOptions.map((visibility) => (
                          <Listbox.Option
                            key={visibility.typeName}
                            className={({ active }) =>
                              clsx(
                                active
                                  ? "text-white bg-secondary"
                                  : "text-gray-500",
                                "relative transition cursor-pointer select-none py-2 pl-2 pr-9 rounded-md mx-2 my-1"
                              )
                            }
                            value={visibility}
                          >
                            {({ selected, active }) => (
                              <>
                                <div className="flex items-center">
                                  <span
                                    className={clsx(
                                      selected
                                        ? "font-semibold"
                                        : "font-normal",
                                      "ml-3 block truncate"
                                    )}
                                  >
                                    {visibility.typeName}
                                  </span>
                                </div>

                                {selected ? (
                                  <span
                                    className={clsx(
                                      active ? "text-white" : "text-blue-700",
                                      "absolute inset-y-0 right-0 flex items-center pr-4"
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
              name="visibility[typeName]"
              value={visibilitySelected.type}
              ref={visibilityRef}
            />
          </div>
        </div>

        <div className="max-w-md pt-4 sm:pt-4">
          <label
            htmlFor="name"
            className="block text-gray-500 text-sm font-semibold"
          >
            Bot Name
          </label>
          <div className="pt-2 mb-8">
            <input
              className="w-full border px-1.5 py-sm bg-secondary trsn bg border-gray-800 text-gray-400 sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
              type="text"
              ref={nameRef}
              placeholder="bot name"
            />
          </div>
        </div>
      </div>
      <div className="border-t border-gray-800">
        <Button htmlType="submit" type="success" loading={isLoading}>
          Create Bot Project
        </Button>
      </div>
    </form>
  );
};

export const NewProject = () => {
  return (
    <NewProjectModal>
      <NewProjectHandler />
    </NewProjectModal>
  );
};
