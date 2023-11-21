"use client";

import clsx from "clsx";
import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import { useAuth, VIEWS } from "../provider";
import supabase from "@/supabase/browser";
import Template from "./tmp";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import { Button } from "@/components/Button";

const SignInSchema = Yup.object().shape({
  email: Yup.string().email("Invalid email").required("Required"),
  password: Yup.string().required("Required"),
});

const SignIn = () => {
  const { setView } = useAuth();

  const signIn = async (formData: any) => {
    const { error } = await supabase.auth.signInWithPassword({
      email: formData.email,
      password: formData.password,
    });

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    }
  };

  return (
    <Template>
      <h2 className="text-xl font-farray font-medium md:text-2xl text-white">
        Welcome to Botway
      </h2>

      <p className="text-[13px] md:text-sm text-gray-500 font-medium pt-2 cursor-pointer">
        Don't have an Account?{" "}
        <a
          onClick={() => setView(VIEWS.SIGN_UP)}
          className="text-blue-700 font-mono"
        >
          Sign up for an account
        </a>
      </p>

      <p className="text-[13px] md:text-sm text-gray-500 font-medium pt-1.5 cursor-pointer">
        Forget Password?{" "}
        <a
          onClick={() => setView(VIEWS.FORGOTTEN_PASSWORD)}
          className="text-blue-700 font-mono"
        >
          Reset
        </a>
      </p>

      <div className="my-2 mb-2 pt-8">
        <Formik
          initialValues={{
            email: "",
            password: "",
          }}
          validationSchema={SignInSchema}
          onSubmit={signIn}
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
                className={clsx(
                  "input",
                  errors.email && touched.email && "bg-red-50",
                )}
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

              <div className="pb-6" />

              <label
                htmlFor="password"
                className="block text-gray-500 pb-2 text-sm font-medium"
              >
                Password
              </label>

              <Field
                className={clsx(
                  "input",
                  errors.password && touched.password && "bg-red-50",
                )}
                id="password"
                placeholder="••••••••••••••••"
                name="password"
                type="password"
              />

              {errors.password && touched.password ? (
                <div className="text-red-600 text-sm font-medium pt-1">
                  {errors.password}
                </div>
              ) : null}

              <Button htmlType="submit" className="button w-full p-2">
                Sign in
              </Button>
            </Form>
          )}
        </Formik>
      </div>
    </Template>
  );
};

export default SignIn;
