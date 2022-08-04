import os
import std/json

let botwayConfig = readFile(joinPath(getHomeDir(), ".botway", "botway.json"))
let config = parseJson(botwayConfig)["botway"]["bots"]["mox"]

proc GetToken*(): string =
  return config["bot_token"].getStr

proc GetAppId*(): string =
  return config["bot_app_id"].getStr

proc GetGuildId*(serverName: string): string =
  if config["type"].getStr != "discord":
    raise newException(IOError, "ERROR: This function/feature is only working with discord bots")
  else:
    return config["guilds"][serverName]["server_id"].getStr
