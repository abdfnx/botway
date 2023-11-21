"use client";

import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import { useAuth, VIEWS } from "@/supabase/auth/provider";
import supabase from "@/supabase/browser";
import clsx from "clsx";
import { toastStyle } from "@/tools/toast-style";
import { toast } from "react-hot-toast";
import Template from "./tmp";
import { Button } from "@/components/Button";

const ResetPasswordSchema = Yup.object().shape({
  email: Yup.string().email("Invalid email").required("Required"),
});

const ResetPassword = () => {
  const { setView } = useAuth();

  const origin =
    typeof window !== "undefined" && window.location.origin
      ? window.location.origin
      : "";

  const resetPassword = async (formData: any) => {
    const { error } = await supabase.auth.resetPasswordForEmail(
      formData?.email,
      {
        redirectTo: origin,
      },
    );

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    } else {
      toast.success("Password reset instructions sent", toastStyle);
    }
  };

  return (
    <Template>
      <h2 className="text-xl font-farray font-medium md:text-2xl text-white">
        Reset Password
      </h2>

      <p className="text-sm font-medium pt-1 cursor-pointer">
        <a
          onClick={() => setView(VIEWS.SIGN_IN)}
          className="text-blue-700 font-mono"
        >
          Return to Sign In page
        </a>
      </p>

      <div className="my-2 mb-2 pt-8">
        <Formik
          initialValues={{
            email: "",
          }}
          validationSchema={ResetPasswordSchema}
          onSubmit={resetPassword}
        >
          {({ errors, touched }) => (
            <Form className="column w-full">
              <label
                htmlFor="email"
                className="block text-gray-500 pb-2 text-sm font-medium"
              >
                Email
              </label>

              <Field
                className={clsx("input", errors.email && "bg-red-50")}
                id="email"
                name="email"
                placeholder="Email Address"
                type="email"
              />

              {errors.email && touched.email ? (
                <div className="text-red-600 text-sm font-medium pt-1">
                  {errors.email}
                </div>
              ) : null}

              <Button htmlType="submit" className="button w-full p-2">
                Send Instructions
              </Button>
            </Form>
          )}
        </Formik>
      </div>
    </Template>
  );
};

export default ResetPassword;
