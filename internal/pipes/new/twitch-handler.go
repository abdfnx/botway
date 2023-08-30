package new

import "github.com/botwayorg/templates"

func TwitchHandler(m model) {
	if m.PlatformChoice == 3 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "twitch", "pip")
	} else if m.PlatformChoice == 3 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "twitch", "pipenv")
	} else if m.PlatformChoice == 3 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "twitch", "poetry")
	} else if m.PlatformChoice == 3 && m.LangChoice == 1 {
		templates.GoTemplate(opts.BotName, "twitch")
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "twitch", false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "twitch", false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "twitch", false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "twitch", false)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "twitch", true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "twitch", true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "twitch", true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 3 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "twitch", true)
	} else if m.PlatformChoice == 3 && m.LangChoice == 4 {
		templates.DenoTemplate(opts.BotName, "twitch")
	} else if m.PlatformChoice == 3 && m.LangChoice == 5 {
		templates.JavaTemplate(opts.BotName, "twitch")
	}
}
