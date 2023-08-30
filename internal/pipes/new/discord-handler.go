package new

import "github.com/botwayorg/templates"

func DiscordHandler(m model) {
	if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "discord", "pip")
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "discord", "pipenv")
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "discord", "poetry")
	} else if m.PlatformChoice == 0 && m.LangChoice == 1 {
		templates.GoTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "discord", false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "discord", false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "discord", false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "discord", false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "discord", true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "discord", true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "discord", true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "discord", true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 {
		templates.RubyTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 && m.PMChoice == 0 {
		templates.RustTemplate(opts.BotName, "discord", "cargo")
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 && m.PMChoice == 1 {
		templates.RustTemplate(opts.BotName, "discord", "fleet")
	} else if m.PlatformChoice == 0 && m.LangChoice == 6 {
		templates.DenoTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 7 {
		templates.CsharpTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 8 {
		templates.DartTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 9 {
		templates.PHPTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 10 {
		templates.KotlinTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 11 {
		templates.JavaTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 12 {
		templates.CppTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 13 {
		templates.NimTemplate(opts.BotName, "discord")
	} else if m.PlatformChoice == 0 && m.LangChoice == 14 {
		templates.CTemplate(opts.BotName)
	} else if m.PlatformChoice == 0 && m.LangChoice == 15 {
		templates.CrystalTemplate(opts.BotName)
	}
}
