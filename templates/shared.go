package templates

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/resto/core/api"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
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

	if strings.Contains(respone, "#include <{{.BotName}}/{{.BotName}}.h>") && strings.Contains(templateName, "telegram") {
		respone = strings.ReplaceAll(respone, "#include <{{.BotName}}/{{.BotName}}.h>", "")
	} else if strings.Contains(respone, `#include "botway/botway.hpp"`) && strings.Contains(templateName, "telegram") {
		respone = strings.ReplaceAll(respone, `#include "botway/botway.hpp"`, `#include "botway.hpp"`)
	} else if strings.Contains(arg, "pubspec.yaml") {
		respone = strings.ReplaceAll(respone, "{{.BotName}}", strings.ReplaceAll(botName, "-", ""))
	}

	respone = strings.ReplaceAll(respone, "{{.BotName}}", botName)

	viper.SetConfigType("json")

	viper.ReadConfig(bytes.NewBuffer(constants.BotwayConfig))

	respone = strings.ReplaceAll(respone, "{{.Author}}", viper.GetString("github.username"))

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
