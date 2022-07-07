package csharp

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("csharp.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("telegram/csharp.md", "resources", "")
}

func MainCsContent() string {
	return templates.Content("src/Main.cs", "telegram-csharp", "")
}

func BotCSharpProj() string {
	return templates.Content("telegram-csharp.csproj", "telegram-csharp", "")
}
