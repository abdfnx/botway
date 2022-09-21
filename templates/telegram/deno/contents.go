package deno

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/deno.dockerfile", hostService), "botway", botName, "telegram")
}

func Resources() string {
	return templates.Content("telegram/deno.md", "resources", "", "")
}

func MainTsContent() string {
	return templates.Content("main.ts", "telegram-deno", "", "")
}
