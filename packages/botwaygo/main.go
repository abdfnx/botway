package botwaygo

import (
	"errors"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func GetBotInfo(value string) string {
	viper.AddConfigPath(".")
	viper.SetConfigName(".botway")
	viper.SetConfigType("yaml")

	return viper.GetString("bot." + value)
}

func GetToken() string {
	if constants.Berr != nil {
		panic(constants.Berr)
	}

	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Botway is not running in Golang"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_token").String()

		return data
	}
}

func GetAppId() string {
	if constants.Berr != nil {
		panic(constants.Berr)
	}

	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Botway is not running in Golang"))
	} else {
		data := ""

		if GetBotInfo("type") == "slack" {
			data = gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_app_token").String()
		} else {
			data = gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_app_id").String()
		}

		return data
	}
}

func GetGuildId(serverName string) string {
	if constants.Berr != nil {
		panic(constants.Berr)
	}

	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Botway is not running in Golang"))
	} else if GetBotInfo("type") != "discord" {
		panic(errors.New("ERROR: This function/feature is only working with discord bots."))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".guilds." + serverName + ".server_id").String()

		return data
	}
}
