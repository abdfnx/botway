package python

import "github.com/abdfnx/botway/templates"

func MainPyContent() string {
	return templates.Content("main.py", "slack-python", "", "")
}

func Resources() string {
	return templates.Content("slack/python.md", "resources", "", "")
}
