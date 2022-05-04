package app

import (
	"github.com/abdfnx/botway/internal/pipes/docker/build-image"
	"github.com/spf13/cobra"
)

func DockerCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker",
		Short: "Manage your bots docker images.",
		Long: "With `botway docker` command you can manage all your bots docker images.",
	}

	cmd.AddCommand(DockerBuildCMD())

	return cmd
}

func DockerBuildCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build your bot docker image.",
		Run: func(cmd *cobra.Command, args []string) {
			build_image.DockerBuildImage()
		},
	}

	return cmd
}
