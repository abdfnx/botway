package app

import "github.com/spf13/cobra"

func DBCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "database",
		Aliases: []string{"db"},
		Short:   "Manage your bot database (With Railway)",
	}

	dbAddCmd := &cobra.Command{
		Use:     "add",
		Short:   "Add a new database plugin to your bot project",
		RunE:    Contextualize(handler.Add, handler.Panic),
	}

	dbConnectCmd := &cobra.Command{
		Use:     "connect",
		Short:   "Open an interactive shell to a database",
		RunE:    Contextualize(handler.Connect, handler.Panic),
	}
	
	cmd.AddCommand(dbAddCmd)
	cmd.AddCommand(dbConnectCmd)

	return cmd
}
