import { parse } from "https://deno.land/std@0.137.0/encoding/yaml.ts";
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
    console.log("ERROR: Botway is not running in Deno");
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
    console.log("ERROR: Botway is not running in Deno");
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
    console.log("ERROR: Botway is not running in Deno");
  } else if (getBotInfo("type") != "discord") {
    console.log("ERROR: Botway is not running in Discord");
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
