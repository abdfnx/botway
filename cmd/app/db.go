package app

import (
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func DBCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "database",
		Aliases: []string{"db"},
		Short:   "Manage your bot database (With Railway)",
	}

	dbAddCmd := &cobra.Command{
		Use:    "add",
		Short:  "Add a new database plugin to your bot project",
		Args:   cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		RunE:   Contextualize(handler.Add, handler.Panic),
	}

	dbConnectCmd := &cobra.Command{
		Use:    "connect",
		Short:  "Open an interactive shell to a database",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		RunE:   Contextualize(handler.Connect, handler.Panic),
	}

	cmd.AddCommand(dbAddCmd)
	cmd.AddCommand(dbConnectCmd)

	return cmd
}
