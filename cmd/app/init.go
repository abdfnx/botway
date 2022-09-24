package app

import (
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/spf13/cobra"
)

var opts = options.InitOptions{
	NoRepo:   false,
}

func InitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Initialize ~/.botway",
		Aliases: []string{"."},
		Run: func(cmd *cobra.Command, args []string) {
			initx.Init()
			initx.InitBWBots()

			if !opts.NoRepo {
				initx.SetupGitRepo()
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.NoRepo, "no-repo", "", false, "Don't create a private git repo under my account")

	return cmd
}
