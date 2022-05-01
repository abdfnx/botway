package new

import (
	"github.com/abdfnx/botway/tools/templates/slack/go"
	"github.com/abdfnx/botway/tools/templates/slack/python/pip"
	"github.com/abdfnx/botway/tools/templates/slack/python/pipenv"
)

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.SlackPythonPip(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.SlackPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 {
		sgo.SlackGo(opts.BotName)
	}
}
