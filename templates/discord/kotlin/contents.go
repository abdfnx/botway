package kotlin

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/gradle.dockerfile", hostService), "botway", botName, "discord")
}

func Resources() string {
	return templates.Content("discord/kotlin.md", "resources", "", "")
}

func BotlinContent() string {
	return templates.Content("packages/botlin/main.kt", "botway", "", "")
}

func MainKtContent() string {
	return templates.Content("app/src/main/kotlin/core/Bot.kt", "discord-kotlin", "", "")
}

func BuildGradleKtsContent() string {
	return templates.Content("app/build.gradle.kts", "discord-kotlin", "", "")
}

func GradleWrapperPropsContent() string {
	return templates.Content("gradle/wrapper/gradle-wrapper.properties", "discord-kotlin", "", "")
}

func DotGitattributesContent() string {
	return templates.Content(".gitattributes", "discord-kotlin", "", "")
}

func GradlewContent() string {
	return templates.Content("gradlew", "discord-kotlin", "", "")
}

func GradlewBatContent() string {
	return templates.Content("gradlew.bat", "discord-kotlin", "", "")
}

func SettingsGradleKts() string {
	return templates.Content("settings.gradle.kts", "discord-kotlin", "", "")
}
