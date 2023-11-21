"use client";

import { Button } from "@/components/Button";
import { toastStyle } from "@/tools/toast-style";
import toast from "react-hot-toast";
import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import supabase from "@/supabase/browser";
import clsx from "clsx";
import Template from "./tmp";
import { VIEWS, useAuth } from "../provider";

const SignUpSchema = Yup.object().shape({
  email: Yup.string().email("Invalid email").required("Required"),
  password: Yup.string().required("Required"),
  name: Yup.string().required("Required"),
});

const SignUp = () => {
  const { setView } = useAuth();

  const signUp = async (formData: any) => {
    const signUpFunc = async (
      params: any,
    ): Promise<{ auth: any; error: Error | null }> => {
      const { data, error } = await supabase.auth.signUp({
        email: params.email,
        password: params.password,
        options: {
          data: {
            name: params.name,
            githubApiToken: "",
            zeaburApiToken: "",
          },
        },
      });

      let authError = null;

      // User exists, but is fake. See https://supabase.com/docs/reference/javascript/auth-signup
      if (
        data.user &&
        data.user.identities &&
        data.user.identities.length === 0
      ) {
        authError = {
          name: "AuthApiError",
          message: "User already exists",
        };
      } else if (error)
        authError = {
          name: error.name,
          message: error.message,
        };

      return { auth: data, error: authError };
    };

    const { error } = await signUpFunc({
      email: formData.email,
      password: formData.password,
      name: formData.name,
    });

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    } else {
      toast.success(
        "Your account has been created\nPlease check your email for further instructions",
        toastStyle,
      );
    }
  };

  return (
    <Template>
      <h2 className="text-xl font-farray font-medium md:text-2xl text-white">
        Create a new Account
      </h2>

      <p className="text-sm text-gray-500 font-medium pt-1 cursor-pointer">
        Already have an account?{" "}
        <a
          onClick={() => setView(VIEWS.SIGN_IN)}
          className="text-blue-700 font-mono"
        >
          Sign in
        </a>
      </p>

      <div className="my-2 mb-2 pt-8">
        <Formik
          initialValues={{
            email: "",
            password: "",
            name: "",
          }}
          validationSchema={SignUpSchema}
          onSubmit={signUp}
        >
          {({ errors, touched }) => (
            <Form className="column w-full">
              <label
                htmlFor="name"
                className="block text-gray-500 pb-2 text-sm font-medium"
              >
                Name
              </label>

              <Field
                className="input"
                id="name"
                name="name"
                placeholder="Your Name"
                type="text"
              />

              {errors.name && touched.name ? (
                <div className="text-red-600 text-sm font-medium pt-1">
                  {errors.name}
                </div>
              ) : null}

              <div className="pb-6" />

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
                Create
              </Button>
            </Form>
          )}
        </Formik>
      </div>
    </Template>
  );
};

export default SignUp;
