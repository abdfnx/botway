package ruby

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

func TelegramRuby(botName, hostService string) {
	_, err := looker.LookPath("ruby")
	bundlePath, berr := looker.LookPath("bundle")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" ruby is not installed"))
	} else if berr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" bundler is not installed"))
	} else {
		bundlerInit := bundlePath + " init"

		cmd := exec.Command("bash", "-c", bundlerInit)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", bundlerInit)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.rb"), []byte(MainRbContent()), 0644)
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

		bundleConfig := bundlePath + " config set --local path '~/.gem'"

		configCmd := exec.Command("bash", "-c", bundleConfig)

		if runtime.GOOS == "windows" {
			configCmd = exec.Command("powershell.exe", bundleConfig)
		}

		configCmd.Dir = botName
		configCmd.Stdin = os.Stdin
		configCmd.Stdout = os.Stdout
		configCmd.Stderr = os.Stderr
		err = configCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		bundleAdd := bundlePath + " add telegram-bot-ruby bwrb"

		addCmd := exec.Command("bash", "-c", bundleAdd)

		if runtime.GOOS == "windows" {
			addCmd = exec.Command("powershell.exe", bundleAdd)
		}

		addCmd.Dir = botName
		addCmd.Stdin = os.Stdin
		addCmd.Stdout = os.Stdout
		addCmd.Stderr = os.Stderr
		err = addCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
