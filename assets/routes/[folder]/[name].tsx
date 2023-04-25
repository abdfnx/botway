import { supabase } from "../../supabase/config.ts";
import { Handlers, PageProps } from "$fresh/server.ts";

export const handler: Handlers = {
  async GET(_, ctx) {
    const { folder, name } = ctx.params;

    const { data, error } = supabase.storage
      .from("cdn")
      .getPublicUrl(`${folder}/${name}`);

    if (error) {
      return ctx.render(null);
    }

    return ctx.render(data.publicUrl);
  },
};

export default function Page({ data }: PageProps) {
  if (!data) {
    return <h1>Not found</h1>;
  }

  return (
    <iframe src={data} className="w-full h-full min-h-screen">
      <p>This browser does not support file type!</p>
    </iframe>
  );
}
