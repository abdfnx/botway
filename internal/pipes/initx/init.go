package initx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/tran/dfs"
	"github.com/spf13/viper"
)

func Init() {
	err := dfs.CreateDirectory(filepath.Join(constants.HomeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(constants.BotwayDirPath)
	viper.SetConfigName("botway")
	viper.SetConfigType("json")

	viper.SetDefault("botway.bots", map[string]string{})
	viper.SetDefault("botway.settings.auto_sync", true)
	viper.SetDefault("botway.settings.check_updates", true)
	viper.SetDefault("botway.bots_names", []string{})

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
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Initialization Successful"))
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Initialization Failed, try again"))
	}
}
