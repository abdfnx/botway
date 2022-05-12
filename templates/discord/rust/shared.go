package rust

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/templates"
	"github.com/abdfnx/botway/templates/discord"
	"github.com/abdfnx/looker"
	"github.com/charmbracelet/lipgloss"
)

func MainRsContent() string {
	return templates.Content("discord/rust/assets/src/main.rs", "")
}

func CargoFileContent(botName string) string {
	return templates.Content("discord/rust/assets/Cargo.toml", botName)
}

func Resources() string {
	return `# Botway Discord (Rust 🦀) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [A Rust library for the Discord API](https://github.com/serenity-rs/serenity)
- [An async Rust library for the Discord voice API](https://github.com/serenity-rs/songbird)
- [Discord Server](https://discord.gg/serenity-rs)

## Examples

- [serenity examples](https://github.com/serenity-rs/serenity/tree/current/examples)
- [songbird examples](https://github.com/serenity-rs/songbird/tree/current/examples)

big thanks to [**@serenity-rs**](https://github.com/serenity-rs) org`
}

func DiscordRust(botName, pm string) {
	_, err := looker.LookPath("cargo")
	pmPath, perr := looker.LookPath(pm)
	messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" cargo is not installed"))
	} else if perr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" " + pm + "  is not installed"))
	} else {
		if runtime.GOOS == "linux" {
			fmt.Println(messageStyle.Render("> Installing some required linux packages"))

			discord.InstallCommandRust()
		}

		DockerfileContent := templates.Content("assets/" + pm + ".dockerfile", botName)

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.rs"), []byte(MainRsContent()), 0644)
		cargoFile := os.WriteFile(filepath.Join(botName, "Cargo.toml"), []byte(CargoFileContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

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

		pmBuild := pmPath + " build"
		buildCmd := exec.Command("bash", "-c", pmBuild)

		if runtime.GOOS == "windows" {
			buildCmd = exec.Command("powershell.exe", pmBuild)
		}

		buildCmd.Dir = botName
		buildCmd.Stdin = os.Stdin
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		err = buildCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		if pm == "fleet" {
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
		}

		templates.CheckProject(botName, "discord")
	}
}
