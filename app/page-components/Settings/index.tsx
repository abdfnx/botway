import { Button } from "@/components/Button";
import Layout from "@/components/Layout";
import { LoadingDots } from "@/components/LoadingDots";
import { fetcher } from "@/lib/fetch";
import { useCurrentUser } from "@/lib/user";
import { toastStyle } from "@/tools/toast-style";
import { useRouter } from "next/router";
import { useCallback, useEffect, useRef, useState } from "react";
import toast from "react-hot-toast";
import { EncryptJWT, jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/api-tokens";

const AccountInfo = ({ user, mutate }: any) => {
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

        toast.success("Your profile has been updated", toastStyle);
      } catch (e: any) {
        toast.error(e.message, toastStyle);
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
      <div className="flex-1 border-b border-gray-800 shadow-sm p-5 mb-8">
        <div className="mb-4">
          <h1 className="text-xl text-gray-400 font-semibold">
            Account Information
          </h1>
        </div>
        <form onSubmit={onSubmit} className="w-full">
          <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-2">
            <div className="max-w-md">
              <label
                htmlFor="email"
                className="block text-gray-500 text-sm font-semibold"
              >
                Email
              </label>
              <div className="pt-2">
                <input
                  className="w-full border px-1.5 py-sm bg-secondary trsn bg border-gray-800 text-gray-400 cursor-not-allowed sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                  aria-label="Email Address"
                  type="email"
                  value={user.email}
                  disabled
                />

                <p className="mb-3 py-2 text-sm text-gray-400">
                  You cannot change your email
                </p>
              </div>
            </div>
            <div className="max-w-md lg:pl-6">
              <label
                htmlFor="name"
                className="block text-gray-500 text-sm font-semibold"
              >
                Name
              </label>
              <div className="pt-2">
                <input
                  className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                  ref={nameRef}
                  placeholder="name"
                  type="text"
                />
              </div>
            </div>
            <div className="max-w-md pt-4 sm:pt-4">
              <label
                htmlFor="username"
                className="block text-gray-500 text-sm font-semibold"
              >
                Username
              </label>
              <div className="pt-2 mb-6">
                <input
                  className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                  ref={usernameRef}
                  placeholder="username"
                  type="text"
                />
              </div>
            </div>
          </div>
          <div className="mb-5">
            <Button htmlType="submit" type="success" loading={isLoading}>
              Update Information
            </Button>
          </div>
        </form>
      </div>
    </>
  );
};

const ChangePassword = () => {
  const currentPasswordRef: any = useRef();
  const newPasswordRef: any = useRef();
  const [isLoading, setIsLoading] = useState(false);

  const onSubmit = useCallback(async (e: any) => {
    e.preventDefault();

    try {
      setIsLoading(true);

      await fetcher("/api/user/password", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          currentPassword: currentPasswordRef.current.value,
          newPassword: newPasswordRef.current.value,
        }),
      });

      toast.success("Your password has been updated", toastStyle);
    } catch (e: any) {
      toast.error(e.message, toastStyle);
    } finally {
      setIsLoading(false);

      currentPasswordRef.current.value = "";
      newPasswordRef.current.value = "";
    }
  }, []);

  return (
    <div className="flex-1 border-b border-gray-800 shadow-sm p-5 mb-8">
      <div className="mb-4">
        <h1 className="text-xl text-gray-400 font-semibold">Change Password</h1>
      </div>
      <form onSubmit={onSubmit} className="w-full">
        <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-1">
          <div className="max-w-md">
            <label
              htmlFor="current-password"
              className="block text-gray-500 text-sm font-semibold"
            >
              Current Password
            </label>
            <div className="pt-2">
              <input
                className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                ref={currentPasswordRef}
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
        <div className="mb-5">
          <Button htmlType="submit" type="success" loading={isLoading}>
            Update Password
          </Button>
        </div>
      </form>
    </div>
  );
};

