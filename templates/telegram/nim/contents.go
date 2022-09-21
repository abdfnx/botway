package nim

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/nim.dockerfile", hostService), "botway", botName, "telegram")
}

func Resources() string {
	return templates.Content("telegram/nim.md", "resources", "", "")
}

func MainNimContent() string {
	return templates.Content("src/main.nim", "telegram-nim", "", "")
}

func BotnimContent(botName string) string {
	return templates.Content("packages/botnim/main.nim", "botway", botName, "")
}

func NimbleFileContent() string {
	return templates.Content("telegram_nim.nimble", "telegram-nim", "", "")
}
