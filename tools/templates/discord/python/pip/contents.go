package pip

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "python", "pip/Dockerfile", botName)
}
