package ruby

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/ruby.dockerfile", hostService), "botway", botName)
}

func MainRbContent() string {
	return templates.Content("main.rb", "discord-ruby", "")
}

func Resources() string {
	return templates.Content("discord/ruby.md", "resources", "")
}
