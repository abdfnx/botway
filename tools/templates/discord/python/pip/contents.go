package pip

import "github.com/abdfnx/botway/tools/templates/discord/python"

func DockerfileContent(botName string) string {
	return python.Content("pip/Dockerfile", botName)
}
