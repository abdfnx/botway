import { Button } from "@/components/Button";
import Layout from "@/components/Layout";
import { LoadingDots } from "@/components/LoadingDots";
import { fetcher } from "@/lib/fetch";
import { useCurrentUser } from "@/lib/user";
import { useRouter } from "next/router";
import { useCallback, useEffect, useRef, useState } from "react";
import toast from "react-hot-toast";

export const AccountInfo = ({ user, mutate }: any) => {
  const usernameRef: any = useRef();
  const nameRef: any = useRef();
  const [isLoading, setIsLoading] = useState(false);

  const onSubmit = useCallback(
    async (e: any) => {
      e.preventDefault();

      try {
        setIsLoading(true);

        const formData = new FormData();

        formData.append("username", usernameRef.current.value);
        formData.append("name", nameRef.current.value);

        const response = await fetcher("/api/user", {
          method: "PATCH",
          body: formData,
        });

        mutate({ user: response.user }, false);

        toast.success("Your profile has been updated", {
          style: {
            borderRadius: "10px",
          },
          className: "bg-secondary",
        });
      } catch (e: any) {
        toast.error(e.message, {
          style: {
            borderRadius: "10px",
          },
          className: "bg-secondary",
        });
      } finally {
        setIsLoading(false);
      }
    },
    [mutate]
  );

  useEffect(() => {
    usernameRef.current.value = user.username;
    nameRef.current.value = user.name;
  }, [user]);

  return (
    <>
      <div className="my-2 mb-2">
        <form onSubmit={onSubmit}>
          <span className="text-gray-400 text-lg">Account Information</span>

          <div className="grid gap-4 grid-cols-2 grid-rows-2">
            <div className="pt-4">
              <label
                htmlFor="email"
                className="block text-gray-500 text-sm font-semibold"
              >
                Email
              </label>
              <div className="pt-2">
                <input
                  className="trsn bg border border-gray-800 placeholder:text-gray-400 cursor-not-allowed text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                  aria-label="Email Address"
                  type="email"
                  value={user.email}
                  disabled
                />
              </div>
            </div>
            <div className="pt-4 lg:pl-4">
              <label
                htmlFor="name"
                className="block text-gray-500 text-sm font-semibold"
              >
                Name
              </label>
              <div className="pt-2">
                <input
                  className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                  ref={nameRef}
                  placeholder="name"
                  type="text"
                />
              </div>
            </div>

            <div className="pt-4">
              <label
                htmlFor="username"
                className="block text-gray-500 text-sm font-semibold"
              >
                Username
              </label>
              <div className="pt-2">
                <input
                  className="trsn bg border border-gray-800 placeholder:text-gray-400 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block w-full p-2"
                  ref={usernameRef}
                  placeholder="username"
                  type="text"
                />
              </div>
            </div>
          </div>

          <Button htmlType="submit" type="success" loading={isLoading}>
            Save Changes
          </Button>
        </form>
      </div>
    </>
  );
};

export const Settings = () => {
  const { data, error, mutate } = useCurrentUser();
  const loading = !data && !error;
  const router = useRouter();

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
            {/* <AboutYou user={data.user} mutate={mutate} /> */}
            <AccountInfo user={data.user} mutate={mutate} />
          </span>
        </Layout>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};
