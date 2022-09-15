package app

import (
	"github.com/abdfnx/botway/internal/pipes/compose"
	"github.com/spf13/cobra"
)

func ComposeCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compose",
		Short: "Run and manage mulit-bots with botway compose",
	}

	cmd.AddCommand(ComposeUpCMD())
	cmd.AddCommand(ComposeBuildCMD())
	cmd.AddCommand(ComposeListCMD())

	return cmd
}

func ComposeUpCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Create and start bots compose",
		Run: func(cmd *cobra.Command, args []string) {
			compose.Compose(false, false)
		},
	}

	return cmd
}

func ComposeBuildCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build or rebuild bots",
		Run: func(cmd *cobra.Command, args []string) {
			compose.Compose(true, false)
		},
	}

	return cmd
}

func ComposeListCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List bots services",
		Run: func(cmd *cobra.Command, args []string) {
			compose.Compose(false, true)
		},
	}

	return cmd
}
