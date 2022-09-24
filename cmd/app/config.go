package app

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/botwayorg/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ConfigCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Update, Sync, Push changes, or Clone your botway configuration settings",
	}

	cmd.AddCommand(ConfigUpdateCMD())
	cmd.AddCommand(ConfigSyncCMD())
	cmd.AddCommand(ConfigCloneCMD())

	return cmd
}

func ConfigUpdateCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"set"},
		Short:   "Update botway configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigType("json")

			viper.ReadConfig(bytes.NewBuffer(constants.BotwayConfig))

			var newValue interface{} = args[1]

			if args[1] == "true" {
				newValue = true
			} else if args[1] == "false" {
				newValue = false
			}

			// set new key value but keep existing values
			viper.Set("botway.settings."+args[0], newValue)

			// write config to file
			err := viper.WriteConfigAs(constants.BotwayConfigFile)

			if err != nil {
				return err
			}

			fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
			fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Updated botway configuration"))

			return nil
		},
	}

	return cmd
}

func ConfigSyncCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync your config file at your private repo (" + git.GitConfig() + "/.botway)",
		Run: func(cmd *cobra.Command, args []string) {
			initx.UpdateConfig()
		},
	}

	return cmd
}

func ConfigCloneCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone your .botway repo",
		Run: func(cmd *cobra.Command, args []string) {
			clone := `botway gh-repo clone .botway ` + constants.BotwayDirPath

			cloneCommand := exec.Command("bash", "-c", clone)

			if runtime.GOOS == "windows" {
				cloneCommand = exec.Command("powershell.exe", "-Command", clone)
			}

			cloneCommand.Stdin = os.Stdin
			cloneCommand.Stdout = os.Stdout
			cloneCommand.Stderr = os.Stderr
			err := cloneCommand.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}
		},
	}

	return cmd
}
