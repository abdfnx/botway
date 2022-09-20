package poetry

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/templates"
	"github.com/abdfnx/botway/templates/telegram/python"
	"github.com/abdfnx/looker"
)

func PyProjectContent(botName string) string {
	return templates.Content("pyproject.toml", "discord-python", botName)
}

func TelegramPythonPoetry(botName, hostService string) {
	pythonPath := "python3"

	if runtime.GOOS == "windows" {
		pythonPath = "python"
	}

	_, err := looker.LookPath(pythonPath)
	poetry, perr := looker.LookPath("poetry")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" python is not installed"))
	} else if perr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" poetry is not installed"))
	} else {
		dockerFileContent := templates.Content(fmt.Sprintf("dockerfiles/%s/poetry.dockerfile", hostService), "botway", botName)

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.py"), []byte(python.MainPyContent()), 0644)
		pyprojectFile := os.WriteFile(filepath.Join(botName, "pyproject.toml"), []byte(PyProjectContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(dockerFileContent), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(python.Resources()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if pyprojectFile != nil {
			log.Fatal(pyprojectFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		poetryAdd := poetry + " add python-telegram-bot botway.py pyyaml cryptography PySocks ujson"

		cmd := exec.Command("bash", "-c", poetryAdd)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", poetryAdd)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "discord")
	}
}
