import slug from "slug";

export const slugger = (username: any) => slug(username, "_");
