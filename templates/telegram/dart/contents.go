package dart

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/dart.dockerfile", hostService), "botway", botName, "telegram")
}

func Resources() string {
	return templates.Content("telegram/dart.md", "resources", "", "")
}

func MainDartContent() string {
	return templates.Content("src/main.dart", "telegram-dart", "", "")
}

func PubspecFileContent(botName string) string {
	return templates.Content("pubspec.yaml", "telegram-dart", botName, "")
}
