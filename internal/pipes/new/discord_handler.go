package new

import (
	"github.com/abdfnx/botway/templates/discord/c"
	"github.com/abdfnx/botway/templates/discord/cpp"
	"github.com/abdfnx/botway/templates/discord/crystal"
	"github.com/abdfnx/botway/templates/discord/csharp"
	"github.com/abdfnx/botway/templates/discord/dart"
	"github.com/abdfnx/botway/templates/discord/deno"
	dgo "github.com/abdfnx/botway/templates/discord/go"
	"github.com/abdfnx/botway/templates/discord/java"
	"github.com/abdfnx/botway/templates/discord/kotlin"
	"github.com/abdfnx/botway/templates/discord/nim"
	"github.com/abdfnx/botway/templates/discord/nodejs"
	"github.com/abdfnx/botway/templates/discord/php"
	"github.com/abdfnx/botway/templates/discord/python/pip"
	"github.com/abdfnx/botway/templates/discord/python/pipenv"
	"github.com/abdfnx/botway/templates/discord/python/poetry"
	"github.com/abdfnx/botway/templates/discord/ruby"
	"github.com/abdfnx/botway/templates/discord/rust"
)

func DiscordHandler(m model) {
	if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 0 {
		pip.DiscordPythonPip(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 1 {
		pipenv.DiscordPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 2 {
		poetry.DiscordPythonPoetry(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 1 {
		dgo.DiscordGo(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 0 {
		nodejs.DiscordNodejs(opts.BotName, "npm")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 1 {
		nodejs.DiscordNodejs(opts.BotName, "yarn")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 2 {
		nodejs.DiscordNodejs(opts.BotName, "pnpm")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 3 {
		nodejs.DiscordNodejs(opts.BotName, "bun")
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 {
		ruby.DiscordRuby(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 && m.PMChoice == 0 {
		rust.DiscordRust(opts.BotName, "cargo")
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 && m.PMChoice == 1 {
		rust.DiscordRust(opts.BotName, "fleet")
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 {
		deno.DiscordDeno(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 6 {
		csharp.DiscordCsharp(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 7 {
		dart.DiscordDart(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 8 {
		php.DiscordPHP(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 9 {
		kotlin.DiscordKotlin(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 10 {
		java.DiscordJava(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 11 {
		cpp.DiscordCpp(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 12 {
		nim.DiscordNim(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 13 {
		c.DiscordC(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 14 {
		crystal.DiscordCrystal(opts.BotName)
	}
}
