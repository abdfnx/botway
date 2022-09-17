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
	"github.com/charmbracelet/lipgloss"
)

func DiscordRuby(botName string) {
	_, err := looker.LookPath("ruby")
	bundlePath, berr := looker.LookPath("bundle")
	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" ruby is not installed"))
	} else if berr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" bundler is not installed"))
	} else {
		if runtime.GOOS == "windows" {
			fmt.Println(messageStyle.Render(`On Windows, follow these steps:

Download the latest libsodium-X.Y.Z-msvc.zip from https://download.libsodium.org/libsodium/releases.
From the downloaded zip file, extract the 'x64/Release/v120/dynamic/libsodium.dll' file to somewhere.
Copy that to any folder within the Ruby '$LOAD_PATH' or 'C:\Windows\System32' and rename it to 'sodium.dll'.
You can add a folder to your '$LOAD_PAT'H either at runtime or via the -I command line flag (ruby -I ./my_dlls bot.rb).`))
		}

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

		bundleAdd := bundlePath + " add discordrb --git https://github.com/shardlab/discordrb --branch main"

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

		bundleAddBotwayrb := bundlePath + " add bwrb"

		addBotwayCmd := exec.Command("bash", "-c", bundleAddBotwayrb)

		if runtime.GOOS == "windows" {
			addBotwayCmd = exec.Command("powershell.exe", bundleAddBotwayrb)
		}

		addBotwayCmd.Dir = botName
		addBotwayCmd.Stdin = os.Stdin
		addBotwayCmd.Stdout = os.Stdout
		addBotwayCmd.Stderr = os.Stderr
		err = addBotwayCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "discord")
	}
}
