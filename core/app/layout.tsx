import Head from "./head";
import createClient from "@/supabase/server";
import { AuthProvider } from "@/supabase/auth/provider";
import "@/assets/main.scss";

// do not cache this layout
export const revalidate = 0;

const RootLayout = async ({ children }: { children: React.ReactNode }) => {
  const supabase = createClient();

  const {
    data: { session },
  } = await supabase.auth.getSession();

  const accessToken = session?.access_token || null;

  return (
    <html lang="en">
      <body>
        <Head />
        <AuthProvider accessToken={accessToken}>{children}</AuthProvider>
      </body>
    </html>
  );
};

export default RootLayout;
