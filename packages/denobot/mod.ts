import { parse } from "https://deno.land/std/encoding/yaml.ts";
import { readJson } from "https://deno.land/x/deno_json/mod.ts";

let botwayConfigPath = Deno.env.get("HOME") + "/.botway" + "/botway.json";

if (Deno.build.os === "windows") {
  botwayConfigPath =
    Deno.env.get("USERPROFILE") + "\\.botway" + "\\botway.json";
}

const botwayConfig: any = await readJson(botwayConfigPath);
const botConfig: any = parse(await Deno.readTextFile("./.botway.yaml"));

const getBotInfo = (value: string) => {
  return botConfig["bot"][value];
};

export const getToken = () => {
  try {
    return botwayConfig["botway"]["bots"][getBotInfo("name")]["bot_token"];
  } catch (err) {
    console.log(err.stack || String(err));
  }
};

export const getAppId = () => {
  try {
    if (getBotInfo("type") === "slack") {
      return botwayConfig["botway"]["bots"][getBotInfo("name")][
        "bot_app_token"
      ];
    } else {
      return botwayConfig["botway"]["bots"][getBotInfo("name")]["bot_app_id"];
    }
  } catch (err) {
    console.log(err.stack || String(err));
  }
};

export const getGuildId = (serverName: string) => {
  if (getBotInfo("type") != "discord") {
    console.log(
      "ERROR: This function/feature is only working with discord bots",
    );
  } else {
    try {
      return botwayConfig["botway"]["bots"][getBotInfo("name")]["guilds"][
        serverName
      ]["server_id"];
    } catch (err) {
      console.log(err.stack || String(err));
    }
  }
};

export const getSecret = () => {
  try {
    let value = "";

    if (getBotInfo("type") === "slack") {
      value = "signing_secret";
    } else if (getBotInfo("type") === "twitch") {
      value = "bot_client_secret";
    }

    return botwayConfig["botway"]["bots"][getBotInfo("name")][value];
  } catch (err) {
    console.log(err.stack || String(err));
  }
};
