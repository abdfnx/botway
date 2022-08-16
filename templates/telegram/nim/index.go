package nim

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/templates"
	"github.com/abdfnx/looker"
)

func TelegramNim(botName string) {
	_, err := looker.LookPath("nim")
	nimblePath, nerr := looker.LookPath("nimble")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" nim is not installed"))
	} else if nerr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" nimble is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.nim"), []byte(MainNimContent()), 0644)
		botnimFile := os.WriteFile(filepath.Join(botName, "src", "botnim.nim"), []byte(BotnimContent(botName)), 0644)
		nimbleFile := os.WriteFile(filepath.Join(botName, botName + ".nimble"), []byte(NimbleFileContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if botnimFile != nil {
			log.Fatal(botnimFile)
		}

		if nimbleFile != nil {
			log.Fatal(nimbleFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		nimbleInstall := nimblePath + " install -y"

		installCmd := exec.Command("bash", "-c", nimbleInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", nimbleInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
