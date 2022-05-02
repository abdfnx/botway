package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/tran/dfs"
	"github.com/tidwall/sjson"
)

func AddBotToConfig(botName, botType, botPath string) {
	homeDir, _ := dfs.GetHomeDirectory()
	botwayConfigPath := filepath.Join(homeDir, ".botway", "botway.json")
	botwayConfig, err := ioutil.ReadFile(botwayConfigPath)

	if err != nil {
		panic(err)
	}

	botTypeContent, _ := sjson.Set(string(botwayConfig), "botway.bots." + botName + ".type", botType)
	botPathContent, _ := sjson.Set(botTypeContent, "botway.bots." + botName + ".path", botPath)

	remove := os.Remove(botwayConfigPath)

	if remove != nil {
        log.Fatal(remove)
    }

	newBotConfig := os.WriteFile(botwayConfigPath, []byte(botPathContent), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}
