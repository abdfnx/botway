package new

import (
	"github.com/abdfnx/botway/templates/nodejs"
	"github.com/abdfnx/botway/templates/telegram/cpp"
	"github.com/abdfnx/botway/templates/telegram/csharp"
	"github.com/abdfnx/botway/templates/telegram/dart"
	"github.com/abdfnx/botway/templates/telegram/deno"
	tgo "github.com/abdfnx/botway/templates/telegram/go"
	"github.com/abdfnx/botway/templates/telegram/java"
	"github.com/abdfnx/botway/templates/telegram/kotlin"
	"github.com/abdfnx/botway/templates/telegram/nim"
	"github.com/abdfnx/botway/templates/telegram/php"
	"github.com/abdfnx/botway/templates/telegram/python/pip"
	"github.com/abdfnx/botway/templates/telegram/python/pipenv"
	"github.com/abdfnx/botway/templates/telegram/python/poetry"
	"github.com/abdfnx/botway/templates/telegram/ruby"
	"github.com/abdfnx/botway/templates/telegram/rust"
	"github.com/abdfnx/botway/templates/telegram/swift"
	"github.com/abdfnx/botway/templates/ts"
)

func TelegramHandler(m model) {
	if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMChoice == 0 {
		pip.TelegramPythonPip(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMChoice == 1 {
		pipenv.TelegramPythonPipenv(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMChoice == 2 {
		poetry.TelegramPythonPoetry(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 1 {
		tgo.TelegramGo(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 0 {
		nodejs.Nodejs(opts.BotName, "npm", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 1 {
		nodejs.Nodejs(opts.BotName, "yarn", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 2 {
		nodejs.Nodejs(opts.BotName, "pnpm", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 3 {
		nodejs.Nodejs(opts.BotName, "bun", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 0 {
		ts.NodejsTS(opts.BotName, "npm", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 1 {
		ts.NodejsTS(opts.BotName, "yarn", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 2 {
		ts.NodejsTS(opts.BotName, "pnpm", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 3 {
		ts.NodejsTS(opts.BotName, "bun", "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 4 {
		ruby.TelegramRuby(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 5 && m.PMChoice == 0 {
		rust.TelegramRust(opts.BotName, "cargo")
	} else if m.PlatformChoice == 1 && m.LangChoice == 5 && m.PMChoice == 1 {
		rust.TelegramRust(opts.BotName, "fleet")
	} else if m.PlatformChoice == 1 && m.LangChoice == 6 {
		deno.TelegramDeno(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 7 {
		csharp.TelegramCsharp(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 8 {
		dart.TelegramDart(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 9 {
		php.TelegramPHP(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 10 {
		kotlin.TelegramKotlin(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 11 {
		java.TelegramJava(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 12 {
		cpp.TelegramCpp(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 13 {
		nim.TelegramNim(opts.BotName)
	} else if m.PlatformChoice == 1 && m.LangChoice == 14 {
		swift.TelegramSwift(opts.BotName)
	}
}
