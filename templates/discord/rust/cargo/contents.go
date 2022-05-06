package cargo

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "rust", "cargo/Dockerfile", botName)
}
