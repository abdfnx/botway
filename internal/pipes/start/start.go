package start

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func Start() {
	if _, err := os.Stat(".botway.yaml"); err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		panic(constants.FAIL_FOREGROUND.Render("You need to run this command in your bot directory"))
	}

	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	fmt.Println(messageStyle.Render("\n\n======= Starting Your Bot 🤖 ======\n\n"))

	startCmd := viper.GetString("bot.start_cmd")
	botPath := gjson.Get(string(constants.BotwayConfig), "botway.bots." + viper.GetString("bot.name") + ".path").String()

	cmd := exec.Command("bash", "-c", startCmd)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", startCmd)
	}

	cmd.Dir = botPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmdErr := cmd.Run()

	if cmdErr != nil {
		log.Printf("error: %v\n", cmdErr)
	}
}
