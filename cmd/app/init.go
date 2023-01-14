package app

import (
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/cobra"
)

var opts = options.InitOptions{
	NoRepo: false,
}

func InitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Initialize ~/.botway",
		Aliases: []string{"."},
		Run: func(cmd *cobra.Command, args []string) {
			initx.Init()

			if !opts.NoRepo {
				initx.SetupGitRepo()
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.NoRepo, "no-repo", "", false, "Don't create a private git repo under my account")

	return cmd
}

func GenerateCConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "c-init",
		Short:  "Initialize config for c projects (only for discord bots)",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			tools.GenerateCConfig(botwaygo.GetToken())
		},
	}

	return cmd
}
