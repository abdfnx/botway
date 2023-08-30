package new

import "github.com/botwayorg/templates"

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "slack", "pip")
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "slack", "pipenv")
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "slack", "poetry")
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "slack", false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "slack", false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "slack", false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "slack", false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "slack", true)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "slack", true)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "slack", true)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "slack", true)
	}
}
