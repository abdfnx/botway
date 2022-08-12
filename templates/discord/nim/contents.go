package nim

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("nim.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("discord/nim.md", "resources", "")
}

func MainNimContent() string {
	return templates.Content("src/main.nim", "discord-nim", "")
}

func BotnimContent(botName string) string {
	return templates.Content("packages/botnim/main.nim", "botway", botName)
}

func PngFileContent() string {
	return templates.Content("assets/facepalm.png", "discord-nim", "")
}

func NimbleFileContent() string {
	return templates.Content("discord_nim.nimble", "discord-nim", "")
}
