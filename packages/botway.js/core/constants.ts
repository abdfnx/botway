import { join } from "path";

export const HOMEDIR: any = process.env.HOME || process.env.USERPROFILE;
export const DOT_BOTWAY_PATH = join(HOMEDIR, ".botway");
export const BOTWAY_CONFIG_PATH: any = join(DOT_BOTWAY_PATH, "botway.json");
