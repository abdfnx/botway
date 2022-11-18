package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/tidwall/sjson"
)

func main() {
	updateVersionOnPackageJSON(os.Args[1])
	gitCommit()
	gitTag(os.Args[1])
	gitPushOrigin(os.Args[1])
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

func gitCommit() {
	cmd := "git add . && git commit -m 'bump' && git push"

	commitCmd := exec.Command("bash", "-c", cmd)

	commitCmd.Stdin = os.Stdin
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	err := commitCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func gitTag(version string) {
	cmd := "git tag -a " + version + " -m \"New release " + version + "\""

	runCmd := exec.Command("bash", "-c", cmd)

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

	runCmd.Stdin = os.Stdin
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	err := runCmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
