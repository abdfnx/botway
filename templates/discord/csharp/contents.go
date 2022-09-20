package csharp

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/crystal.dockerfile", hostService), "botway", botName)
}

func Resources() string {
	return templates.Content("discord/csharp.md", "resources", "")
}

func MainCsContent() string {
	return templates.Content("src/Main.cs", "discord-csharp", "")
}

func BotCSharpProj() string {
	return templates.Content("discord-csharp.csproj", "discord-csharp", "")
}
