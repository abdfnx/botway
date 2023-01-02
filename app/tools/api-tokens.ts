import { toast } from "react-hot-toast";
import { bgSecondary } from "./colors";

export const CheckAPITokens = (user: any) => {
  if (!user.railwayApiToken) {
    toast.error(
      "Your Railway API Token is not set, please set your railway api token in the settings page",
      {
        style: {
          borderRadius: "10px",
          backgroundColor: bgSecondary,
          color: "#fff",
        },
      }
    );
  } else if (!user.githubApiToken) {
    toast.error(
      "Your GitHub Token is not set, please set your github token in the settings page",
      {
        style: {
          borderRadius: "10px",
          backgroundColor: bgSecondary,
          color: "#fff",
        },
      }
    );
  }
};
