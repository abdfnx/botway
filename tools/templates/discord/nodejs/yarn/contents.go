package yarn

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "nodejs", "yarn/Dockerfile", botName)
}
