package new

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/viper"
)

func createRepoWindows(botName string, isPrivate bool) string {
	privateFlag := ""

	if isPrivate {
		privateFlag = "--private"
	}

	return fmt.Sprintf(`
		$username = botway gh get-username
		git init
		botway gh-repo create %s -d "My Awesome bot" %s -y
		git add .
		git commit -m "new botway bot project"
		git branch -M main
		git push -u origin main
	`, botName, privateFlag)
}

func createRepoUnix(botName string, isPrivate bool) string {
	repoStatus := "--public"

	if isPrivate {
		repoStatus = "--private"
	}

	return fmt.Sprintf(`
		username=$(botway gh get-username)
		git init
		botway gh-repo create %s -d "My Awesome bot" %s -y
		git add .
		git commit -m "new botway bot project"
		git branch -M main
		git push -u origin main
	`, botName, repoStatus)
}

func CreateRepo(o *options.NewOptions, botName string) {
	if o.RepoName == "" {
		o.RepoName = botName
	}

	viper.Set("bot.repo", "github.com/"+botwaygo.GetBotInfo("author")+"/"+o.RepoName)

	newConfig := viper.WriteConfigAs(".botway.yaml")

	if newConfig != nil {
		log.Fatal(newConfig)
	}

	createRepoCmd := exec.Command("bash", "-c", createRepoUnix(o.RepoName, o.IsPrivate))

	if runtime.GOOS == "windows" {
		createRepoCmd = exec.Command("powershell.exe", "-Command", createRepoWindows(o.RepoName, o.IsPrivate))
	}

	createRepoCmd.Dir = botName
	createRepoCmd.Stdin = os.Stdin
	createRepoCmd.Stdout = os.Stdout
	createRepoCmd.Stderr = os.Stderr
	err := createRepoCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
