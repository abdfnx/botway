package pip

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("pip.dockerfile", "dockerfiles", botName)
}

func RequirementsContent() string {
	return templates.Content("requirements.txt", "discord-python", "")
}
