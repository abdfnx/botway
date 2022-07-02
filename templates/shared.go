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
	url := fmt.Sprintf("https://raw.githubusercontent.com/botwayorg/%s/main/%s", templateName, arg)
	respone, status, _, err := api.BasicGet(url, "GET", "", "", "", "", false, 0, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	if status == "404" || status == "401" || strings.Contains(respone, "404") {
		fmt.Println("404")
		fmt.Println(respone)
		os.Exit(0)
	}

	if strings.Contains(arg, ".dockerfile") || strings.Contains(arg, "Cargo.toml") || strings.Contains(arg, "shard.yml") || strings.Contains(arg, "pubspec.yaml") || strings.Contains(arg, "pyproject.toml") {
		respone = strings.ReplaceAll(respone, "{{.BotName}}", botName)

		viper.SetConfigType("json")

		viper.ReadConfig(bytes.NewBuffer(constants.BotwayConfig))

		respone = strings.ReplaceAll(respone, "{{.Author}}", viper.GetString("user.github_username"))
	}

	return respone
}

func CheckProject(botName, botType string) {
	if _, err := os.Stat(botName); !os.IsNotExist(err) {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + botName + " Created successfully 🎉"))
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("NEXT"))
		fmt.Println(" Now, run " + lipgloss.NewStyle().Foreground(constants.GRAY_COLOR).Render("botway tokens set --" + botType + " " + botName) + " command to add tokens of your bot 🔑")
	}
}
