package pipenv

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("dockerfiles/pipenv.dockerfile", "botway", botName)
}
