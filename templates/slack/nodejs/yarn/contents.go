package yarn

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("slack", "nodejs", "yarn/Dockerfile", botName)
}
