package templates

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/config"
	"github.com/abdfnx/resto/core/api"
	"github.com/charmbracelet/lipgloss"
)

func Content(arg, templateName, botName string) string {
	org := "botwayorg"

	if templateName == "botway" {
		org = "abdfnx"
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", org, templateName, arg)
	respone, status, _, err := api.BasicGet(url, "GET", "", "", "", "", false, 0, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	if status == "404" || status == "401" || strings.Contains(respone, "404") {
		fmt.Println("404: " + url)
		os.Exit(0)
	}

	// TODO: fix botway c++ telegram template
	if strings.Contains(respone, "#include <{{.BotName}}/{{.BotName}}.h>") && strings.Contains(templateName, "telegram") {
		respone = strings.ReplaceAll(respone, "#include <{{.BotName}}/{{.BotName}}.h>", "")
	} else if strings.Contains(respone, `#include "botway/botway.hpp"`) && strings.Contains(templateName, "telegram") {
		respone = strings.ReplaceAll(respone, `#include "botway/botway.hpp"`, `#include "botway.hpp"`)
	} else if strings.Contains(arg, "pubspec.yaml") {
		respone = strings.ReplaceAll(respone, "{{.BotName}}", strings.ReplaceAll(botName, "-", ""))
	}

	respone = strings.ReplaceAll(respone, "{{.BotName}}", botName)

	author := config.Get("github.username")

	if author == "" {
		author = "botway"
	}

	respone = strings.ReplaceAll(respone, "{{.Author}}", author)

	return respone
}

func CheckProject(botName, botType string) {
	if _, err := os.Stat(botName); !os.IsNotExist(err) {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + botName + " Created successfully 🎉"))
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("NEXT"))
		fmt.Println(" Now, run " + lipgloss.NewStyle().Foreground(constants.GRAY_COLOR).Render("botway tokens set --"+botType+" "+botName) + " command to add tokens of your bot 🔑")
	}
}
