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
	"github.com/abdfnx/botway/templates/discord/php"
	"github.com/abdfnx/botway/templates/discord/python/pip"
	"github.com/abdfnx/botway/templates/discord/python/pipenv"
	"github.com/abdfnx/botway/templates/discord/python/poetry"
	"github.com/abdfnx/botway/templates/discord/ruby"
	"github.com/abdfnx/botway/templates/discord/rust"
	"github.com/abdfnx/botway/templates/nodejs"
	"github.com/abdfnx/botway/templates/ts"
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
		nodejs.Nodejs(opts.BotName, "npm", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 1 {
		nodejs.Nodejs(opts.BotName, "yarn", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 2 {
		nodejs.Nodejs(opts.BotName, "pnpm", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 3 {
		nodejs.Nodejs(opts.BotName, "bun", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 0 {
		ts.NodejsTS(opts.BotName, "npm", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 1 {
		ts.NodejsTS(opts.BotName, "yarn", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 2 {
		ts.NodejsTS(opts.BotName, "pnpm", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 3 {
		ts.NodejsTS(opts.BotName, "bun", "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 {
		ruby.DiscordRuby(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 && m.PMChoice == 0 {
		rust.DiscordRust(opts.BotName, "cargo")
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 && m.PMChoice == 1 {
		rust.DiscordRust(opts.BotName, "fleet")
	} else if m.PlatformChoice == 0 && m.LangChoice == 6 {
		deno.DiscordDeno(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 7 {
		csharp.DiscordCsharp(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 8 {
		dart.DiscordDart(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 9 {
		php.DiscordPHP(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 10 {
		kotlin.DiscordKotlin(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 11 {
		java.DiscordJava(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 12 {
		cpp.DiscordCpp(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 13 {
		nim.DiscordNim(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 14 {
		c.DiscordC(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 15 {
		crystal.DiscordCrystal(opts.BotName)
	}
}
