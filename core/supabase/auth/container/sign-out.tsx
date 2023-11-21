"use client";

import { useAuth } from "@/supabase/auth/provider";
import { SignOutIcon } from "@primer/octicons-react";

export const SignOut = () => {
  const { signOut } = useAuth();

  return (
    <div
      onClick={signOut}
      className="flex text-red-600 justify-between items-center"
    >
      <SignOutIcon size={24} className="mr-2" />
      Sign Out
    </div>
  );
};
