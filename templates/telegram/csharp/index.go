package csharp

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

func TelegramCsharp(botName string) {
	dotnetPath, err := looker.LookPath("dotnet")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" dotnet is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "Main.cs"), []byte(MainCsContent()), 0644)
		csprojFile := os.WriteFile(filepath.Join(botName, botName+".csproj"), []byte(BotCSharpProj()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if csprojFile != nil {
			log.Fatal(csprojFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		dotNetRestore := dotnetPath + " restore"

		restoreCmd := exec.Command("bash", "-c", dotNetRestore)

		if runtime.GOOS == "windows" {
			restoreCmd = exec.Command("powershell.exe", dotNetRestore)
		}

		restoreCmd.Dir = botName
		restoreCmd.Stdin = os.Stdin
		restoreCmd.Stdout = os.Stdout
		restoreCmd.Stderr = os.Stderr
		err = restoreCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
