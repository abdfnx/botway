package python

import "github.com/abdfnx/botway/templates"

func MainPyContent() string {
	return templates.Content("main.py", "discord-python", "")
}

func Resources() string {
	return templates.Content("discord/python.md", "resources", "")
}
