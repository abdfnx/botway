package fleet

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "rust", "fleet/Dockerfile", botName)
}
