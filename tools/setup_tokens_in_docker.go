package tools

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func SetupTokensInDocker() {
	if _, err := os.Stat(".botway.yaml"); err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Print(" ")
		panic(constants.FAIL_FOREGROUND.Render("You need to run this command in your bot directory"))
	}

	botName := initx.GetBotName()
	t := initx.GetBotType()

	bot_token := ""
	app_token := ""
	cid := ""

	if t == "discord" {
		bot_token = "DISCORD_TOKEN"
		app_token = "DISCORD_CLIENT_ID"
		cid = "bot_app_id"
	} else if t == "slack" {
		bot_token = "SLACK_TOKEN"
		app_token = "SLACK_APP_TOKEN"
		cid = "bot_app_token"
	} else if t == "telegram" {
		bot_token = "TELEGRAM_TOKEN"
	}

	botwayConfig := viper.New()
	env := viper.New()

	botwayConfig.SetConfigType("json")

	if constants.Berr != nil {
		panic(constants.Berr)
	}

	botwayConfig.ReadConfig(bytes.NewBuffer(constants.BotwayConfig))
	botPath := botwayConfig.GetString("botway.bots." + botName + ".path")

	env.AddConfigPath(filepath.Join(botPath, "config"))
	env.SetConfigName("botway-deploy-tokens")
	env.SetConfigType("env")

	bot_token_content := botwayConfig.GetString("botway.bots." + botName + ".bot_token")
	app_token_content := botwayConfig.GetString("botway.bots." + botName + "." + cid)

	if bot_token_content == "" || app_token_content == "" {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Print(" ")
		panic(constants.FAIL_FOREGROUND.Render("You didn't set bot token or app token"))
	}

	env.SetDefault(bot_token, bot_token_content)

	if t != "telegram" {
		env.SetDefault(app_token, app_token_content)
	}

	if t == "discord" {
		if constants.Gerr != nil {
			panic(constants.Gerr)
		} else {
			guilds := gjson.Get(string(constants.Guilds), "guilds.#")

			for x := 0; x < int(guilds.Int()); x++ {
				server := gjson.Get(string(constants.Guilds), "guilds."+fmt.Sprint(x)).String()

				sgi := strings.ToUpper(server) + "_GUILD_ID"
				sgi_content := botwayConfig.GetString("botway.bots." + botName + ".guilds." + server + ".server_id")

				env.Set(sgi, sgi_content)
			}
		}
	}

	if err := env.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = env.WriteConfig()

			if err != nil {
				panic(err)
			}
		}
	}

	if err := env.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		}
	}
}

func RemoveConfig() {
	os.Remove(filepath.Join("config", "botway-deploy-tokens.env"))
}
