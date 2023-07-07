import { base64url } from "jose";

export const BW_SECRET_KEY = base64url.decode(
  process.env.NEXT_PUBLIC_BW_SECRET_KEY!,
);
