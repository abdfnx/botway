package app

import (
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/cobra"
)

var opts = options.InitOptions{
	Docker: false,
	NoRepo: false,
}

func InitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Initialize ~/.botway",
		Aliases: []string{"."},
	}

	if opts.Docker {
		if botwaygo.GetBotInfo("bot.host") == "railway.app" {
			cmd.RunE = Contextualize(handler.DockerInit, handler.Panic)
		} else if botwaygo.GetBotInfo("bot.host") == "render.com" {
			cmd.Run = func(cmd *cobra.Command, args []string) {
				initx.DockerInit()
			}
		}
	} else {
		cmd.Run = func(cmd *cobra.Command, args []string) {
			initx.Init()

			if !opts.NoRepo {
				initx.SetupGitRepo()
			}
		}
	}

	cmd.Flags().BoolVarP(&opts.Docker, "docker", "", false, "Initialize botway config in docker")
	cmd.Flags().BoolVarP(&opts.NoRepo, "no-repo", "", false, "Don't create a private git repo under my account")

	return cmd
}
