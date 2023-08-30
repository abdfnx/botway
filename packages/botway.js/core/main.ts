import { load } from "js-yaml";
import { readFileSync } from "fs";
import { BOTWAY_CONFIG_PATH } from "./constants";

const format = (data: string): string => {
  return JSON.stringify(data);
};

const getBotInfo = (value: string) => {
  try {
    const contents = readFileSync(".botway.yaml", "utf8"),
      data: any = load(contents);

    const raw = format(data);
    const json = JSON.parse(raw);

    return json["bot"][value];
  } catch (err: any) {
    console.log(err.stack || String(err));
  }
};

const contents = readFileSync(BOTWAY_CONFIG_PATH, "utf8");

const json = JSON.parse(contents);

export const GetToken = () => {
  try {
    return json["botway"]["bots"][getBotInfo("name")]["bot_token"];
  } catch (err: any) {
    console.log(err.stack || String(err));
  }
};

export const GetAppId = () => {
  try {
    if (getBotInfo("type") === "slack") {
      return json["botway"]["bots"][getBotInfo("name")]["bot_app_token"];
    } else {
      return json["botway"]["bots"][getBotInfo("name")]["bot_app_id"];
    }
  } catch (err: any) {
    console.log(err.stack || String(err));
  }
};

export const GetGuildId = (serverName: string) => {
  if (getBotInfo("type") != "discord") {
    console.log(
      "ERROR: This function/feature is only working with discord bots",
    );
  } else {
    try {
      return json["botway"]["bots"][getBotInfo("name")]["guilds"][serverName][
        "server_id"
      ];
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};

export const GetSecret = () => {
  if (getBotInfo("type") != "slack") {
    console.log("ERROR: This function/feature is only working with slack bots");
  } else {
    try {
      var value = "";

      if (getBotInfo("type") === "slack") {
        value = "signing_secret";
      } else if (getBotInfo("type") === "twitch") {
        value = "bot_client_secret";
      }

      return json["botway"]["bots"][getBotInfo("name")][value];
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};
