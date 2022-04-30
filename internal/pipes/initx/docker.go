package initx

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/tran/dfs"
	"github.com/spf13/viper"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func GetBotName() string {
	botwayConfigPath := filepath.Join(".botway.yaml")

	c := config.New(".")
	c.AddDriver(yaml.Driver)
	c.LoadFiles(botwayConfigPath)

	return c.String("bot.name")
}

func DockerInit() {
	var err error

	homeDir, err := dfs.GetHomeDirectory()

	if err != nil {
		log.Fatal(err)
	}

	err = dfs.CreateDirectory(filepath.Join(homeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	botwayDirPath := ""

	if runtime.GOOS == "windows" {
		botwayDirPath = `$HOME\\.botway`
	} else {
		botwayDirPath = `$HOME/.botway`
	}

	viper.AddConfigPath(botwayDirPath)
	viper.SetConfigName("botway")
	viper.SetConfigType("yaml")

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
