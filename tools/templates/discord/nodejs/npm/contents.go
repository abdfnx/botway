package npm

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "nodejs", "npm/Dockerfile", botName)
}
