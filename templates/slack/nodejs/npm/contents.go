package npm

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("slack", "nodejs", "npm/Dockerfile", botName)
}
