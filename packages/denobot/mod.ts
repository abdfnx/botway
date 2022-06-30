import { parse } from "https://deno.land/std/encoding/yaml.ts";
import { readJson } from "./json.ts";

let botwayConfigPath: any = Deno.env.get("HOME") + "/.botway" + "/botway.json";

if (Deno.build.os == "windows") {
  botwayConfigPath =
    Deno.env.get("USERPROFILE") + "\\.botway" + "\\botway.json";
}

const botwayConfig: any = await readJson(botwayConfigPath);
const botConfig: any = parse(await Deno.readTextFile("./.botway.yaml"));

const getBotInfo = (value: string) => {
  return botConfig["bot"][value];
};

export const getToken = () => {
  if (getBotInfo("lang") != "deno") {
    console.log("ERROR: Your Bot framework is not Deno");
  } else {
    try {
      return botwayConfig["botway"]["bots"][getBotInfo("name")]["bot_token"];
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};

export const getAppId = () => {
  if (getBotInfo("lang") != "deno") {
    console.log("ERROR: Your Bot framework is not Deno");
  } else {
    try {
      if (getBotInfo("type") == "slack") {
        return botwayConfig["botway"]["bots"][getBotInfo("name")][
          "bot_app_token"
        ];
      } else {
        return botwayConfig["botway"]["bots"][getBotInfo("name")]["bot_app_id"];
      }
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};

export const getGuildId = (serverName: string) => {
  if (getBotInfo("lang") != "deno") {
    console.log("ERROR: Your Bot framework is not Deno");
  } else if (getBotInfo("type") != "discord") {
    console.log(
      "ERROR: This function/feature is only working with discord bots"
    );
  } else {
    try {
      return botwayConfig["botway"]["bots"][getBotInfo("name")]["guilds"][
        serverName
      ]["server_id"];
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};

export const getSigningSecret = () => {
  if (getBotInfo("lang") != "deno") {
    console.log("ERROR: Your Bot framework is not Deno");
  } else if (getBotInfo("type") != "slack") {
    console.log("ERROR: This function/feature is only working with slack bots");
  } else {
    try {
      return botwayConfig["botway"]["bots"][getBotInfo("name")][
        "signing_secret"
      ];
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};
