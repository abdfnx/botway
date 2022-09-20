package swift

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/swift.dockerfile", hostService), "botway", botName)
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
