package app

import (
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func PocketBaseCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pocketbase",
		Short:   "Manage your work with PocketBase",
		Aliases: []string{"pb"},
	}

	cmd.AddCommand(PocketBaseInitCMD())

	return cmd
}

func PocketBaseInitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "init",
		Short:  "Initialize pocketbase in your docker compose config file",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run: func(cmd *cobra.Command, args []string) {
			tools.InitInDockerCompose("pocketbase")
		},
	}

	return cmd
}
