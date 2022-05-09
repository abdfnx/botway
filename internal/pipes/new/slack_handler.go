package new

import (
	"github.com/abdfnx/botway/templates/slack/go"
	"github.com/abdfnx/botway/templates/slack/nodejs"
	"github.com/abdfnx/botway/templates/slack/python/pip"
	"github.com/abdfnx/botway/templates/slack/python/pipenv"
	"github.com/abdfnx/botway/templates/slack/ruby"
)

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.SlackPythonPip(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.SlackPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 {
		sgo.SlackGo(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMCoice == 0 {
		nodejs.SlackNodejs(opts.BotName, "npm")
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMCoice == 1 {
		nodejs.SlackNodejs(opts.BotName, "yarn")
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMCoice == 2 {
		nodejs.SlackNodejs(opts.BotName, "pnpm")
	} else if m.PlatformChoice == 2 && m.LangChoice == 3 {
		ruby.SlackRuby(opts.BotName)
	}
}
