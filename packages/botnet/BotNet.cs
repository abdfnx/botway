namespace BotNet;

using System.IO;
using Newtonsoft.Json.Linq;
using YamlDotNet.Serialization;

public class Core {
    static string path = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.UserProfile), ".botway", "botway.json");

    dynamic BotwayConfigJson = JObject.Parse(File.ReadAllText(path));

    private dynamic BotwayConfig() {
        if (File.Exists(path)) {
            return BotwayConfigJson;
        } else {
            Exception e = new Exception("ERROR: Botway config file not found!");
            throw e;
        }
    }

    public string GetBotInfo(string value) {
        var deserializer = new Deserializer();

        if (File.Exists(".botway.yaml")) {
            dynamic BotConfig = deserializer.Deserialize<dynamic>(File.ReadAllText(".botway.yaml"));

            return BotConfig["bot"][value].ToString();
        } else {
            Exception e = new Exception("ERROR: Bot config file not found");
            throw e;
        }
    }

    public string GetToken() {
        return BotwayConfig()["botway"]["bots"][GetBotInfo("name")]["bot_token"].ToString();
    }

    public string GetAppId() {
        string value = "bot_app_id";

        if (GetBotInfo("type") == "slack") {
            value = "bot_app_token";
        }

        return BotwayConfig()["botway"]["bots"][GetBotInfo("name")][value].ToString();
    }

    public string GetGuildId(string serverName) {
        if (GetBotInfo("type") != "discord") {
            Exception e = new Exception("ERROR: This function/feature is only working with discord bots.");
            throw e;
        } else {
            return BotwayConfig()["botway"]["bots"][GetBotInfo("name")]["guilds"][serverName]["server_id"].ToString();
        }
    }

    public string GetSigningSecret() {
        if (GetBotInfo("type") != "slack") {
            Exception e = new Exception("ERROR: This function/feature is only working with slack bots.");
            throw e;
        } else {
            return BotwayConfig()["botway"]["bots"][GetBotInfo("name")]["signing_secret"].ToString();
        }
    }
}
