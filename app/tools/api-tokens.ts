import { base64url } from "jose";
import { toast } from "react-hot-toast";
import { toastStyle } from "./toast-style";

const key: any = process.env.NEXT_PUBLIC_BW_SECRET_KEY;

export const BW_SECRET_KEY = base64url.decode(key);

const message = (msg: string) => {
  return toast.error(
    `Your ${msg} is not set, please set your ${msg.toLowerCase()} in the settings page`,
    toastStyle
  );
};

export const CheckAPITokens = (user: any, hostService: any) => {
  if (hostService) {
    if (hostService == "railway" && !user.railwayApiToken) {
      message("Railway API Token");
    } else if (hostService == "render" && !user.renderApiToken) {
      message("Render API Token");
    }
  } else if (!user.githubApiToken) {
    message("GitHub Token");
  }
};
