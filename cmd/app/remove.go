package app

import (
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/remove"
	"github.com/spf13/cobra"
)

func RemoveCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a botway project.",
		Aliases: []string{"delete"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts := &options.CommonOptions{
					BotName: args[0],
				}

				remove.Remove(opts)
			} else {
				cmd.Help()
			}
		},
	}

	return cmd
}
