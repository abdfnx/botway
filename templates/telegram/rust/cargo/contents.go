package cargo

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("telegram", "rust", "cargo/Dockerfile", botName)
}