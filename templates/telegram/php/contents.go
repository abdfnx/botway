package php

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/php.dockerfile", hostService), "botway", botName)
}

func MainPHPContent() string {
	return templates.Content("src/main.php", "telegram-php", "")
}

func BotwayPHPContent() string {
	return templates.Content("packages/bw-php/main.php", "botway", "")
}

func Resources() string {
	return templates.Content("telegram/php.md", "resources", "")
}

func ComposerFileContent(botName string) string {
	return templates.Content("composer.json", "telegram-php", botName)
}
