package npm

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/tools/templates/discord/js"
	"github.com/abdfnx/looker"
)

func DiscordNodejsNpm(botName string) {
	npm, err := looker.LookPath("npm")

	if err != nil {
		log.Fatalf("error: %s is not installed", npm)
	} else {
		npmInit := npm + " init -y"

		cmd := exec.Command("bash", "-c", npmInit)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", npmInit)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		indexFile := os.WriteFile(filepath.Join(botName, "src", "index.js"), []byte(js.IndexJSContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: node ./src/index.js"), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(js.Resources()), 0644)

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

		npmInstall := npm + " install " + js.Packages

		installCmd := exec.Command("bash", "-c", npmInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", npmInstall)
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
			installWindowsBuildTools := exec.Command("powershell.exe", npm + " install --global --production --add-python-to-path windows-build-tools")

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
