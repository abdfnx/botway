package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/looker"
	"github.com/botwayorg/templates"
	"github.com/spf13/cobra"
)

func ComposeCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "compose",
		Short:  "Run and manage mulit-bots with docker compose",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run: func(cmd *cobra.Command, args []string) {
			_, err := looker.LookPath("docker")

			if err != nil {
				fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
				panic(constants.FAIL_FOREGROUND.Render(" docker is not installed"))
			}

			tools.CreateEnvFile()

			dockerCompose := exec.Command("docker-compose", args...)

			dockerCompose.Stdin = os.Stdin
			dockerCompose.Stdout = os.Stdout
			dockerCompose.Stderr = os.Stderr
			err = dockerCompose.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}
		},
	}

	cmd.AddCommand(ComposeInitCMD())

	return cmd
}

func ComposeInitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "init",
		Short:  "Initialize `docker-compose.yaml`",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run: func(cmd *cobra.Command, args []string) {
			dockerComposeFile := os.WriteFile("docker-compose.yaml", []byte(templates.Content("dockerfiles/compose/docker-compose.yaml", "botway", "", "")), 0644)

			if dockerComposeFile != nil {
				log.Fatal(dockerComposeFile)
			}
		},
	}

	return cmd
}
