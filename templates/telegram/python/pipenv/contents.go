package pipenv

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("assets/pipenv.dockerfile", botName)
}
