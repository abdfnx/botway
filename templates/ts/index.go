package ts

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

func MainTSContent(platform string) string {
	return templates.Content("main.ts", platform+"-nodejs-ts", "", "")
}

func Resources(platform string) string {
	return templates.Content(platform+"/nodejs.md", "resources", "", "")
}

func NodejsTS(botName, pm, platform, hostService string) {
	_, nerr := looker.LookPath("npm")
	pmPath, err := looker.LookPath(pm)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" " + pm + " is not installed"))
	} else {
		if nerr != nil {
			fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
			fmt.Println(constants.FAIL_FOREGROUND.Render(" npm is not installed"))
		} else {
			dockerfileContent := templates.Content(fmt.Sprintf("dockerfiles/%s/%s.dockerfile", hostService, pm), "botway", botName, platform)

			mainFile := os.WriteFile(filepath.Join(botName, "src", "main.ts"), []byte(MainTSContent(platform)), 0644)
			dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(dockerfileContent), 0644)
			resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform)), 0644)
			tsConfigFile := os.WriteFile(filepath.Join(botName, "tsconfig.json"), []byte(templates.Content("tsconfig.json", platform+"-nodejs-ts", "", "")), 0644)
			packageFile := os.WriteFile(filepath.Join(botName, "package.json"), []byte(templates.Content("package.json", platform+"-nodejs-ts", "", "")), 0644)

			if resourcesFile != nil {
				log.Fatal(resourcesFile)
			}

			if mainFile != nil {
				log.Fatal(mainFile)
			}

			if dockerFile != nil {
				log.Fatal(dockerFile)
			}

			if tsConfigFile != nil {
				log.Fatal(tsConfigFile)
			}

			if packageFile != nil {
				log.Fatal(packageFile)
			}

			pmInstall := pmPath + " i"

			if pm == "yarn" {
				pmInstall = pmPath
			}

			installCmd := exec.Command("bash", "-c", pmInstall)

			if runtime.GOOS == "windows" {
				installCmd = exec.Command("powershell.exe", pmInstall)
			}

			installCmd.Dir = botName
			installCmd.Stdin = os.Stdin
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			err = installCmd.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}

			templates.CheckProject(botName, platform)
		}
	}
}
