package pm

import (
	"github.com/abdfnx/botway/templates/slack/nodejs"
	"github.com/abdfnx/botway/templates/slack/python/pip"
	"github.com/abdfnx/botway/templates/slack/python/pipenv"
)

func SlackHandler(m model, botName, pm string) {
	if m.platform == "slack" && m.lang == "python" && pm == "pip" {
		pip.SlackPythonPip(botName)
	} else if m.platform == "slack" && m.lang == "python" && pm == "pipenv" {
		pipenv.SlackPythonPipenv(botName)
	} else if m.platform == "slack" && m.lang == "nodejs" {
		nodejs.SlackNodejs(botName, "npm")
	}
}
