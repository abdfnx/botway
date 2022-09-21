package ruby

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/ruby.dockerfile", hostService), "botway", botName, "telegram")
}

func MainRbContent() string {
	return templates.Content("main.rb", "telegram-ruby", "", "")
}

func Resources() string {
	return templates.Content("telegram/ruby.md", "resources", "", "")
}
