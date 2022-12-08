import { Dialog, Transition } from "@headlessui/react";
import { ZapIcon } from "@primer/octicons-react";
import { Fragment, useState } from "react";

export const NewProjectModal = ({ children }: any) => {
  const [isOpen, setIsOpen] = useState(false);

  const closeModal = () => {
    setIsOpen(false);
  };

  const openModal = () => {
    setIsOpen(true);
  };

  return (
    <>
      <button
        className="h-9 px-2 py-3.5 rounded-md border border-gray-800 inline-flex flex-shrink-0 whitespace-nowrap items-center gap-2 transition-colors duration-200 ease-in-out leading-none border-1 cursor-pointer text-white hover:color-primary hover:bg-secondary"
        onClick={openModal}
        aria-current="page"
      >
        <ZapIcon size={16} className="fill-blue-700" /> New Project
      </button>

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
                <Dialog.Panel className="w-full max-w-5xl max-h-full border-2 border-dashed border-gray-800 transform overflow-hidden rounded-2xl bg p-6 text-left align-middle shadow-xl transition-all">
                  {children}
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition>
    </>
  );
};
