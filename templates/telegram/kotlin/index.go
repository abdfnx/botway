package kotlin

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

func createDirs(botName string) {
	if err := os.Mkdir(filepath.Join(botName, "gradle"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "gradle", "wrapper"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", "kotlin"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", "kotlin", "botway"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", "kotlin", "core"), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func TelegramKotlin(botName string) {
	createDirs(botName)

	gradle, err := looker.LookPath("gradle")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" gradle is not wrappered"))
	} else {
		botlinFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "kotlin", "botway", "Botway.kt"), []byte(BotlinContent()), 0644)
		mainFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "kotlin", "core", "Bot.kt"), []byte(MainKtContent()), 0644)
		buildGradleFile := os.WriteFile(filepath.Join(botName, "app", "build.gradle.kts"), []byte(BuildGradleKtsContent()), 0644)
		gradleWrapperPropsFile := os.WriteFile(filepath.Join(botName, "gradle", "wrapper", "gradle-wrapper.properties"), []byte(GradleWrapperPropsContent()), 0644)
		gradlewFile := os.WriteFile(filepath.Join(botName, "gradlew"), []byte(GradlewContent()), 0644)
		gradlewBatFile := os.WriteFile(filepath.Join(botName, "gradlew.bat"), []byte(GradlewBatContent()), 0644)
		settingsFile := os.WriteFile(filepath.Join(botName, "settings.gradle.kts"), []byte(SettingsGradleKts()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)
		gitattributesFile := os.WriteFile(filepath.Join(botName, ".gitattributes"), []byte(DotGitattributesContent()), 0644)

		if botlinFile != nil {
			log.Fatal(botlinFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if buildGradleFile != nil {
			log.Fatal(buildGradleFile)
		}

		if gradleWrapperPropsFile != nil {
			log.Fatal(gradleWrapperPropsFile)
		}

		if gradlewFile != nil {
			log.Fatal(gradlewFile)
		}

		if gradlewBatFile != nil {
			log.Fatal(gradlewBatFile)
		}

		if settingsFile != nil {
			log.Fatal(settingsFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if gitattributesFile != nil {
			log.Fatal(gitattributesFile)
		}

		gradleWrapper := gradle + " wrapper"

		wrapperCmd := exec.Command("bash", "-c", gradleWrapper)

		if runtime.GOOS == "windows" {
			wrapperCmd = exec.Command("powershell.exe", gradleWrapper)
		}

		wrapperCmd.Dir = botName
		wrapperCmd.Stdin = os.Stdin
		wrapperCmd.Stdout = os.Stdout
		wrapperCmd.Stderr = os.Stderr
		err = wrapperCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		templates.CheckProject(botName, "telegram")
	}
}
