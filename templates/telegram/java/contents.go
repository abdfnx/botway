package java

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/gradle.dockerfile", hostService), "botway", botName)
}

func Resources() string {
	return templates.Content("telegram/java.md", "resources", "")
}

func BotlinContent() string {
	return templates.Content("packages/botlin/main.kt", "botway", "")
}

func MainJavaContent() string {
	return templates.Content("app/src/main/java/core/Bot.java", "telegram-java", "")
}

func BotHandlerContent() string {
	return templates.Content("app/src/main/java/core/BotHandler.java", "telegram-java", "")
}

func TGBotContent() string {
	return templates.Content("app/src/main/java/core/TGBot.java", "telegram-java", "")
}

func BuildGradleContent() string {
	return templates.Content("app/build.gradle", "telegram-java", "")
}

func GradleWrapperPropsContent() string {
	return templates.Content("gradle/wrapper/gradle-wrapper.properties", "telegram-java", "")
}

func DotGitattributesContent() string {
	return templates.Content(".gitattributes", "telegram-java", "")
}

func GradlewContent() string {
	return templates.Content("gradlew", "telegram-java", "")
}

func GradlewBatContent() string {
	return templates.Content("gradlew.bat", "telegram-java", "")
}

func SettingsGradle() string {
	return templates.Content("settings.gradle", "telegram-java", "")
}
