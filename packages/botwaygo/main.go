package botwaygo

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/tran/dfs"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var (
	homeDir, err = dfs.GetHomeDirectory()
	botwayConfigFile = filepath.Join(homeDir, ".botway", "botway.json")
	botwayConfig, berr = ioutil.ReadFile(botwayConfigFile)
)

func GetBotInfo(value string) string {
	viper.AddConfigPath(".")
	viper.SetConfigName(".botway")
	viper.SetConfigType("yaml")

	return viper.GetString("bot." + value)
}

func GetToken() string {
	if err != nil {
		log.Fatal(err)
	}

	if berr != nil {
		panic(berr)
	}

	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Botway is not running in Golang"))
	} else {
		data := gjson.Get(string(botwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_token").String()

		return data
	}

	return ""
}

func GetAppId() string {
	if err != nil {
		log.Fatal(err)
	}

	if berr != nil {
		panic(berr)
	}

	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Botway is not running in Golang"))
	} else {
		data := ""

		if GetBotInfo("type") == "slack" {
			data = gjson.Get(string(botwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_app_token").String()
		} else {
			data = gjson.Get(string(botwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_app_id").String()
		}

		return data
	}

	return ""
}

func GetGuildId(serverName string) string {
	if err != nil {
		log.Fatal(err)
	}

	if berr != nil {
		panic(berr)
	}

	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Botway is not running in Golang"))
	} else if GetBotInfo("type") != "discord" {
		panic(errors.New("ERROR: This function/feature is only working with discord bots."))
	} else {
		data := gjson.Get(string(botwayConfig), "botway.bots." + GetBotInfo("name") + ".guilds." + serverName + ".server_id").String()

		return data
	}

	return ""
}
