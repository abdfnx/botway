package tools

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var (
	botName = botwaygo.GetBotInfo("bot.name")
	botType = botwaygo.GetBotInfo("bot.type")

	bot_token      = ""
	app_token      = ""
	signing_secret = "SLACK_SIGNING_SECRET"
	cid            = ""
)

func SetupTokensInDocker() {
	CheckDir()

	if botType == "discord" {
		bot_token = "DISCORD_TOKEN"
		app_token = "DISCORD_CLIENT_ID"
		cid = "bot_app_id"
	} else if botType == "slack" {
		bot_token = "SLACK_TOKEN"
		app_token = "SLACK_APP_TOKEN"
		cid = "bot_app_token"
	} else if botType == "telegram" {
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
	env.SetConfigName("botway-tokens")
	env.SetConfigType("env")

	bot_token_content := botwayConfig.GetString("botway.bots." + botName + ".bot_token")
	app_token_content := botwayConfig.GetString("botway.bots." + botName + "." + cid)

	if bot_token_content == "" || botType != "telegram" && app_token_content == "" || botType == "slack" && signing_secret == "" {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Print(" ")
		panic(constants.FAIL_FOREGROUND.Render("You didn't set bot token or app token or signing secret"))
	}

	env.SetDefault(bot_token, bot_token_content)

	if botType != "telegram" {
		env.SetDefault(app_token, app_token_content)
	}

	if botType == "discord" {
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

	if botType == "slack" {
		signing_secret_content := botwayConfig.GetString("botway.bots." + botName + ".signing_secret")

		env.SetDefault(signing_secret, signing_secret_content)
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

func SetupTokensInDockerRender() {
	CheckDir()

	if botType == "discord" {
		bot_token = "DISCORD_TOKEN"
		app_token = "DISCORD_CLIENT_ID"
		cid = "bot_app_id"
	} else if botType == "slack" {
		bot_token = "SLACK_TOKEN"
		app_token = "SLACK_APP_TOKEN"
		cid = "bot_app_token"
	} else if botType == "telegram" {
		bot_token = "TELEGRAM_TOKEN"
	}

	env := viper.New()

	pwd, _ := os.Getwd()

	env.AddConfigPath(filepath.Join(pwd, "config"))
	env.SetConfigName("botway-tokens")
	env.SetConfigType("env")

	env.SetDefault(bot_token, os.Getenv(bot_token))

	if botType != "telegram" {
		env.SetDefault(app_token, os.Getenv(bot_token))
	}

	if botType == "discord" {
		if constants.Gerr != nil {
			panic(constants.Gerr)
		} else {
			guilds := gjson.Get(string(constants.Guilds), "guilds.#")

			for x := 0; x < int(guilds.Int()); x++ {
				server := gjson.Get(string(constants.Guilds), "guilds."+fmt.Sprint(x)).String()

				sgi := strings.ToUpper(server) + "_GUILD_ID"

				env.Set(sgi, os.Getenv(sgi))
			}
		}
	}

	if botType == "slack" {
		env.SetDefault(signing_secret, os.Getenv(signing_secret))
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
	os.Remove(filepath.Join("config", "botway-tokens.env"))
}
