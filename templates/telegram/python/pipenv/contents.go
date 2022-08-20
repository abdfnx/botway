package pipenv

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("pipenv.dockerfile", "botway/dockerfiles", botName)
}
