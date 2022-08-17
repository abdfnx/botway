package new

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/internal/options"
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

	viper.SetConfigType("yaml")

	configFile := filepath.Join(botName, ".botway.yaml")

	configFileContent, _ := ioutil.ReadFile(configFile)

	viper.ReadConfig(bytes.NewBuffer(configFileContent))

	viper.Set("bot.repo", "github.com/"+viper.GetString("author")+"/"+o.RepoName)

	newConfig := viper.WriteConfigAs(configFile)

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
