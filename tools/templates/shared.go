package templates

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/resto/core/api"
)

func Content(platform, lang, fileName, botName string) string {
	url := fmt.Sprintf("https://raw.githubusercontent.com/abdfnx/botway/main/tools/templates/%s/%s/assets/%s", platform, lang, fileName)
	respone, status, _, err := api.BasicGet(url, "GET", "", "", "", "", false, 0, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	if status == "404" || status == "401" || strings.Contains(respone, "404") {
		fmt.Println("404")
		os.Exit(0)
	}

	if strings.Contains(fileName, "Dockerfile") || strings.Contains(fileName, "Cargo.toml") {
		return strings.ReplaceAll(respone, "{{.BotName}}", botName)
	} else {
		return respone
	}
}
