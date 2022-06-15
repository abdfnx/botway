package ruby

import "github.com/abdfnx/botway/templates"

func MainRbContent() string {
	return templates.Content("main.rb", "telegram-ruby", "")
}

func DockerfileContent(botName string) string {
	return templates.Content("ruby.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("telegram/ruby.md", "resources", "")
}
