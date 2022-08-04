package cpp

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("cmake-discord.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("discord/cpp.md", "resources", "")
}

func FindDppCmakeContent() string {
	return templates.Content("cmake/FindDPP.cmake", "discord-cpp", "")
}

func BWCPPFileContent(botName string) string {
	return templates.Content("include/botway/botway.h", "discord-cpp", botName)
}

func MainIncludeFileContent() string {
	return templates.Content("include/bwbot/bwbot.h", "discord-cpp", "")
}

func MainCppContent(botName string) string {
	return templates.Content("src/main.cpp", "discord-cpp", botName)
}

func DotDockerIgnoreContent() string {
	return templates.Content(".dockerignore", "discord-cpp", "")
}

func CmakeListsContent(botName string) string {
	return templates.Content("CMakeLists.txt", "discord-cpp", botName)
}

func RunPsFileContent() string {
	return templates.Content("run.ps1", "discord-cpp", "")
}
