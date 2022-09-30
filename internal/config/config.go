package config

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func Get(value string) string {
	c := gjson.Get(string(constants.BotwayConfig), value)

	return c.String()
}

func GetBotInfoFromArg(botName, value string) string {
	c := viper.New()

	c.SetConfigType("yaml")

	botConfig, err := os.ReadFile(filepath.Join(botName, ".botway.yaml"))

	if err != nil {
		panic(err)
	}

	c.ReadConfig(bytes.NewBuffer(botConfig))

	return c.GetString(value)
}
