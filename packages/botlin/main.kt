package botway

import java.io.File
import java.nio.file.Paths
import net.mamoe.yamlkt.*
import com.beust.klaxon.*

val BotwayConfigPath = Paths.get(System.getProperty("user.home"), ".botway", "botway.json").toString()
val BotwayConfig = File(BotwayConfigPath).readText().trimIndent()

fun getBotInfo(value: String): String {
    val botConfigPath = Paths.get(System.getProperty("user.dir"), "..").toString()

    val botConfig = File(botConfigPath, ".botway.yaml").readText().trimIndent()

    val botConfigMap: YamlMap = Yaml.decodeYamlMapFromString(botConfig)

    val config = botConfigMap["bot"] as YamlMap

    return config[value].toString()
}

val parser: Parser = Parser.default()

val stringBuilder: StringBuilder = StringBuilder(BotwayConfig)

val botwayJson: JsonObject = parser.parse(stringBuilder) as JsonObject

val botway = botwayJson.get("botway") as JsonObject

val bots = botway.get("bots") as JsonObject

val bot = bots.get(getBotInfo("name")) as JsonObject

fun GetToken(): String {
    return bot.get("bot_token").toString()
}

fun GetAppId(): String {
    return bot.get("bot_app_id").toString()
}

fun GetGuildId(serverName: String): String {
    if (getBotInfo("type") != "discord") {
        throw Exception("ERROR: This function/feature is only working with discord bots")
    } else {
        val guilds = bot.get("guilds") as JsonObject

        val sn = guilds.get(serverName) as JsonObject

        return sn.get("server_id").toString()
    }
}

fun GetSecret(): String {
    var value = ""

    if (getBotInfo("type") == "slack") {
        value = "signing_secret"
    } else if (getBotInfo("type") == "twitch") {
        value = "bot_client_secret"
    }

    return bot.get(value).toString()
}
