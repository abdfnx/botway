package app

import "github.com/spf13/cobra"

func VarsCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "variables",
		Aliases: []string{"vars"},
		Short:   "Show variables for active environment",
		RunE:    Contextualize(handler.Variables, handler.Panic),
	}

	cmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	variablesAddCmd := &cobra.Command{
		Use:     "get a variable",
		Short:   "Get the value of a variable",
		RunE:    Contextualize(handler.VariablesGet, handler.Panic),
		Args:    cobra.MinimumNArgs(1),
		Example: "  botway variables get PORT",
	}

	cmd.AddCommand(variablesAddCmd)
	variablesAddCmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	variablesSetCmd := &cobra.Command{
		Use:     "set key=value",
		Short:   "Create or update the value of a variable",
		RunE:    Contextualize(handler.VariablesSet, handler.Panic),
		Args:    cobra.MinimumNArgs(1),
		Example: "  botway variables set KEY=VALUE",
	}

	cmd.AddCommand(variablesSetCmd)
	variablesSetCmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	variablesRemoveCmd := &cobra.Command{
		Use:     "remove a variable",
		Aliases: []string{"rm", "delete"},
		Short:   "Delete a variable",
		RunE:    Contextualize(handler.VariablesDelete, handler.Panic),
		Example: "  botway variables remove MY_KEY",
	}

	cmd.AddCommand(variablesRemoveCmd)
	variablesRemoveCmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	return cmd
}
