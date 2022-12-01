import { LoadingDots } from "@/components/LoadingDots";
import { useCurrentUser } from "@/lib/user";
import Layout from "@/components/Layout";
import { UserAvatar } from "../../components/UserAvatar";
import { useRouter } from "next/router";
import { useEffect } from "react";

const Index = () => {
  const { data, error } = useCurrentUser();
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
        <Layout title="Dashboard">
          <span className="flex items-center">
            <UserAvatar data={data.user.email} size={30} />
            <span className="text-gray-400 text-2xl pl-2">
              {data.user.name}
            </span>
          </span>
        </Layout>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};

export default Index;
