package app

import (
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func ExecCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "exec",
		Short:              "Execute a local command using variables from the active environment (Only for bots that hosted on Railway)",
		PreRun:             func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		RunE:               Contextualize(handler.Exec, handler.Panic),
		DisableFlagParsing: true,
	}

	cmd.Flags().Bool("ephemeral", false, "Execute the local command in an ephemeral environment")
	cmd.Flags().String("service", "", "Fetch variables accessible to a specific service")

	return cmd
}
