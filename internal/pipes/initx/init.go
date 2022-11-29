package initx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/token"
	"github.com/abdfnx/tran/dfs"
	"github.com/spf13/viper"
)

func Init() {
	err := dfs.CreateDirectory(filepath.Join(constants.HomeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	botwayConfig := viper.New()

	botwayConfig.AddConfigPath(constants.BotwayDirPath)
	botwayConfig.SetConfigName("botway")
	botwayConfig.SetConfigType("json")

	botwayConfig.SetDefault("botway.bots", map[string]string{})
	botwayConfig.SetDefault("botway.settings.auto_sync", true)
	botwayConfig.SetDefault("botway.settings.check_updates", true)
	botwayConfig.SetDefault("botway.bots_names", []string{})
	botwayConfig.SetDefault("user.token", token.Generator())

	if err := botwayConfig.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = botwayConfig.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := botwayConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(constants.BotwayConfigFile); err == nil {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Initialization Successful"))
		fmt.Print(constants.INFO_BACKGROUND.Render("NEXT STEP"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" You can login into github by running " + constants.COMMAND_FOREGROUND.Render("`botway gh login`")))
		fmt.Println(constants.INFO_FOREGROUND.Render("Or You can get started and create your first bot by running " + constants.COMMAND_FOREGROUND.Render("`botway new BOT_NAME`")))
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Initialization Failed, try again"))
	}
}
