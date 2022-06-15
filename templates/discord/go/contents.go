package dgo

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("go.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("discord/go.md", "resources", "")
}

func MainGoContent() string {
	return templates.Content("main.go", "discord-go", "")
}
