package php

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

func TelegramPHP(botName, hostService string) {
	_, err := looker.LookPath("php")
	composerPath, serr := looker.LookPath("composer")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" php is not installed"))
	} else if serr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" composer is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.php"), []byte(MainPHPContent()), 0644)
		botwayFile := os.WriteFile(filepath.Join(botName, "src", "botway.php"), []byte(BotwayPHPContent()), 0644)
		composerFile := os.WriteFile(filepath.Join(botName, "composer.json"), []byte(ComposerFileContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if botwayFile != nil {
			log.Fatal(botwayFile)
		}

		if composerFile != nil {
			log.Fatal(composerFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		composerInstall := composerPath + " install"

		installCmd := exec.Command("bash", "-c", composerInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", composerInstall)
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
