import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useState } from "react";
import { ProjectMain } from "./project";

export const Project = ({ project, mutate, user }: any) => {
  const [isOpen, setIsOpen] = useState(false);

  const closeModal = () => {
    mutate();
    setIsOpen(false);
  };

  const openModal = () => {
    setIsOpen(true);
  };

  return (
    <>
      <div className="flex items-center justify-between gap-3 px-5 py-0 rounded-lg border-2 border-dashed border-gray-800 hover:bg-secondary transition-colors duration-200">
        <div className="block flex-1 py-5 cursor-pointer" onClick={openModal}>
          <h2>
            <strong className="text-base text-white leading-tight font-medium align-middle">
              {project.name}
              {project.icon != "" ?? (
                <img src={project.icon} alt="project icon" width={16} />
              )}
            </strong>
          </h2>
          <br />
          <p className="flex items-center gap-1.5 mt-1.5 text-sm text-gray-500">
            <img
              src={`https://cdn-botway.deno.dev/icons/${project.platform}.svg`}
              alt={`${project.platform} icon`}
              width={16}
            />
            {project.platform}

            <img
              src={`https://cdn-botway.deno.dev/icons/${project.lang}.svg`}
              alt={`${project.lang} icon`}
              width={16}
            />
            {project.lang}
          </p>
        </div>
      </div>

      <Transition.Root show={isOpen} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={closeModal}>
          <Transition.Child
            as={Fragment}
            enter="ease-in-out duration-500"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in-out duration-500"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0" />
          </Transition.Child>

          <div className="fixed inset-0 overflow-hidden max-w-full">
            <div className="absolute inset-0 overflow-hidden">
              <div className="pointer-events-none fixed inset-y-0 right-0 flex max-w-full max-h-full pl-16">
                <Transition.Child
                  as={Fragment}
                  enter="transform transition ease-in-out duration-500 sm:duration-700"
                  enterFrom="translate-x-full"
                  enterTo="translate-x-0"
                  leave="transform transition ease-in-out duration-500 sm:duration-700"
                  leaveFrom="translate-x-0"
                  leaveTo="translate-x-full"
                >
                  <Dialog.Panel className="pointer-events-auto pt-16 relative w-screen max-w-full lg:pl-32">
                    <Transition.Child
                      as={Fragment}
                      enter="ease-in-out duration-500"
                      enterFrom="opacity-0"
                      enterTo="opacity-100"
                      leave="ease-in-out duration-500"
                      leaveFrom="opacity-100"
                      leaveTo="opacity-0"
                    >
                      <div className="absolute top-0 left-0 -ml-20 flex pt-4 pr-2 sm:-ml-20 sm:pr-4" />
                    </Transition.Child>
                    <div className="flex h-full w-full flex-col overflow-y-scroll bg py-6 shadow-xl rounded-tl-xl border-l-2 border-t-2 border-dashed border-l-gray-800 border-t-gray-800">
                      <div className="px-4 sm:px-6">
                        <Dialog.Title className="text-lg font-medium text-white">
                          {project.name}
                        </Dialog.Title>
                      </div>
                      <div className="relative mt-6 flex-1 px-4 sm:px-6">
                        <ProjectMain
                          project={project}
                          mutate={mutate}
                          user={user}
                        />
                      </div>
                    </div>
                  </Dialog.Panel>
                </Transition.Child>
              </div>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  );
};
