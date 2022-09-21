package tgo

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/go.dockerfile", hostService), "botway", botName, "telegram")
}

func MainGoContent() string {
	return templates.Content("main.go", "telegram-go", "", "")
}

func Resources() string {
	return templates.Content("telegram/go.md", "resources", "", "")
}
