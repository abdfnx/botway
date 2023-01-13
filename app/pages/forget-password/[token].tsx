import { findTokenByIdAndType } from "@/api/db";
import { getMongoDb } from "@/api/mongodb";
import { ForgetPasswordToken } from "@/page-components/ForgetPassword/token";
import { GetServerSideProps } from "next";
import Head from "next/head";

const ResetPasswordTokenPage = ({ valid, token }: any) => {
  return (
    <>
      <Head>
        <title>Botway - Forget password</title>
      </Head>
      <ForgetPasswordToken valid={valid} token={token} />
    </>
  );
};

export const getServerSideProps: GetServerSideProps = async (context) => {
  const db = await getMongoDb();

  const { token }: any = context.params;

  const tokenDoc = await findTokenByIdAndType(db, token, "passwordReset");

  return { props: { token, valid: !!tokenDoc } };
};

export default ResetPasswordTokenPage;
