package new

import "github.com/botwayorg/templates"

func SlackHandler(m model) {
	if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 0 {
		templates.PythonTemplate(opts.BotName, "slack", "pip", HostServiceName(m))
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 1 {
		templates.PythonTemplate(opts.BotName, "slack", "pipenv", HostServiceName(m))
	} else if m.PlatformChoice == 2 && m.LangChoice == 0 && m.PMChoice == 2 {
		templates.PythonTemplate(opts.BotName, "slack", "poetry", HostServiceName(m))
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "slack", HostServiceName(m), false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "slack", HostServiceName(m), false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "slack", HostServiceName(m), false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 1 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "slack", HostServiceName(m), false)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 0 {
		templates.NodejsTemplate(opts.BotName, "npm", "slack", HostServiceName(m), true)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 1 {
		templates.NodejsTemplate(opts.BotName, "yarn", "slack", HostServiceName(m), true)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 2 {
		templates.NodejsTemplate(opts.BotName, "pnpm", "slack", HostServiceName(m), true)
	} else if m.PlatformChoice == 2 && m.LangChoice == 2 && m.PMChoice == 3 {
		templates.NodejsTemplate(opts.BotName, "bun", "slack", HostServiceName(m), true)
	}
}
