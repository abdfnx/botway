package crystal

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("crystal.dockerfile", "botway/dockerfiles", botName)
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
