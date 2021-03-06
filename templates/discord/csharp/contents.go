package csharp

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("csharp.dockerfile", "dockerfiles", botName)
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
