package new

import (
	"github.com/abdfnx/botway/tools/templates/slack/go"
	"github.com/abdfnx/botway/tools/templates/slack/nodejs/npm"
	"github.com/abdfnx/botway/tools/templates/slack/nodejs/yarn"
	"github.com/abdfnx/botway/tools/templates/slack/python/pip"
	"github.com/abdfnx/botway/tools/templates/slack/python/pipenv"
	"github.com/abdfnx/botway/tools/templates/slack/nodejs/pnpm"
)

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.SlackPythonPip(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.SlackPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 {
		sgo.SlackGo(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMCoice == 0 {
		npm.SlackNodejsNpm(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMCoice == 1 {
		yarn.SlackNodejsYarn(opts.BotName)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMCoice == 2 {
		pnpm.SlackNodejsPnpm(opts.BotName)
	}
}
