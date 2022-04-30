package pipenv

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/looker"
	"github.com/abdfnx/botway/tools/templates/telegram/python/pip"
)


func TelegramPythonPipenv(botName string) {
	python := "python3"

	if runtime.GOOS == "windows" {
		python = "python"
	}

	_, err := looker.LookPath(python)
	pipenv, perr := looker.LookPath("pipenv")
	
	if err != nil {
		log.Fatal("error: python is not installed")
	} else if perr != nil {
		log.Fatal("error: pipenv is not installed")
	} else {
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(pip.Resources()), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		pipenvInstall := pipenv + " install python-telegram-bot botway.py cryptography PySocks ujson"

		cmd := exec.Command("bash", "-c", pipenvInstall)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", pipenvInstall)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.py"), []byte(pip.MainPyContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: pipenv run python3 ./src/main.py"), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if procFile != nil {
			log.Fatal(procFile)
		}
	}
}
