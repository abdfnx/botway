package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/botwaygo"
	"github.com/abdfnx/looker"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func DockerRunImage() {
	_, derr := looker.LookPath("docker")

	if derr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		panic(constants.FAIL_FOREGROUND.Render(" docker is not installed"))
	}

	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	fmt.Println(messageStyle.Render("\n\n======= Start Running Your Bot Docker Image 🐳 ======\n\n"))

	tools.CheckDir()

	runCmd := botwaygo.GetBotInfo("docker.cmds.run")
	botPath := gjson.Get(string(constants.BotwayConfig), "botway.bots."+viper.GetString("bot.name")+".path").String()

	cmd := exec.Command("bash", "-c", runCmd)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", runCmd)
	}

	cmd.Dir = botPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
