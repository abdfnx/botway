package crystal

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/crystal.dockerfile", hostService), "botway", botName)
}

func Resources() string {
	return templates.Content("discord/crystal.md", "resources", "")
}

func MainCrContent() string {
	return templates.Content("src/main.cr", "discord-crystal", "")
}

func ShardFileContent(botName string) string {
	return templates.Content("shard.yml", "discord-crystal", botName)
}
