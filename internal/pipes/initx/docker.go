package initx

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/tran/dfs"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/spf13/viper"
)

func GetBotName() string {
	botwayConfigPath := filepath.Join(".botway.yaml")

	c := config.New(".")
	c.AddDriver(yaml.Driver)
	c.LoadFiles(botwayConfigPath)

	return c.String("bot.name")
}

func DockerInit() {
	err := dfs.CreateDirectory(filepath.Join(constants.HomeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(constants.BotwayDirPath())
	viper.SetConfigName("botway")
	viper.SetConfigType("json")

	viper.SetDefault("botway.bots." + GetBotName() + ".discord_token", os.Getenv("DISCORD_TOKEN"))
	viper.SetDefault("botway.bots." + GetBotName() + ".discord_client_id", os.Getenv("DISCORD_CLIENT_ID"))

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

	fmt.Println("🐋 Done")
}
