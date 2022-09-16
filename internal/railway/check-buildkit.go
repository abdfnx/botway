package railway

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/internal/config"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/viper"
)

func CheckBuildKit() {
	tools.CheckDir()

	checkBuildKit := botwaygo.GetBotInfo("docker.enable_buildkit")

	setVarCmd := "botway vars set --no-redeploy-hint DOCKER_BUILDKIT=0"

	if checkBuildKit == "true" {
		setVarCmd = "botway vars set --no-redeploy-hint DOCKER_BUILDKIT=1"
	}

	botPath := config.Get("botway.bots." + viper.GetString("bot.name") + ".path")

	cmd := exec.Command("bash", "-c", setVarCmd)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", setVarCmd)
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
