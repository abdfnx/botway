package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/tran/dfs"
	"github.com/gookit/config/v2"
	cyaml "github.com/gookit/config/v2/yaml"
)

func AddBotToConfig(botName, botType, botPath string) {
	homeDir, _ := dfs.GetHomeDirectory()
	botwayConfigPath := filepath.Join(homeDir, ".botway", "botway.yaml")
	botwayConfig, err := ioutil.ReadFile(botwayConfigPath)

	if err != nil {
		panic(err)
	}

	bc := config.New(".")
	bc.AddDriver(cyaml.Driver)
	bcp := bc.LoadSources(config.Yaml, botwayConfig)

	if bcp != nil {
		panic(bcp)
	}

	bc.Set("botway.bots." + botName + ".type", botType)
	bc.Set("botway.bots." + botName + ".path", botPath)

	remove := os.Remove(botwayConfigPath)

	if remove != nil {
        log.Fatal(remove)
    }

	newBotConfig := os.WriteFile(botwayConfigPath, []byte(string(tools.ToYaml(bc.Data()))), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}
