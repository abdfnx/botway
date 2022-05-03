import { load } from "js-yaml";
import { readFileSync } from "fs";
import { BOTWAY_CONFIG_PATH } from "./constants";

function format(data: string): string {
  return JSON.stringify(data);
}

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

export const GetToken = () => {
  if (getBotInfo("lang") != "nodejs") {
    console.log("ERROR: Botway is not running in NodeJS");
  } else {
    try {
      const contents = readFileSync(BOTWAY_CONFIG_PATH, "utf8");

      const json = JSON.parse(contents);

      return json["botway"]["bots"][getBotInfo("name")]["bot_token"];
    } catch (err: any) {
      console.log(err.stack || String(err));
    }
  }
};

export const GetAppId = () => {
  if (getBotInfo("lang") != "nodejs") {
    console.log("ERROR: Botway is not running in NodeJS");
  } else {
    if (getBotInfo("type") != "discord") {
      console.log(
        "ERROR: This function/feature is only working with discord bots."
      );
    } else {
      try {
        const contents = readFileSync(BOTWAY_CONFIG_PATH, "utf8");

        const json = JSON.parse(contents);

        return json["botway"]["bots"][getBotInfo("name")]["bot_app_id"];
      } catch (err: any) {
        console.log(err.stack || String(err));
      }
    }
  }
};

export const GetGuildId = (serverName: string) => {
  if (getBotInfo("lang") != "nodejs") {
    console.log("ERROR: Botway is not running in NodeJS");
  } else {
    if (getBotInfo("type") != "discord") {
      console.log(
        "ERROR: This function/feature is only working with discord bots."
      );
    } else {
      try {
        const contents = readFileSync(BOTWAY_CONFIG_PATH, "utf8");

        const json = JSON.parse(contents);

        return json["botway"]["bots"][getBotInfo("name")]["guilds"][serverName][
          "server_id"
        ];
      } catch (err: any) {
        console.log(err.stack || String(err));
      }
    }
  }
};
