package tools

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
)

func GenerateCConfig(token string) {
	viper.AddConfigPath(constants.BotwayDirPath)
	viper.SetConfigName("botway-c-config")
	viper.SetConfigType("json")

	viper.SetDefault("logging.level", "trace")
	viper.SetDefault("logging.filename", "bot.log")
	viper.SetDefault("logging.quiet", false)
	viper.SetDefault("logging.overwrite", true)
	viper.SetDefault("logging.use_color", true)
	viper.SetDefault("logging.http.enable", true)
	viper.SetDefault("logging.http.filename", "http.log")
	viper.SetDefault("logging.disable_modules", []string{"WEBSOCKETS", "USER_AGENT"})
	viper.SetDefault("discord.token", token)
	viper.SetDefault("discord.default_prefix.enable", true)

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

	if _, err := os.Stat(constants.BotwayConfigFile); err == nil {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Botway C Config Initialization Successful"))
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Botway C Config Initialization Failed, try again"))
	}
}
