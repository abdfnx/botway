package dgo

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/go.dockerfile", hostService), "botway", botName)
}

func Resources() string {
	return templates.Content("discord/go.md", "resources", "")
}

func MainGoContent() string {
	return templates.Content("main.go", "discord-go", "")
}
