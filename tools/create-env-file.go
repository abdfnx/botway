package tools

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/config"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func CreateEnvFile() {
	viper.AddConfigPath("config")
	viper.SetConfigName("botway-tokens")
	viper.SetConfigType("env")

	botType := botwaygo.GetBotInfo("bot.type")
	field1 := ""
	field2 := ""
	field3 := ""
	cid := ""
	secret_value := ""

	if botType == "discord" {
		field1 = "DISCORD_TOKEN"
		field2 = "DISCORD_CLIENT_ID"
		cid = "bot_app_id"
	} else if botType == "slack" {
		field1 = "SLACK_TOKEN"
		field2 = "SLACK_APP_TOKEN"
		field3 = "SLACK_SIGNING_SECRET"
		cid = "bot_app_token"
		secret_value = "signing_secret"
	} else if botType == "telegram" {
		field1 = "TELEGRAM_TOKEN"
	} else if botType == "twitch" {
		field1 = "TWITCH_OAUTH_TOKEN"
		field2 = "TWITCH_CLIENT_ID"
		field3 = "TWITCH_CLIENT_SECRET"
		cid = "bot_client_id"
		secret_value = "bot_client_secret"
	}

	viper.SetDefault(field1, config.Get("botway.bots."+botwaygo.GetBotInfo("bot.name")+".bot_token"))

	if botType != "telegram" {
		viper.SetDefault(field2, config.Get("botway.bots."+botwaygo.GetBotInfo("bot.name")+"."+cid))
	}

	if botType == "slack" || botType == "twitch" {
		viper.SetDefault(field3, config.Get("botway.bots."+botwaygo.GetBotInfo("bot.name")+"."+secret_value))
	}

	if botType == "discord" {
		if constants.Gerr != nil {
			panic(constants.Gerr)
		} else {
			guilds := gjson.Get(string(constants.Guilds), "guilds.#")

			for x := 0; x < int(guilds.Int()); x++ {
				server := gjson.Get(string(constants.Guilds), "guilds."+fmt.Sprint(x)).String()

				sgi := strings.ToUpper(server) + "_GUILD_ID"

				viper.Set(sgi, config.Get("botway.bots."+botwaygo.GetBotInfo("bot.name")+".guilds."+server+".server_id"))
			}
		}
	}

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}

	fmt.Println(constants.HEADING + constants.BOLD.Render("Done"))
}
