package pipenv

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "python", "pipenv/Dockerfile", botName)
}
