package app

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/new"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/looker"
	"github.com/spf13/cobra"
)

func NewCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new botway project.",
		Long: "With `botway new` command you can create your botway project.",
		Aliases: []string{"create"},
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts := &options.CommonOptions{
					BotName: args[0],
				}

				new.New(opts)
			} else {
				cmd.Help()
			}
		},
		PostRunE: Contextualize(handler.Init, handler.Panic),
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println(messageStyle.Render("> Installing some required packages"))

			installCmd := exec.Command("bash", "-c", tools.Packages())

			if runtime.GOOS == "linux" {
				installCmd.Stdin = os.Stdin
				installCmd.Stdout = os.Stdout
				installCmd.Stderr = os.Stderr
				err := installCmd.Run()

				if err != nil {
					panic(err)
				}
			} else if runtime.GOOS == "darwin" {
				_, err := looker.LookPath("brew")

				if err != nil {
					panic("error: brew is not installed")
				} else {
					installCmd = exec.Command("bash", "-c", "brew install opus libsodium")

					err = installCmd.Run()

					if err != nil {
						panic(err)
					}
				}
			}
		},
	}

	return cmd
}
