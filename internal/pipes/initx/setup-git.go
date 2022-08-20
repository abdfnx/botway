package initx

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
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
