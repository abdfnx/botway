package pip

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("pip.dockerfile", "botway/dockerfiles", botName)
}

func RequirementsContent() string {
	return templates.Content("requirements.txt", "telegram-python", "")
}
