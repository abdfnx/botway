<?php
namespace MyBot;

include __DIR__ . "/../vendor/autoload.php";

use Dallgoot\Yaml;

function pathJoin($path1, $path2) {
    $paths = func_get_args();
    $last_key = func_num_args() - 1;

    array_walk($paths, function (&$val, $key) use ($last_key) {
        switch ($key) {
            case 0:
                $val = rtrim($val, "/ ");

                break;

            case $last_key:
                $val = ltrim($val, "/ ");

                break;

            default:
                $val = trim($val, "/ ");

                break;
        }
    });

    $first = array_shift($paths);
    $last = array_pop($paths);
    $paths = array_filter($paths);

    array_unshift($paths, $first);

    $paths[] = $last;

    return implode("/", $paths);
}

function homeDir() {
    if (isset($_SERVER["HOME"])) {
        $result = $_SERVER["HOME"];
    } else {
        $result = getenv("HOME");
    }

    if (empty($result) && function_exists("exec")) {
        if(strncasecmp(PHP_OS, "WIN", 3) === 0) {
            $result = exec("echo %userprofile%");
        } else {
            $result = exec("echo ~");
        }
    }

    return $result;
}

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
