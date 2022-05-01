package pip

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("telegram", "python", "pip/Dockerfile", botName)
}

func RequirementsContent() string {
	return templates.Content("telegram", "python", "pip/requirements.txt", "")
}
