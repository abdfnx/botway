package nim

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("nim.dockerfile", "botway/dockerfiles", botName)
}

func Resources() string {
	return templates.Content("telegram/nim.md", "resources", "")
}

func MainNimContent() string {
	return templates.Content("src/main.nim", "telegram-nim", "")
}

func BotnimContent(botName string) string {
	return templates.Content("packages/botnim/main.nim", "botway", botName)
}

func NimbleFileContent() string {
	return templates.Content("telegram_nim.nimble", "telegram-nim", "")
}
