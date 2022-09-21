package php

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/php.dockerfile", hostService), "botway", botName, "discord")
}

func MainPHPContent() string {
	return templates.Content("src/main.php", "discord-php", "", "")
}

func BotwayPHPContent() string {
	return templates.Content("packages/bw-php/main.php", "botway", "", "")
}

func Resources() string {
	return templates.Content("discord/php.md", "resources", "", "")
}

func ComposerFileContent(botName string) string {
	return templates.Content("composer.json", "discord-php", botName, "")
}
