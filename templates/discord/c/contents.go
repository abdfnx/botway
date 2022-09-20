package c

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/c-discord.dockerfile", hostService), "botway", botName)
}

func Resources() string {
	return templates.Content("discord/c.md", "resources", "")
}

func MainCContent() string {
	return templates.Content("src/main.c", "discord-c", "")
}

func BWCContent(botName string) string {
	return templates.Content("packages/bwc/main.h", "botway", botName)
}

func RunPsFileContent() string {
	return templates.Content("run.ps1", "discord-c", "")
}
