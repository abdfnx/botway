package pipenv

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("slack", "python", "pipenv/Dockerfile", botName)
}
