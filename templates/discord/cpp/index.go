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
	"github.com/charmbracelet/lipgloss"
)

func DiscordCpp(botName, hostService string) {
	_, err := looker.LookPath("gcc")
	_, cerr := looker.LookPath("cmake")

	if err := os.Mkdir(filepath.Join(botName, "cmake"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "include"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "include", "botway"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "include", botName), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" gcc is not installed"))
	} else if cerr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" cmake is not installed"))
	} else {
		findDPPCmakeFile := os.WriteFile(filepath.Join(botName, "cmake", "FindDPP.cmake"), []byte(FindDppCmakeContent()), 0644)
		botwayHeader := os.WriteFile(filepath.Join(botName, "include", "botway", "botway.hpp"), []byte((BWCPPFileContent(botName))), 0644)
		mainIncludeFile := os.WriteFile(filepath.Join(botName, "include", botName, botName+".h"), []byte((MainIncludeFileContent())), 0644)
		dotDockerIgnoreFile := os.WriteFile(filepath.Join(botName, ".dockerignore"), []byte(DotDockerIgnoreContent()), 0644)
		cmakeListsFile := os.WriteFile(filepath.Join(botName, "CMakeLists.txt"), []byte(CmakeListsContent(botName)), 0644)
		runPsFile := os.WriteFile(filepath.Join(botName, "run.ps1"), []byte(RunPsFileContent()), 0644)
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.cpp"), []byte(MainCppContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

		if findDPPCmakeFile != nil {
			log.Fatal(findDPPCmakeFile)
		}

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

		if mainIncludeFile != nil {
			log.Fatal(mainIncludeFile)
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

		build, shell := "mkdir build; cd build; cmake ..; make -j", "bash"

		if runtime.GOOS == "windows" {
			build, shell = `.\run.ps1`, "powershell.exe"
			messageStyle := lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)

			fmt.Println(messageStyle.Render(`On Windows, follow instructions at https://dpp.dev/buildwindows.html`))
		} else {
			pos := "linux"

			if runtime.GOOS == "darwin" {
				pos = "osx"
			}

			installCmd := exec.Command("bash", "-c", "curl -sL https://bit.ly/dpp-"+pos+" | bash")

			installCmd.Dir = botName
			installCmd.Stdin = os.Stdin
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr

			err = installCmd.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}
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

		templates.CheckProject(botName, "discord")
	}
}
