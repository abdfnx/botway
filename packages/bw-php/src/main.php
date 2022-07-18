<?php
namespace Botway;

require __DIR__."/../vendor/autoload.php";

use Dallgoot\Yaml;

$BotwayConfig = json_decode(file_get_contents(pathJoin(homeDir(), ".botway", "botway.json")), true);

class Botway {
    function GetBotInfo(string $value) {
        $botConfig = Yaml::parseFile(".botway.yaml", 0, 0);

        return $botConfig->bot->$value;
    }

    function BotwayConfig(string $value) {
        global $BotwayConfig;

        return $BotwayConfig["botway"]["bots"][$this->GetBotInfo("name")][$value];
    }

    function GetToken() {
        return $this->BotwayConfig("bot_token");
    }

    function GetAppId() {
        if ($this->GetBotInfo("type") == "slack") {
            return $this->BotwayConfig("bot_app_token");
        } else {
            return $this->BotwayConfig("bot_app_id");
        }
    }

    function GetGuildId(string $serverName) {
        if ($this->GetBotInfo("type") != "discord") {
            echo "ERROR: This function/feature is only working with discord bots";
        } else {
            global $BotwayConfig;

            return $BotwayConfig["botway"]["bots"][$this->GetBotInfo("name")]["guilds"][$serverName]["server_id"];
        }
    }

    function GetSigningSecret() {
        if ($this->GetBotInfo("type") != "slack") {
            echo "ERROR: This function/feature is only working with slack bots";
        } else {
            return $this->BotwayConfig("signing_secret");
        }
    }
}
