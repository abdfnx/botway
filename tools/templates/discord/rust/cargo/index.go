package cargo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/tools/templates/discord"
	"github.com/abdfnx/botway/tools/templates/discord/rust"
	"github.com/abdfnx/looker"
)

func DiscordRustCargo(botName string) {
	_, err := looker.LookPath("rust")
	cargoPath, cerr := looker.LookPath("cargo")

	if err != nil {
		log.Fatal("error: rust is not installed")
	} else if cerr != nil {
		log.Fatal("error: cargo is not installed")
	} else {
		if runtime.GOOS == "linux" {
			fmt.Println("Installing some required linux packages")

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

		cargoBuild := cargoPath + " build"

		buildCmd := exec.Command("bash", "-c", cargoBuild)

		if runtime.GOOS == "windows" {
			buildCmd = exec.Command("powershell.exe", cargoBuild)
		}

		buildCmd.Dir = botName
		buildCmd.Stdin = os.Stdin
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		err = buildCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
}
