package build_image

import (
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

func DockerBuildImage() {
	dockerPath, derr := looker.LookPath("docker")

	if derr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		panic(constants.FAIL_FOREGROUND.Render(" docker is not installed"))
	}

	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	fmt.Println(messageStyle.Render("\n\n======= Start Building Your Bot Docker Image 🐳 ======\n\n"))


	if _, err := os.Stat(".botway.yaml"); err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		panic(constants.FAIL_FOREGROUND.Render("You need to run this command in your bot directory"))
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".botway")
	viper.SetConfigType("yaml")

	buildCmd := viper.GetString("docker.build_cmd")
	botPath := gjson.Get(string(constants.BotwayConfig), "botway.bots." + viper.GetString("bot.name") + ".path").String()
	dockerBuild := dockerPath + " " + buildCmd

	cmd := exec.Command("bash", "-c", dockerBuild)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", dockerBuild)
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
