package tgo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools/templates"
	"github.com/abdfnx/looker"
)

func TelegramGo(botName string) {
	goPath, err := looker.LookPath("go")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" go is not installed"))
	} else {
		goInit := goPath + " mod init " + botName

		cmd := exec.Command("bash", "-c", goInit)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", goInit)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.go"), []byte(MainGoContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: ./" + botName), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if procFile != nil {
			log.Fatal(procFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		goTidy := goPath + " mod tidy"

		tidyCmd := exec.Command("bash", "-c", goTidy)

		if runtime.GOOS == "windows" {
			tidyCmd = exec.Command("powershell.exe", goTidy)
		}

		tidyCmd.Dir = botName
		tidyCmd.Stdin = os.Stdin
		tidyCmd.Stdout = os.Stdout
		tidyCmd.Stderr = os.Stderr
		err = tidyCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName)
	}
}
