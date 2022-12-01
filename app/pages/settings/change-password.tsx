import { ChangePassword } from "@/page-components/Settings/change-password";
import Head from "next/head";

const ChangePasswordSettingsPage = () => {
  return (
    <>
      <Head>
        <title>🤖 Botway - Settings</title>
      </Head>
      <ChangePassword />
    </>
  );
};

export default ChangePasswordSettingsPage;
