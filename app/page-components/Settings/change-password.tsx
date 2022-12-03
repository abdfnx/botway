import { useCallback, useEffect, useRef, useState } from "react";
import { fetcher } from "@/lib/fetch";
import toast from "react-hot-toast";
import { bgSecondary } from "@/tools/colors";
import { SettingsLayout } from ".";
import { Button } from "@/components/Button";
import { useCurrentUser } from "@/lib/user";
import { useRouter } from "next/router";
import { LoadingDots } from "@/components/LoadingDots";
import Layout from "@/components/Layout";

export const ChangePassword = () => {
  const oldPasswordRef: any = useRef();
  const newPasswordRef: any = useRef();
  const [isLoading, setIsLoading] = useState(false);
  const { data, error } = useCurrentUser();
  const loading = !data && !error;
  const router = useRouter();

  const onSubmit = useCallback(async (e: any) => {
    e.preventDefault();

    try {
      setIsLoading(true);

      await fetcher("/api/user/password", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          oldPassword: oldPasswordRef.current.value,
          newPassword: newPasswordRef.current.value,
        }),
      });

      toast.success("Your password has been updated", {
        style: {
          borderRadius: "10px",
          backgroundColor: bgSecondary,
          color: "#fff",
        },
      });
    } catch (e: any) {
      toast.error(e.message, {
        style: {
          borderRadius: "10px",
          backgroundColor: bgSecondary,
          color: "#fff",
        },
      });
    } finally {
      setIsLoading(false);

      oldPasswordRef.current.value = "";
      newPasswordRef.current.value = "";
    }
  }, []);

  const PushToSignIn = () => {
    useEffect(() => {
      if (!data?.user) {
        router.push("/sign-in");
      }
    }, []);

    return <></>;
  };

  return (
    <>
      {loading ? (
        <LoadingDots>Loading</LoadingDots>
      ) : data?.user ? (
        <Layout title="General Settings">
          <span className="flex items-center">
            <SettingsLayout user={data.user}>
              <div className="flex-1">
                <div className="mb-4">
                  <h1 className="text-xl text-gray-400 font-semibold">
                    Change Password
                  </h1>
                </div>
                <form onSubmit={onSubmit}>
                  <input type="hidden" name="_method" value="PATCH" />

                  <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-1">
                    <div className="max-w-md">
                      <label
                        htmlFor="old-password"
                        className="block text-gray-500 text-sm font-semibold"
                      >
                        Old Password
                      </label>
                      <div className="pt-2">
                        <input
                          className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                          ref={oldPasswordRef}
                          autoComplete="current-password"
                          placeholder="••••••••••••••••"
                          type="password"
                        />

                        <p className="lg:hidden mb-3" />
                      </div>
                    </div>
                    <div className="max-w-md lg:pl-6">
                      <label
                        htmlFor="new-password"
                        className="block text-gray-500 text-sm font-semibold"
                      >
                        New Password
                      </label>
                      <div className="pt-2 mb-6">
                        <input
                          className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                          ref={newPasswordRef}
                          autoComplete="new-password"
                          placeholder="••••••••••••••••"
                          type="password"
                        />
                      </div>
                    </div>
                  </div>
                  <div className="border-t border-gray-800">
                    <Button
                      htmlType="submit"
                      type="success"
                      loading={isLoading}
                    >
                      Save Changes
                    </Button>
                  </div>
                </form>
              </div>
            </SettingsLayout>
          </span>
        </Layout>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};
