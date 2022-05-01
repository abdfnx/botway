package pip

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/tools/templates/slack/python"
	"github.com/abdfnx/looker"
)

func SlackPythonPip(botName string) {
	pip := "pip3"
	pythonPath := "python3"

	if runtime.GOOS == "windows" {
		pip = "pip"
		pythonPath = "python"
	}

	_, err := looker.LookPath(pythonPath)
	_, perr := looker.LookPath(pip)

	if err != nil {
		log.Fatal("error: python is not installed")
	} else if perr != nil {
		log.Fatal("error: pip is not installed")
	} else {
		requirementsFile := os.WriteFile(filepath.Join(botName, "requirements.txt"), []byte(RequirementsContent()), 0644)

		if requirementsFile != nil {
			log.Fatal(requirementsFile)
		}

		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(python.Resources()), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		pipInstall := pip + " install -r requirements.txt"

		cmd := exec.Command("bash", "-c", pipInstall)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", pipInstall)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.py"), []byte(python.MainPyContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: python3 ./src/main.py"), 0644)
		runtimeFile := os.WriteFile(filepath.Join(botName, "runtime.txt"), []byte("python-3.9.6"), 0644)
		flake8File := os.WriteFile(filepath.Join(botName, ".flake8"), []byte(python.Flake8Content()), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if procFile != nil {
			log.Fatal(procFile)
		}

		if runtimeFile != nil {
			log.Fatal(runtimeFile)
		}

		if flake8File != nil {
			log.Fatal(flake8File)
		}
	}
}
