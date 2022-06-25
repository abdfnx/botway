package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/tidwall/sjson"
)

func main() {
	updateVersionOnPackageJSON(os.Args[1])
	gitTag(os.Args[1])
	gitPushOrigin(os.Args[1])
	publishOnNPM()
}

func updateVersionOnPackageJSON(version string) {
	version = version[1:]

	packageJSON, err := ioutil.ReadFile("package.json")

	if err != nil {
		log.Fatal(err)
	}

	newVersion, err := sjson.Set(string(packageJSON), "version", version)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("package.json", []byte(newVersion), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func gitTag(version string) {
	cmd := "git tag -a " + version + " -m \"New release " + version + "\""

	runCmd := exec.Command("bash", "-c", cmd)

	if runtime.GOOS == "windows" {
		runCmd = exec.Command("powershell.exe", cmd)
	}

	runCmd.Stdin = os.Stdin
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	err := runCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func gitPushOrigin(version string) {
	cmd := "git push origin " + version
	runCmd := exec.Command("bash", "-c", cmd)

	if runtime.GOOS == "windows" {
		runCmd = exec.Command("powershell.exe", cmd)
	}

	runCmd.Stdin = os.Stdin
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	err := runCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func publishOnNPM() {
	cmd := "yarn publish"

	runCmd := exec.Command("bash", "-c", cmd)

	if runtime.GOOS == "windows" {
		runCmd = exec.Command("powershell.exe", cmd)
	}

	runCmd.Stdin = os.Stdin
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	err := runCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
