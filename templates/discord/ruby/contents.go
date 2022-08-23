package ruby

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("dockerfiles/ruby.dockerfile", "botway", botName)
}

func MainRbContent() string {
	return templates.Content("main.rb", "discord-ruby", "")
}

func Resources() string {
	return templates.Content("discord/ruby.md", "resources", "")
}
