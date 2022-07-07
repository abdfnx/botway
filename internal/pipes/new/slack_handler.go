package new

import (
	"github.com/abdfnx/botway/templates/slack/nodejs"
	"github.com/abdfnx/botway/templates/slack/python/pip"
	"github.com/abdfnx/botway/templates/slack/python/pipenv"
	"github.com/abdfnx/botway/templates/slack/python/poetry"
)

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.SlackPythonPip(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.SlackPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 2 {
		poetry.SlackPythonPoetry(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMCoice == 0 {
		nodejs.SlackNodejs(opts.BotName, "npm")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMCoice == 1 {
		nodejs.SlackNodejs(opts.BotName, "yarn")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMCoice == 2 {
		nodejs.SlackNodejs(opts.BotName, "pnpm")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMCoice == 3 {
		nodejs.SlackNodejs(opts.BotName, "bun")
	}
}
