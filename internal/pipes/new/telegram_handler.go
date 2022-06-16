package new

import (
	"github.com/abdfnx/botway/templates/telegram/go"
	"github.com/abdfnx/botway/templates/telegram/nodejs"
	"github.com/abdfnx/botway/templates/telegram/python/pip"
	"github.com/abdfnx/botway/templates/telegram/python/pipenv"
	"github.com/abdfnx/botway/templates/telegram/ruby"
	"github.com/abdfnx/botway/templates/telegram/rust"
)

func TelegramHandler(m model) {
	if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.TelegramPythonPip(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.TelegramPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 1 {
		tgo.TelegramGo(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMCoice == 0 {
		nodejs.TelegramNodejs(opts.BotName, "npm")
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMCoice == 1 {
		nodejs.TelegramNodejs(opts.BotName, "yarn")
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMCoice == 2 {
		nodejs.TelegramNodejs(opts.BotName, "pnpm")
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 {
		ruby.TelegramRuby(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 4 && m.PMCoice == 0 {
		rust.TelegramRust(opts.BotName, "cargo")
	} else if m.PlatformChoice == 1 && m.LangChoice == 4 && m.PMCoice == 1 {
		rust.TelegramRust(opts.BotName, "fleet")
	}
}
