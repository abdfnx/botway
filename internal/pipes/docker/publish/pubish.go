package publish

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/looker"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func DockerPublishImage() {
	dockerPath, derr := looker.LookPath("docker")

	if derr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		panic(constants.FAIL_FOREGROUND.Render(" docker is not installed"))
	}

	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	fmt.Println(messageStyle.Render("\n\n======= Start Publishing Your Bot Docker Image 🐳 ======\n\n"))

	if _, err := os.Stat(".botway.yaml"); err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		panic(constants.FAIL_FOREGROUND.Render(" You need to run this command in your bot directory"))
	}

	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	botImage := viper.GetString("docker.image")
	botPath := gjson.Get(string(constants.BotwayConfig), "botway.bots." + viper.GetString("bot.name") + ".path").String()
	dockerPublish := dockerPath + " publish " + botImage

	cmd := exec.Command("bash", "-c", dockerPublish)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", dockerPublish)
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
