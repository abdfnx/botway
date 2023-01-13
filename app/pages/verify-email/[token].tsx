import { findAndDeleteTokenByIdAndType, updateUserById } from "@/api/db";
import { getMongoDb } from "@/api/mongodb";
import { VerifyEmail } from "@/page-components/VerifyEmail";
import { GetServerSideProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect } from "react";

export default function VerifyEmailPage({ valid }: any) {
  const router = useRouter();

  useEffect(() => {
    if (process.env.NEXT_PUBLIC_FULL != "true") {
      router.push("/");
    }
  }, []);

  return (
    <>
      <Head>
        <title>Botway - Email verification</title>
      </Head>
      <VerifyEmail valid={valid} />
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const db = await getMongoDb();

  const { token }: any = context.params;

  const deletedToken = await findAndDeleteTokenByIdAndType(
    db,
    token,
    "emailVerify"
  );

  if (!deletedToken) return { props: { valid: false } };

  await updateUserById(db, deletedToken.creatorId, {
    emailVerified: true,
  });

  return { props: { valid: true } };
};
