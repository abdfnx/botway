package fleet

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools/templates"
	"github.com/abdfnx/botway/tools/templates/discord"
	"github.com/abdfnx/botway/tools/templates/discord/rust"
	"github.com/abdfnx/looker"
	"github.com/charmbracelet/lipgloss"
)

func DiscordRustFleet(botName string) {
	_, err := looker.LookPath("rustc")
	fleetPath, ferr := looker.LookPath("fleet")
	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" rust is not installed"))
	} else if ferr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" fleet is not installed"))
	} else {
		if runtime.GOOS == "linux" {
			fmt.Println(messageStyle.Render("> Installing some required linux packages"))

			discord.InstallCommandRust()
		}

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.rs"), []byte(rust.MainRsContent()), 0644)
		cargoFile := os.WriteFile(filepath.Join(botName, "Cargo.toml"), []byte(rust.CargoFileContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: ./" + botName), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(rust.Resources()), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if cargoFile != nil {
			log.Fatal(cargoFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if procFile != nil {
			log.Fatal(procFile)
		}

		rustUpPath, err := looker.LookPath("rustup")

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		rustUpCmd := rustUpPath + " default nightly"

		rustUp := exec.Command("bash", "-c", rustUpCmd)

		if runtime.GOOS == "windows" {
			rustUp = exec.Command("powershell.exe", rustUpCmd)
		}

		rustUp.Dir = botName
		rustUp.Stdin = os.Stdin
		rustUp.Stdout = os.Stdout
		rustUp.Stderr = os.Stderr
		err = rustUp.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		fleetBuild := fleetPath + " build"

		buildCmd := exec.Command("bash", "-c", fleetBuild)

		if runtime.GOOS == "windows" {
			buildCmd = exec.Command("powershell.exe", fleetBuild)
		}

		buildCmd.Dir = botName
		buildCmd.Stdin = os.Stdin
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		err = buildCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "discord")
	}
}
