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

    public function GetToken(): string {
        return $this->BotwayConfig("bot_token");
    }

    public function GetAppId(): string {
        if ($this->GetBotInfo("type") == "slack") {
            return $this->BotwayConfig("bot_app_token");
        } else {
            return $this->BotwayConfig("bot_app_id");
        }
    }

    public function GetGuildId(string $serverName): string {
        if ($this->GetBotInfo("type") != "discord") {
            echo "ERROR: This function/feature is only working with discord bots";
        } else {
            global $BotwayConfig;

            return $BotwayConfig["botway"]["bots"][$this->GetBotInfo("name")]["guilds"][$serverName]["server_id"];
        }
    }

    public function GetSigningSecret(): string {
        if ($this->GetBotInfo("type") != "slack") {
            echo "ERROR: This function/feature is only working with slack bots";
        } else {
            return $this->BotwayConfig("signing_secret");
        }
    }
}
