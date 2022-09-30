package initx

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/config"
)

func windowsCommands() string {
	return `
		$username = botway gh get-username
		git init
		botway gh-repo create .botway -d "My botway config - $username" --private -y
		git add .
		git commit -m "new .botway repo"
		git branch -M main
		git push -u origin main
	`
}

func unixCommands() string {
	return `
		username=$(botway gh get-username)
		git init
		botway gh-repo create .botway -d "My botway config - $username" --private -y
		git add .
		git commit -m "new .botway repo"
		git branch -M main
		git push -u origin main
	`
}

func SetupGitRepo() {
	createRepoCmd := exec.Command("bash", "-c", unixCommands())

	if runtime.GOOS == "windows" {
		createRepoCmd = exec.Command("powershell.exe", "-Command", windowsCommands())
	}

	createRepoCmd.Dir = constants.BotwayDirPath
	createRepoCmd.Stdin = os.Stdin
	createRepoCmd.Stdout = os.Stdout
	createRepoCmd.Stderr = os.Stderr
	err := createRepoCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func UpdateConfig() {
	if config.Get("botway.settings.auto_sync") == "true" {
		cmd := `git add .
			git commit -m "New changes"
			git push`

		updateCommand := exec.Command("bash", "-c", cmd)

		if runtime.GOOS == "windows" {
			updateCommand = exec.Command("powershell.exe", "-Command", cmd)
		}

		updateCommand.Dir = constants.BotwayDirPath
		updateCommand.Stdin = os.Stdin
		updateCommand.Stdout = os.Stdout
		updateCommand.Stderr = os.Stderr
		err := updateCommand.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Configuration synced"))
	}
}
