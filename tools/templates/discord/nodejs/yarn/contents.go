package yarn

import "github.com/abdfnx/botway/tools/templates/discord/nodejs"

func DockerfileContent(botName string) string {
	return nodejs.Content("yarn/Dockerfile", botName)
}
