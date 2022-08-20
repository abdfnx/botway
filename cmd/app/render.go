package app

import (
	"github.com/abdfnx/botway/internal/render"
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func RenderCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "render",
		Short:   "Manage your work with Render",
		Aliases: []string{"rn"},
	}

	cmd.AddCommand(RenderLogoutCMD())
	cmd.AddCommand(RenderConnectMD())

	return cmd
}

func RenderLogoutCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout of your Render account",
		Run: func(cmd *cobra.Command, args []string) {
			render.Logout()
		},
	}

	return cmd
}

func RenderConnectMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "connect",
		Short:   "Connect Your Render Service",
		Aliases: []string{"co"},
		Args:    cobra.MaximumNArgs(1),
		PreRun:  func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run: func(cmd *cobra.Command, args []string) {
			render.ConnectService()
		},
	}

	return cmd
}
