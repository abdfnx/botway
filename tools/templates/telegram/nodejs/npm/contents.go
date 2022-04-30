package npm

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("telegram", "nodejs", "npm/Dockerfile", botName)
}
