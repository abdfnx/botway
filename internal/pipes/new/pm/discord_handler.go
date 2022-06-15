package pm

import (
	"github.com/abdfnx/botway/templates/discord/deno"
	"github.com/abdfnx/botway/templates/discord/go"
	"github.com/abdfnx/botway/templates/discord/nodejs"
	"github.com/abdfnx/botway/templates/discord/python/pip"
	"github.com/abdfnx/botway/templates/discord/python/pipenv"
	"github.com/abdfnx/botway/templates/discord/ruby"
	"github.com/abdfnx/botway/templates/discord/rust"
)

func DiscordHandler(m model, botName, pm string) {
	if m.platform == "discord" && m.lang == "python" && pm == "pip" {
		pip.DiscordPythonPip(botName)
	} else if m.platform == "discord" && m.lang == "python" && pm == "pipenv" {
		pipenv.DiscordPythonPipenv(botName)
	} else if m.platform == "discord" && m.lang == "go" {
		dgo.DiscordGo(botName)
	} else if m.platform == "discord" && m.lang == "nodejs" {
		nodejs.DiscordNodejs(botName, pm)
	} else if m.platform == "discord" && m.lang == "ruby" {
		ruby.DiscordRuby(botName)
	} else if m.platform == "discord" && m.lang == "rust" {
		rust.DiscordRust(botName, pm)
	} else if m.platform == "discord" && m.lang == "deno" {
		deno.DiscordDeno(botName)
	}
}

