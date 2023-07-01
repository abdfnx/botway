"use client";

import { LoadingDots } from "@/components/LoadingDots";
import { useAuth, VIEWS } from "../provider";
import ResetPassword from "./reset";
import SignIn from "./sign-in";
import SignUp from "./sign-up";
import UpdatePassword from "./update";

const Auth = ({ view: initialView }: any) => {
  let { view } = useAuth();

  if (initialView) {
    view = initialView;
  }

  switch (view) {
    case VIEWS.UPDATE_PASSWORD:
      return <UpdatePassword />;

    case VIEWS.FORGOTTEN_PASSWORD:
      return <ResetPassword />;

    case VIEWS.SIGN_UP:
      return <SignUp />;

    case VIEWS.SIGN_IN:
      return <SignIn />;

    default:
      return (
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      );
  }
};

export default Auth;
