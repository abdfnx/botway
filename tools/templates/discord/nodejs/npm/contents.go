package npm

import "github.com/abdfnx/botway/tools/templates/discord/nodejs"

func DockerfileContent(botName string) string {
	return nodejs.Content("npm/Dockerfile", botName)
}
