package app

import (
	"github.com/abdfnx/botway/internal/config"
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/remove"
	"github.com/abdfnx/botway/internal/render"
	"github.com/spf13/cobra"
)

func RemoveCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   "Remove a botway project",
		Aliases: []string{"rm", "delete"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts := &options.CommonOptions{
					BotName: args[0],
				}

				if config.GetBotInfoFromArg(args[0], "bot.host_service") == "render.com" {
					render.DeleteRenderService(args[0])
				} else if config.GetBotInfoFromArg(args[0], "bot.host_service") == "railway.app" {
					cmd.PostRunE = Contextualize(handler.Delete, handler.Panic)
				}

				remove.Remove(opts)
			} else {
				cmd.Help()
			}
		},
	}

	return cmd
}
