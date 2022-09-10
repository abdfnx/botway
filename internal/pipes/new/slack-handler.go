package new

import (
	"github.com/abdfnx/botway/templates/nodejs"
	"github.com/abdfnx/botway/templates/slack/python/pip"
	"github.com/abdfnx/botway/templates/slack/python/pipenv"
	"github.com/abdfnx/botway/templates/slack/python/poetry"
	"github.com/abdfnx/botway/templates/ts"
)

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 0 {
		pip.SlackPythonPip(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 1 {
		pipenv.SlackPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 2 {
		poetry.SlackPythonPoetry(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 0 {
		nodejs.Nodejs(opts.BotName, "npm", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 1 {
		nodejs.Nodejs(opts.BotName, "yarn", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 2 {
		nodejs.Nodejs(opts.BotName, "pnpm", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 3 {
		nodejs.Nodejs(opts.BotName, "bun", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 0 {
		ts.NodejsTS(opts.BotName, "npm", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 1 {
		ts.NodejsTS(opts.BotName, "yarn", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 2 {
		ts.NodejsTS(opts.BotName, "pnpm", "slack")
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 3 {
		ts.NodejsTS(opts.BotName, "bun", "slack")
	}
}
