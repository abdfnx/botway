package java

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/gradle.dockerfile", hostService), "botway", botName, "discord")
}

func Resources() string {
	return templates.Content("discord/java.md", "resources", "", "")
}

func BotlinContent() string {
	return templates.Content("packages/botlin/main.kt", "botway", "", "")
}

func MainJavaContent() string {
	return templates.Content("app/src/main/java/core/Bot.java", "discord-java", "", "")
}

func BuildGradleContent() string {
	return templates.Content("app/build.gradle", "discord-java", "", "")
}

func GradleWrapperPropsContent() string {
	return templates.Content("gradle/wrapper/gradle-wrapper.properties", "discord-java", "", "")
}

func DotGitattributesContent() string {
	return templates.Content(".gitattributes", "discord-java", "", "")
}

func GradlewContent() string {
	return templates.Content("gradlew", "discord-java", "", "")
}

func GradlewBatContent() string {
	return templates.Content("gradlew.bat", "discord-java", "", "")
}

func SettingsGradle() string {
	return templates.Content("settings.gradle", "discord-java", "", "")
}
