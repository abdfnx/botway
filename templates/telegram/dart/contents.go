package dart

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("dart.dockerfile", "botway/dockerfiles", botName)
}

func Resources() string {
	return templates.Content("telegram/dart.md", "resources", "")
}

func MainDartContent() string {
	return templates.Content("src/main.dart", "telegram-dart", "")
}

func PubspecFileContent(botName string) string {
	return templates.Content("pubspec.yaml", "telegram-dart", botName)
}
