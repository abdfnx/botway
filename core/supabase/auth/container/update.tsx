"use client";

import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import supabase from "@/supabase/browser";
import clsx from "clsx";
import { toast } from "react-hot-toast";
import { toastStyle } from "@/tools/toast-style";
import Template from "./tmp";
import { Button } from "@/components/Button";

const UpdatePasswordSchema = Yup.object().shape({
  password: Yup.string().required("Required"),
});

const UpdatePassword = () => {
  const updatePassword = async (formData: any) => {
    const { error } = await supabase.auth.updateUser({
      password: formData.password,
    });

    if (error) {
      toast.error(error.message, toastStyle);

      console.log(error);
    }
  };

  return (
    <Template>
      <h2 className="text-lg font-medium md:text-2xl text-white">
        Reset Password
      </h2>

      <p className="text-sm text-gray-500 font-medium pt-1 cursor-pointer">
        Enter a new password for your account
      </p>

      <div className="my-2 mb-2 pt-8">
        <Formik
          initialValues={{
            password: "",
          }}
          validationSchema={UpdatePasswordSchema}
          onSubmit={updatePassword}
        >
          {({ errors, touched }) => (
            <Form className="column w-full">
              <label
                htmlFor="password"
                className="block text-gray-500 pb-2 text-sm font-medium"
              >
                New Password
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
                Update Password
              </Button>
            </Form>
          )}
        </Formik>
      </div>
    </Template>
  );
};

export default UpdatePassword;
