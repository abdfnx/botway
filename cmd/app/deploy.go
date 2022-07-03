package app

import (
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func DeployCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deploy [path]",
		Short:   "Deploy and upload project from the current directory",
		Aliases: []string{"up"},
		PreRun:  func(cmd *cobra.Command, args []string) { tools.SetupTokensInDocker() },
		RunE:    Contextualize(handler.Delpoy, handler.Panic),
		PostRun: func(cmd *cobra.Command, args []string) { tools.RemoveConfig() },
	}

	cmd.AddCommand(DeployDownCMD())
	cmd.AddCommand(DeployLogsCMD())
	cmd.AddCommand(DeployLiveCMD())

	cmd.Flags().BoolP("detach", "d", false, "Detach from cloud build/deploy logs")
	cmd.Flags().StringP("environment", "e", "", "Specify an environment to up onto")
	cmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	return cmd
}

func DeployDownCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "down",
		Short:  "Remove the most recent deployment",
		PreRun: func(cmd *cobra.Command, args []string) { CheckDir() },
		RunE:   Contextualize(handler.Down, handler.Panic),
	}

	return cmd
}

func DeployLogsCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "logs",
		Short:  "View the most-recent deploy's logs",
		PreRun: func(cmd *cobra.Command, args []string) { CheckDir() },
		RunE:   Contextualize(handler.Logs, handler.Panic),
	}

	return cmd
}

func DeployLiveCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "live",
		Short:  "Open the deployed application",
		PreRun: func(cmd *cobra.Command, args []string) { CheckDir() },
		RunE:   Contextualize(handler.OpenApp, handler.Panic),
	}

	return cmd
}
