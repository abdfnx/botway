local yaml = require "lyaml"
local path = require "path"
local json = require "lunajson"

local BotwayConfigFile = io.open(path.join(path.user_home(), ".botway", "botway.json"))

if BotwayConfigFile == nil then
    print("ERROR: Botway config file not found")
    return
end

local BotwayConfig = json.decode(BotwayConfigFile:read("a"))
local BotConfig = io.open(".botway.yaml", "r")

if BotConfig == nil then
    print("ERROR: Bot config file not found")
    return
end

local Botway = {}

function Botway.GetBotInfo(value)
    local config = yaml.load(BotConfig:read("a"))

    return config["bot"][value]
end

function Botway.GetToken()
    return BotwayConfig["botway"]["bots"][Botway.GetBotInfo("name")]["bot_token"]
end

function Botway.GetAppId()
    return BotwayConfig["botway"]["bots"][Botway.GetBotInfo("name")]["bot_app_id"]
end

function Botway.GetGuildId(serverName)
    return BotwayConfig["botway"]["bots"][Botway.GetBotInfo("name")]["guilds"][serverName]["server_id"]
end

return Botway
