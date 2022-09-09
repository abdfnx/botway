package render

import (
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
	"github.com/tidwall/gjson"
)

var (
	id          = gjson.Get(string(constants.RenderConfig), "user.id").String()
	apiToken    = gjson.Get(string(constants.RenderConfig), "user.api_token").String()
	serviceName = strings.ReplaceAll(botwaygo.GetBotInfo("bot.name"), " ", "%20")
	serviceId   = gjson.Get(string(constants.RenderConfig), "projects."+botwaygo.GetBotInfo("bot.name")+".id").String()
)
