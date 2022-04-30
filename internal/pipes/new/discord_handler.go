package new

import (
	"github.com/abdfnx/botway/tools/templates/discord/deno"
	"github.com/abdfnx/botway/tools/templates/discord/go"
	"github.com/abdfnx/botway/tools/templates/discord/nodejs/npm"
	"github.com/abdfnx/botway/tools/templates/discord/nodejs/pnpm"
	"github.com/abdfnx/botway/tools/templates/discord/nodejs/yarn"
	"github.com/abdfnx/botway/tools/templates/discord/python/pip"
	"github.com/abdfnx/botway/tools/templates/discord/python/pipenv"
	"github.com/abdfnx/botway/tools/templates/discord/ruby"
	"github.com/abdfnx/botway/tools/templates/discord/rust/cargo"
	"github.com/abdfnx/botway/tools/templates/discord/rust/fleet"
)

func DiscordHandler(m model) {
	if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMCoice == 0 {
		pip.DiscordPythonPip(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMCoice == 1 {
		pipenv.DiscordPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 1 {
		dgo.DiscordGo(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMCoice == 0 {
		npm.DiscordNodejsNpm(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMCoice == 1 {
		yarn.DiscordNodejsYarn(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMCoice == 2 {
		pnpm.DiscordNodejsPnpm(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 {
		ruby.DiscordRuby(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 && m.PMCoice == 0 {
		cargo.DiscordRustCargo(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 && m.PMCoice == 1 {
		fleet.DiscordRustFleet(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 {
		deno.DiscordDeno(opts.BotName)
	}
}
