package pipenv

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("telegram", "python", "pipenv/Dockerfile", botName)
}
