import { LoadingDots } from "@/components/LoadingDots";
import { useCurrentUser } from "@/lib/user";
import Layout from "@/components/Layout";
import { UserAvatar } from "../../components/UserAvatar";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { ProjectList } from "./ProjectList";
import { NewProject } from "../New/NewProject";

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
        <LoadingDots className="fixed inset-0 flex items-center justify-center" />
      ) : data?.user ? (
        <Layout title="Dashboard">
          <div className="flex items-center justify-between gap-4">
            <div className="flex-1 gap-2 justify-end flex-shrink-0">
              <a className="h-9 mt-1 px-4.5 inline-flex flex-shrink-0 whitespace-nowrap items-center gap-2">
                <UserAvatar data={data.user.email} size={30} />
                <span className="text-gray-400 text-2xl pl-2">
                  {data.user.name}
                </span>
              </a>
            </div>
            <div className="flex gap-2 justify-end flex-shrink-0">
              <NewProject />
            </div>
          </div>

          <ProjectList />
        </Layout>
      ) : (
        <PushToSignIn />
      )}
    </>
  );
};

export default Index;
