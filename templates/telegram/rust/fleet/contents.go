package fleet

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("telegram", "rust", "fleet/Dockerfile", botName)
}
