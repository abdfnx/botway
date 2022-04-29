package yarn

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/tools/templates/discord/nodejs"
	"github.com/abdfnx/looker"
)

func DiscordNodejsYarn(botName string) {
	yarn, err := looker.LookPath("yarn")

	if err != nil {
		log.Fatalf("error: %s is not installed", yarn)
	} else {
		yarnInit := yarn + " init -y"

		cmd := exec.Command("bash", "-c", yarnInit)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", yarnInit)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		indexFile := os.WriteFile(filepath.Join(botName, "src", "index.js"), []byte(nodejs.IndexJSContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: node ./src/index.js"), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(nodejs.Resources()), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if indexFile != nil {
			log.Fatal(indexFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if procFile != nil {
			log.Fatal(procFile)
		}

		yarnInstall := yarn + " add " + nodejs.Packages

		installCmd := exec.Command("bash", "-c", yarnInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", yarnInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		if runtime.GOOS == "windows" {
			installWindowsBuildTools := exec.Command("powershell.exe", yarn + " global add windows-build-tools")

			installWindowsBuildTools.Dir = botName
			installWindowsBuildTools.Stdin = os.Stdin
			installWindowsBuildTools.Stdout = os.Stdout
			installWindowsBuildTools.Stderr = os.Stderr
			err = installWindowsBuildTools.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}
		}
	}
}
