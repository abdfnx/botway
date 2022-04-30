package cli

import (
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/new"
	"github.com/spf13/cobra"
)

func NewCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new botway project.",
		Long: "With `botway new` command you can create your botway project.",
		Aliases: []string{"create"},
		Run: func(cmd *cobra.Command, args []string) {
			opts := &options.NewOptions{
				BotName: args[0],
			}

			new.New(opts)
		},
	}

	return cmd
}
