package pnpm

import "github.com/abdfnx/botway/tools/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "nodejs", "pnpm/Dockerfile", botName)
}
