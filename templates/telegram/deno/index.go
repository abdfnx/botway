package deno

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

func TelegramDeno(botName, hostService string) {
	deno, err := looker.LookPath("deno")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" deno is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "main.ts"), []byte(MainTsContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		denoInstall := deno + " cache main.ts"

		installCmd := exec.Command("bash", "-c", denoInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", denoInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// format files
		denoFormat := deno + " fmt"

		formatCmd := exec.Command("bash", "-c", denoFormat)

		if runtime.GOOS == "windows" {
			formatCmd = exec.Command("powershell.exe", denoFormat)
		}

		formatCmd.Dir = botName
		formatCmd.Stdin = os.Stdin
		formatCmd.Stdout = os.Stdout
		formatCmd.Stderr = os.Stderr
		err = formatCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
