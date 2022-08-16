package cpp

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

func TelegramCpp(botName string) {
	_, err := looker.LookPath("gcc")
	_, cerr := looker.LookPath("cmake")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" gcc is not installed"))
	} else if cerr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" cmake is not installed"))
	} else {
		botwayHeader := os.WriteFile(filepath.Join(botName, "src", "botway.hpp"), []byte((BWCPPFileContent(botName))), 0644)
		dotDockerIgnoreFile := os.WriteFile(filepath.Join(botName, ".dockerignore"), []byte(DotDockerIgnoreContent()), 0644)
		cmakeListsFile := os.WriteFile(filepath.Join(botName, "CMakeLists.txt"), []byte(CmakeListsContent(botName)), 0644)
		runPsFile := os.WriteFile(filepath.Join(botName, "run.ps1"), []byte(RunPsFileContent()), 0644)
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.cpp"), []byte(MainCppContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if botwayHeader != nil {
			log.Fatal(botwayHeader)
		}

		if runPsFile != nil {
			log.Fatal(runPsFile)
		}

		if dotDockerIgnoreFile != nil {
			log.Fatal(dotDockerIgnoreFile)
		}

		if cmakeListsFile != nil {
			log.Fatal(cmakeListsFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		install := "curl -sL https://raw.githubusercontent.com/botwayorg/telegram-cpp/main/scripts/install-deps.sh | bash"
		build, shell := "mkdir build; cd build; cmake ..; make -j", "bash"

		if runtime.GOOS == "windows" {
			install = "irm https://raw.githubusercontent.com/botwayorg/telegram-cpp/main/scripts/install-deps.ps1 | iex"
			build, shell = `.\run.ps1`, "powershell.exe"
		}

		installCmd := exec.Command(shell, "-c", install)

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		run := exec.Command(shell, "-c", build)

		run.Dir = botName
		run.Stdin = os.Stdin
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr

		err = run.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
