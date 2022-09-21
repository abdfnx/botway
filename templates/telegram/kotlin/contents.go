package kotlin

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/gradle.dockerfile", hostService), "botway", botName, "telegram")
}

func Resources() string {
	return templates.Content("telegram/kotlin.md", "resources", "", "")
}

func BotlinContent() string {
	return templates.Content("packages/botlin/main.kt", "botway", "", "")
}

func MainKtContent() string {
	return templates.Content("app/src/main/kotlin/core/Bot.kt", "telegram-kotlin", "", "")
}

func BuildGradleKtsContent() string {
	return templates.Content("app/build.gradle.kts", "telegram-kotlin", "", "")
}

func GradleWrapperPropsContent() string {
	return templates.Content("gradle/wrapper/gradle-wrapper.properties", "telegram-kotlin", "", "")
}

func DotGitattributesContent() string {
	return templates.Content(".gitattributes", "telegram-kotlin", "", "")
}

func GradlewContent() string {
	return templates.Content("gradlew", "telegram-kotlin", "", "")
}

func GradlewBatContent() string {
	return templates.Content("gradlew.bat", "telegram-kotlin", "", "")
}

func SettingsGradleKts() string {
	return templates.Content("settings.gradle.kts", "telegram-kotlin", "", "")
}
