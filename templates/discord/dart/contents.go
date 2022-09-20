package dart

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/dart.dockerfile", hostService), "botway", botName)
}

func Resources() string {
	return templates.Content("discord/dart.md", "resources", "")
}

func MainDartContent() string {
	return templates.Content("src/main.dart", "discord-dart", "")
}

func PubspecFileContent(botName string) string {
	return templates.Content("pubspec.yaml", "discord-dart", botName)
}
