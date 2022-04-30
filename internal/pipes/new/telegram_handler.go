package new

import (
	"github.com/abdfnx/botway/tools/templates/telegram/go"
	"github.com/abdfnx/botway/tools/templates/telegram/nodejs/npm"
	"github.com/abdfnx/botway/tools/templates/telegram/nodejs/pnpm"
	"github.com/abdfnx/botway/tools/templates/telegram/nodejs/yarn"
	"github.com/abdfnx/botway/tools/templates/telegram/python/pip"
	"github.com/abdfnx/botway/tools/templates/telegram/python/pipenv"
	"github.com/abdfnx/botway/tools/templates/telegram/ruby"
	"github.com/abdfnx/botway/tools/templates/telegram/rust/cargo"
	"github.com/abdfnx/botway/tools/templates/telegram/rust/fleet"
)

func TelegramHandler(m model) {
	if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.TelegramPythonPip(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.TelegramPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 1 {
		tgo.TelegramGo(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMCoice == 0 {
		npm.TelegramNodejsNpm(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMCoice == 1 {
		yarn.TelegramNodejsYarn(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMCoice == 2 {
		pnpm.TelegramNodejsPnpm(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 {
		ruby.TelegramRuby(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 4 && m.PMCoice == 0 {
		cargo.TelegramRustCargo(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 4 && m.PMCoice == 1 {
		fleet.TelegramRustFleet(opts.BotName)
	}
}
