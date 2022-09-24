package app

import (
	build_image "github.com/abdfnx/botway/internal/pipes/docker/build-image"
	"github.com/abdfnx/botway/internal/pipes/docker/run"
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func DockerCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker",
		Short: "Manage your bots docker images",
		Long:  "With `botway docker` command you can manage all your bots docker images",
	}

	cmd.AddCommand(DockerBuildCMD())
	cmd.AddCommand(DockerRunCMD())

	return cmd
}

func DockerBuildCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "build",
		Short:  "Build your bot docker image",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run:    func(cmd *cobra.Command, args []string) { build_image.DockerBuildImage() },
	}

	return cmd
}

func DockerRunCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "run",
		Short:  "Run your bot docker image",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run:    func(cmd *cobra.Command, args []string) { run.DockerRunImage() },
	}

	return cmd
}