const Tokens = ({ user, mutate }: any) => {
  const githubApiTokenRef: any = useRef();
  const railwayApiTokenRef: any = useRef();
  const renderApiTokenRef: any = useRef();
  const [isLoading, setIsLoading] = useState(false);

  const onSubmit = useCallback(
    async (e: any) => {
      e.preventDefault();

      try {
        setIsLoading(true);

        const formData = new FormData();

        const githubApiToken = await new EncryptJWT({
          data: githubApiTokenRef.current.value,
        })
          .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
          .encrypt(BW_SECRET_KEY);
        const railwayApiToken = await new EncryptJWT({
          data: railwayApiTokenRef.current.value,
        })
          .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
          .encrypt(BW_SECRET_KEY);
        const renderApiToken = await new EncryptJWT({
          data: renderApiTokenRef.current.value,
        })
          .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
          .encrypt(BW_SECRET_KEY);

        formData.append("githubApiToken", githubApiToken);
        formData.append("railwayApiToken", railwayApiToken);
        formData.append("renderApiToken", renderApiToken);

        const response = await fetcher("/api/user", {
          method: "PATCH",
          body: formData,
        });

        mutate({ user: response.user }, false);

        toast.success("Your account tokens has been updated", toastStyle);
      } catch (e: any) {
        toast.error(e.message, toastStyle);
      } finally {
        setIsLoading(false);
      }
    },
    [mutate]
  );

  useEffect(() => {
    const set = async () => {
      let ghApiToken: any, railwayApiToken: any, renderApiToken: any;

      if (user.githubApiToken) {
        const { payload } = await jwtDecrypt(
          user.githubApiToken,
          BW_SECRET_KEY
        );

        ghApiToken = payload.data;
      }

      if (user.railwayApiToken) {
        const { payload } = await jwtDecrypt(
          user.railwayApiToken,
          BW_SECRET_KEY
        );

        railwayApiToken = payload.data;
      }

      if (user.renderApiToken) {
        const { payload } = await jwtDecrypt(
          user.renderApiToken,
          BW_SECRET_KEY
        );

        renderApiToken = payload.data;
      }

      if (user.githubApiToken != "")
        githubApiTokenRef.current.value = ghApiToken;
      else githubApiTokenRef.current.value = "";

      if (user.railwayApiToken != "")
        railwayApiTokenRef.current.value = railwayApiToken;
      else railwayApiTokenRef.current.value = "";

      if (user.renderApiToken != "")
        renderApiTokenRef.current.value = renderApiToken;
      else renderApiTokenRef.current.value = "";
    };

    set().catch(console.error);
  }, [user]);

  return (
    <>
      <div className="flex-1 shadow-sm p-5 mb-8">
        <div className="mb-4">
          <h1 className="text-xl text-gray-400 font-semibold">API Tokens</h1>
        </div>
        <form onSubmit={onSubmit} className="w-full">
          <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-2">
            <div className="max-w-md">
              <label
                htmlFor="github-api-token"
                className="block text-gray-500 text-sm font-semibold"
              >
                GitHub API Token
              </label>
              <div className="pt-2 mb-6">
                <input
                  className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                  ref={githubApiTokenRef}
                  placeholder="github api token"
                  type="password"
                />
              </div>
            </div>
            <div className="max-w-md lg:pl-6">
              <label
                htmlFor="railway-api-token"
                className="block text-gray-500 text-sm font-semibold"
              >
                Railway API Token
              </label>
              <div className="pt-2 mb-6">
                <input
                  className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                  ref={railwayApiTokenRef}
                  placeholder="railway api token"
                  type="password"
                />
              </div>
            </div>
            <div className="max-w-md">
              <label
                htmlFor="render-api-token"
                className="block text-gray-500 text-sm font-semibold"
              >
                Render API Token
              </label>
              <div className="pt-2 mb-6">
                <input
                  className="w-full border px-1.5 bg-secondary trsn bg border-gray-800 placeholder:text-gray-400 placeholder:pl-1 text-white sm:text-sm rounded-lg focus:outline-none hover:border-blue-700 block p-2"
                  ref={renderApiTokenRef}
                  placeholder="render api token"
                  type="password"
                />
              </div>
            </div>
          </div>
          <div className="mb-5">
            <Button htmlType="submit" type="success" loading={isLoading}>
              Update Configuration
            </Button>
          </div>
        </form>
      </div>
    </>
  );
};

const EmailIsNotVerifed = () => {
  return (
    <div className="flex-1 border rounded-xl shadow-sm p-5 mb-8 bg-orange-25 border-orange-600">
      <div className="w-full">
        <div className="lg:grid lg:gap-2 lg:grid-cols-2 lg:grid-rows-1">
          <div className="max-w-md">
            <span className="text-yellow-500">Your Email is Not Verfied</span>
          </div>
        </div>
      </div>
    </div>
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
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : data?.user ? (
        <Layout title="General Settings">
          {!data.user.emailVerified &&
          process.env.NEXT_PUBLIC_FULL == "true" ? (
            <span className="flex items-center">
              <EmailIsNotVerifed />
            </span>
          ) : (
            <></>
          )}
          <span className="flex items-center">
            <AccountInfo user={data.user} mutate={mutate} />
          </span>
          <span className="flex items-center">
            <ChangePassword />
          </span>
          <span className="flex items-center">
            <Tokens user={data.user} mutate={mutate} />
          </span>
        </Layout>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};
