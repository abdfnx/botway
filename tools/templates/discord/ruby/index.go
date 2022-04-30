package ruby

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/tools/templates/discord"
	"github.com/abdfnx/looker"
)

func DiscordRuby(botName string) {
	bundlePath, err := looker.LookPath("bundle")

	if err != nil {
		log.Fatalf("error: %s is not installed", bundlePath)
	} else {
		if runtime.GOOS == "linux" {
			fmt.Println("Installing some required linux packages")

			discord.InstallCommandRust()
		} else if runtime.GOOS == "darwin" {
			fmt.Println("Installing some required macos packages via homebrew")

			discord.InstallCommandRuby()
		} else if runtime.GOOS == "windows" {
			fmt.Println(`On Windows, follow these steps:

Download the latest libsodium-X.Y.Z-msvc.zip from https://download.libsodium.org/libsodium/releases.
From the downloaded zip file, extract the 'x64/Release/v120/dynamic/libsodium.dll' file to somewhere.
Copy that to any folder within the Ruby '$LOAD_PATH' or 'C:\Windows\System32' and rename it to 'sodium.dll'.
You can add a folder to your '$LOAD_PAT'H either at runtime or via the -I command line flag (ruby -I ./my_dlls bot.rb).`)
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

		bundleInstall := bundlePath + " add discordrb --github shardlab/discordrb --branch main"

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

		bundleInstall = bundlePath + " add botwayrb"

		installCmd = exec.Command("bash", "-c", bundleInstall)

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
	}
}
