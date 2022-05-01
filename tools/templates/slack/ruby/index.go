package ruby

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

func SlackRuby(botName string) {
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
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: bundle exec ruby ./src/main.rb"), 0644)
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

		bundleInstall := bundlePath + " add slack-ruby-bot async-websocket"

		installCmd := exec.Command("bash", "-c", bundleInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", bundleInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "slack")
	}
}
