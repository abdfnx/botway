package app

import (
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/spf13/cobra"
)

func DockerInitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker-init",
		Short: "Initialize ~/.botway for docker containers",
		Run: func(cmd *cobra.Command, args []string) {
			initx.DockerInit()
		},
	}

	return cmd
}
