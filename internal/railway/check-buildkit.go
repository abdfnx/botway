package railway

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func CheckBuildKit() {
	tools.CheckDir()

	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	checkBuildKit := viper.GetBool("docker.enable_buildkit")

	setVarCmd := "botway vars set DOCKER_BUILDKIT=0"

	if checkBuildKit {
		setVarCmd = "botway vars set DOCKER_BUILDKIT=1"
	}

	botPath := gjson.Get(string(constants.BotwayConfig), "botway.bots."+viper.GetString("bot.name")+".path").String()

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
