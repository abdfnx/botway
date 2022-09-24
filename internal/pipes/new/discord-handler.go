package new

import "github.com/botwayorg/templates"

func DiscordHandler(m model) {
	if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "discord", "pip", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "discord", "pipenv", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "discord", "poetry", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 1 {
		templates.GoTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "discord", HostServiceName(m), false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "discord", HostServiceName(m), false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "discord", HostServiceName(m), false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "discord", HostServiceName(m), false)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "discord", HostServiceName(m), true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "discord", HostServiceName(m), true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "discord", HostServiceName(m), true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 3 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "discord", HostServiceName(m), true)
	} else if m.PlatformChoice == 0 && m.LangChoice == 4 {
		templates.RubyTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 && m.PMChoice == 0 {
		templates.RustTemplate(opts.BotName, "discord", "cargo", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 5 && m.PMChoice == 1 {
		templates.RustTemplate(opts.BotName, "discord", "fleet", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 6 {
		templates.DenoTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 7 {
		templates.CsharpTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 8 {
		templates.DartTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 9 {
		templates.PHPTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 10 {
		templates.KotlinTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 11 {
		templates.JavaTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 12 {
		templates.CppTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 13 {
		templates.NimTemplate(opts.BotName, "discord", HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 14 {
		templates.CTemplate(opts.BotName, HostServiceName(m))
	} else if m.PlatformChoice == 0 && m.LangChoice == 15 {
		templates.CrystalTemplate(opts.BotName, HostServiceName(m))
	}
}
