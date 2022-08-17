package app

import (
	"bytes"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
)

func BotConfig(value string) string {
	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	return viper.GetString(value)
}
