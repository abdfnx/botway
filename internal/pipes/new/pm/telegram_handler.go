package pm

import (
	"github.com/abdfnx/botway/templates/telegram/go"
	"github.com/abdfnx/botway/templates/telegram/nodejs"
	"github.com/abdfnx/botway/templates/telegram/python/pip"
	"github.com/abdfnx/botway/templates/telegram/python/pipenv"
	"github.com/abdfnx/botway/templates/telegram/ruby"
	"github.com/abdfnx/botway/templates/telegram/rust"
)

func TelegramHandler(m model, botName, pm string) {
	if m.platform == "telegram" && m.lang == "python" && pm == "pip" {
		pip.TelegramPythonPip(botName)
	} else if m.platform == "telegram" && m.lang == "python" && pm == "pipenv" {
		pipenv.TelegramPythonPipenv(botName)
	} else if m.platform == "telegram" && m.lang == "go" {
		tgo.TelegramGo(botName)
	} else if m.platform == "telegram" && m.lang == "nodejs" {
		nodejs.TelegramNodejs(botName, pm)
	} else if m.platform == "telegram" && m.lang == "ruby" {
		ruby.TelegramRuby(botName)
	} else if m.platform == "telegram" && m.lang == "rust" {
		rust.TelegramRust(botName, pm)
	}
}
