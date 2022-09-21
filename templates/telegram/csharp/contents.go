package csharp

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/csharp.dockerfile", hostService), "botway", botName, "telegram")
}

func Resources() string {
	return templates.Content("telegram/csharp.md", "resources", "", "")
}

func MainCsContent() string {
	return templates.Content("src/Main.cs", "telegram-csharp", "", "")
}

func BotCSharpProj() string {
	return templates.Content("telegram-csharp.csproj", "telegram-csharp", "", "")
}
