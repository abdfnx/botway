package app

import (
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/abdfnx/botway/internal/options"
	"github.com/spf13/cobra"
)

var opts = options.InitOptions{
	Docker: false,
}

func InitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize ~/.botway",
		Aliases: []string{"."},
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Docker {
				initx.DockerInit()
			} else {
				initx.BotwayInit()
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.Docker, "docker", "", false, "Initialize botway config in docker")

	return cmd
}
