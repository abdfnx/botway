package swift

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("swift.dockerfile", "botway/dockerfiles", botName)
}

func Resources() string {
	return templates.Content("telegram/swift.md", "resources", "")
}

func MainSwiftContent() string {
	return templates.Content("Sources/bwbot/main.swift", "telegram-swift", "")
}

func BotwaySwiftContent(botName string) string {
	return templates.Content("packages/botway-swift/main.swift", "botway", botName)
}

func PackageSwiftFileContent(botName string) string {
	return templates.Content("Package.swift", "telegram-swift", botName)
}
