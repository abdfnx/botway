package app

import "github.com/spf13/cobra"

func RunCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "run",
		Short:  "Run a local command using variables from the active environment",
		PreRun: func(cmd *cobra.Command, args []string) { CheckDir() },
		RunE:   Contextualize(handler.Run, handler.Panic),
		DisableFlagParsing: true,
	}

	cmd.Flags().Bool("ephemeral", false, "Run the local command in an ephemeral environment")
	cmd.Flags().String("service", "", "Fetch variables accessible to a specific service")

	return cmd
}
