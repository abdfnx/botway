package app

import (
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func SurrealCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "surrealdb",
		Short:   "Manage your work with PocketBase",
		Aliases: []string{"surreal", "sdb"},
	}

	cmd.AddCommand(SurrealInitCMD())

	return cmd
}

func SurrealInitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "init",
		Short:  "Initialize surrealdb in your docker compose config file",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run: func(cmd *cobra.Command, args []string) {
			tools.InitInDockerCompose("surrealdb")
		},
	}

	return cmd
}
