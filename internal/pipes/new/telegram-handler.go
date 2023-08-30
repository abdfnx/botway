package new

import "github.com/botwayorg/templates"

func TelegramHandler(m model) {
	if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "telegram", "pip")
	} else if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "telegram", "pipenv")
	} else if m.PlatformChoice == 1 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "telegram", "poetry")
	} else if m.PlatformChoice == 1 && m.LangChoice == 1 {
		templates.GoTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "telegram", false)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "telegram", false)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "telegram", false)
	} else if m.PlatformChoice == 1 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "telegram", false)
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "telegram", true)
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "telegram", true)
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "telegram", true)
	} else if m.PlatformChoice == 1 && m.LangChoice == 3 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "telegram", true)
	} else if m.PlatformChoice == 1 && m.LangChoice == 4 {
		templates.RubyTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 5 && m.PMChoice == 0 {
		templates.RustTemplate(opts.BotName, "telegram", "cargo")
	} else if m.PlatformChoice == 1 && m.LangChoice == 5 && m.PMChoice == 1 {
		templates.RustTemplate(opts.BotName, "telegram", "fleet")
	} else if m.PlatformChoice == 1 && m.LangChoice == 6 {
		templates.DenoTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 7 {
		templates.CsharpTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 8 {
		templates.DartTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 9 {
		templates.PHPTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 10 {
		templates.KotlinTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 11 {
		templates.JavaTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 12 {
		templates.CppTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 13 {
		templates.NimTemplate(opts.BotName, "telegram")
	} else if m.PlatformChoice == 1 && m.LangChoice == 14 {
		templates.SwiftTemplate(opts.BotName, "telegram")
	}
}
