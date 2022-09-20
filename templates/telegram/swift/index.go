package swift

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

func TelegramSwift(botName, hostService string) {
	swiftPath, err := looker.LookPath("swift")

	if err := os.Mkdir(filepath.Join(botName, "Sources"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "Sources", botName), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" swift is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "Sources", botName, "main.swift"), []byte(MainSwiftContent()), 0644)
		botwaySwiftFile := os.WriteFile(filepath.Join(botName, "Sources", botName, "botway.swift"), []byte(BotwaySwiftContent(botName)), 0644)
		packageSwiftFile := os.WriteFile(filepath.Join(botName, "Package.swift"), []byte(PackageSwiftFileContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if botwaySwiftFile != nil {
			log.Fatal(botwaySwiftFile)
		}

		if packageSwiftFile != nil {
			log.Fatal(packageSwiftFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		swiftBuild := swiftPath + " build"

		installCmd := exec.Command("bash", "-c", swiftBuild)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", swiftBuild)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
