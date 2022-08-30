package config

import (
	"github.com/abdfnx/botway/constants"
	"github.com/tidwall/gjson"
)

func Get(value string) string {
	c := gjson.Get(string(constants.BotwayConfig), value)

	return c.String()
}
