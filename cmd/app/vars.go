package app

import (
	"github.com/abdfnx/botway/internal/render"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/cobra"
)

func VarsCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "variables",
		Aliases: []string{"vars"},
		Short:   "Show variables for active environment",
		PreRun:  func(cmd *cobra.Command, args []string) { tools.CheckDir() },
	}

	variablesGetCmd := &cobra.Command{
		Use:     "get a variable",
		Short:   "Get the value of a variable",
		Args:    cobra.MinimumNArgs(1),
		PreRun:  func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Example: "  botway variables get PORT",
	}

	variablesSetCmd := &cobra.Command{
		Use:     "set key=value",
		Short:   "Create or update the value of a variable",
		Args:    cobra.MinimumNArgs(1),
		PreRun:  func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Example: "  botway variables set KEY=VALUE",
	}

	variablesRemoveCmd := &cobra.Command{
		Use:     "remove a variable",
		Aliases: []string{"rm", "delete"},
		Short:   "Delete a variable",
		PreRun:  func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Example: "  botway variables remove MY_KEY",
	}

	if botwaygo.GetBotInfo("bot.host_service") == "railway.app" {
		cmd.RunE = Contextualize(handler.Variables, handler.Panic)

		desc := "Fetch variables accessible to a specific service"

		cmd.Flags().StringP("service", "s", "", desc)

		variablesGetCmd.RunE = Contextualize(handler.VariablesGet, handler.Panic)
		variablesSetCmd.RunE = Contextualize(handler.VariablesSet, handler.Panic)
		variablesRemoveCmd.RunE = Contextualize(handler.VariablesDelete, handler.Panic)

		variablesGetCmd.Flags().StringP("service", "s", "", desc)
		variablesSetCmd.Flags().StringP("service", "s", "", desc)
		variablesSetCmd.Flags().BoolP("no-redeploy-hint", "", false, "Don't show re-deploy hints after setting a new variable")
		variablesSetCmd.Flags().BoolP("hidden", "", false, "Hide variable value")
		variablesRemoveCmd.Flags().StringP("service", "s", "", desc)
	} else if botwaygo.GetBotInfo("bot.host_service") == "render.com" {
		cmd.Run = func(cmd *cobra.Command, args []string) {
			render.Vars(false, args)
		}

		variablesGetCmd.Run = func(cmd *cobra.Command, args []string) {
			render.Vars(true, args)
		}

		variablesSetCmd.Run = func(cmd *cobra.Command, args []string) {
			render.SetEnvVars(args)
		}
	}

	cmd.AddCommand(variablesGetCmd)
	cmd.AddCommand(variablesSetCmd)
	cmd.AddCommand(variablesRemoveCmd)

	return cmd
}
