package new

import "github.com/botwayorg/templates"

func TwitchHandler(m model) {
	if m.PlatformChoice == 3 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "twitch", "pip", HostServiceName(m))
	} else if m.PlatformChoice == 3 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "twitch", "pipenv", HostServiceName(m))
	} else if m.PlatformChoice == 3 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "twitch", "poetry", HostServiceName(m))
	} else if m.PlatformChoice == 3 && m.LangChoice == 1 {
		templates.GoTemplate(opts.BotName, "twitch", HostServiceName(m))
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "twitch", HostServiceName(m), false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "twitch", HostServiceName(m), false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "twitch", HostServiceName(m), false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "twitch", HostServiceName(m), false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "twitch", HostServiceName(m), true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "twitch", HostServiceName(m), true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "twitch", HostServiceName(m), true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "twitch", HostServiceName(m), true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 4 {
		templates.DenoTemplate(opts.BotName, "twitch", HostServiceName(m))
	} else if m.PlatformChoice == 3 && m.LangChoice == 5 {
		templates.JavaTemplate(opts.BotName, "twitch", HostServiceName(m))
	}
}
