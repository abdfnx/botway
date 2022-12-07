import { findProjectById } from "@/api/db";
import { getMongoDb } from "@/api/mongodb";
import Layout from "@/components/Layout";
import { GetServerSideProps } from "next";
import Head from "next/head";

export default function ProjectPage({ project }: any) {
  if (typeof project.createdAt !== "string") {
    project.createdAt = new Date(project.createdAt);
  }

  return (
    <>
      <Head>
        <title>🤖 Botway - {project.name}</title>
      </Head>
      <Layout></Layout>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const db = await getMongoDb();

  const project = await findProjectById(db, context.params?.projectId);

  if (!project) {
    return {
      notFound: true,
    };
  }

  project._id = String(project._id);
  project.creatorId = String(project.creatorId);
  project.creator._id = String(project.creator._id);
  project.createdAt = project.createdAt.toJSON();

  return { props: { project } };
};
