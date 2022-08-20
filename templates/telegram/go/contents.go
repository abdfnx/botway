package tgo

import (
	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName string) string {
	return templates.Content("go.dockerfile", "botway/dockerfiles", botName)
}

func MainGoContent() string {
	return templates.Content("main.go", "telegram-go", "")
}

func Resources() string {
	return templates.Content("telegram/go.md", "resources", "")
}
