package pnpm

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "nodejs", "pnpm/Dockerfile", botName)
}
