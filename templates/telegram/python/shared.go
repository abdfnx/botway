package python

import "github.com/abdfnx/botway/templates"

func MainPyContent() string {
	return templates.Content("main.py", "telegram-python", "", "")
}

func Resources() string {
	return templates.Content("telegram/python.md", "resources", "", "")
}
