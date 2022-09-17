package compose

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
)

func Compose(justBuild bool, listMode bool) {
	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotComposeConfig))

	bots := viper.GetStringMapString("bots")

	for bot := range bots {
		b := "bots." + bot

		botPath := viper.GetString(b + ".path")

		if !listMode {
			fmt.Println(constants.HEADING + "Change directory to " + constants.BOLD.Render(botPath))

			if viper.GetBool(b+".run_in_container") == true {
				image := viper.GetString(b + ".container.image")
				ports := viper.GetString(b + ".container.ports")
				removeAfterRunValue := viper.GetBool(b + ".container.remove_after_run")

				portsFlag := ""

				if ports != "" {
					portsFlag = "-p " + ports
				}

				removeAfterRun := ""

				if removeAfterRunValue {
					removeAfterRun = "docker rmi " + image
				}

				dockerRunCmd := fmt.Sprintf(`
					docker run %s %s
					%s
				`, portsFlag, image, removeAfterRun)

				if justBuild {
					dockerRunCmd = ""
				}

				createContainerCmd := fmt.Sprintf(`
					botway docker build
					%s
				`, dockerRunCmd)

				cmd := exec.Command("bash", "-c", createContainerCmd)

				if runtime.GOOS == "windows" {
					cmd = exec.Command("powershell.exe", createContainerCmd)
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

			if !justBuild {
				if viper.GetBool(b+".run_local") == true {
					startCmd := ""

					if viper.GetString(b+".local.cmd") == "default" {
						startCmd = "botway start"
					} else {
						startCmd = viper.GetString(b + ".local.cmd")
					}

					cmd := exec.Command("bash", "-c", startCmd)

					if runtime.GOOS == "windows" {
						cmd = exec.Command("powershell.exe", startCmd)
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

				if viper.GetBool(b+".deploy") == true {
					deployCmd := "botway deploy"

					cmd := exec.Command("bash", "-c", deployCmd)

					if runtime.GOOS == "windows" {
						cmd = exec.Command("powershell.exe", deployCmd)
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
			}
		} else {
			fmt.Println(constants.HEADING + constants.BOLD.Render(bot))
		}
	}
}
