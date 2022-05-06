package pip

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "python", "pip/Dockerfile", botName)
}

func RequirementsContent() string {
	return templates.Content("discord", "python", "pip/requirements.txt", "")
}
