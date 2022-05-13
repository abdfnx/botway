package app

import "github.com/spf13/cobra"

func DeployCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [path]",
		Short: "Deploy and upload project from the current directory",
		RunE:  Contextualize(handler.Delpoy, handler.Panic),
	}

	cmd.Flags().BoolP("detach", "d", false, "Detach from cloud build/deploy logs")
	cmd.Flags().StringP("environment", "e", "", "Specify an environment to up onto")
	cmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	return cmd
}
