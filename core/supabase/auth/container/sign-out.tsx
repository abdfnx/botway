"use client";

import { useAuth } from "@/supabase/auth/provider";
import { SignOutIcon } from "@primer/octicons-react";

export const SignOut = () => {
  const { signOut } = useAuth();

  return (
    <div className="border-b py-5 px-6 border-gray-800">
      <ul className="space-y-1">
        <li role="menuitem" className="outline-none">
          <a
            className="cursor-pointer flex space-x-3 items-center outline-none focus-visible:ring-1 focus-visible:z-10 group py-1 font-normal border-gray-800"
            style={{ marginLeft: "0rem" }}
            onClick={signOut}
          >
            <div className="transition truncate text-sm min-w-fit">
              <SignOutIcon className="fill-red-600" />
            </div>
            <span className="transition truncate text-gray-400 hover:text-white text-sm w-full">
              Sign Out
            </span>
          </a>
        </li>
      </ul>
    </div>
  );
};
