import { findProjectById } from "@/api/db";
import { getMongoDb } from "@/api/mongodb";
import { useCurrentUser } from "@/lib/user";
import { Settings } from "@/page-components/Settings";
import Head from "next/head";
import { useRouter } from "next/router";

const ProjectPage = async () => {
  const { data: { user } = {} } = useCurrentUser();
  const router = useRouter();

  const projectId = router.query;

  const db = await getMongoDb();

  const project = await findProjectById(db, projectId);

  if (!project) {
    return <h1>404</h1>;
  }

  project._id = String(project._id);
  project.creatorId = String(project.creatorId);
  project.creator._id = String(project.creator._id);
  project.createdAt = project.createdAt.toJSON();

  if (project.creatorId != user._id) {
    return <h1>404</h1>;
  }

  return (
    <>
      <Head>
        <title>🤖 Botway - Settings</title>
      </Head>
      <h1>{project.name}</h1>
    </>
  );
};

export default ProjectPage;
