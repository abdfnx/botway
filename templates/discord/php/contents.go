package php

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("php.dockerfile", "dockerfiles", botName)
}

func MainPHPContent() string {
	return templates.Content("src/main.php", "discord-php", "")
}

func BotwayPHPContent() string {
	return templates.Content("src/botway.php", "discord-php", "")
}

func Resources() string {
	return templates.Content("discord/php.md", "resources", "")
}

func ComposerFileContent(botName string) string {
	return templates.Content("composer.json", "discord-php", botName)
}
