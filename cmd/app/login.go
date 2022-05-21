package app

import (
	"github.com/spf13/cobra"
)

func LoginCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with Railway",
		Aliases: []string{"auth"},
		RunE:  Contextualize(handler.Login, handler.Panic),
	}

	// addCmd(cmd, &cobra.Command{
	// 	Use:   "github",
	// 	Short: "Login to your GitHub account",
	// 	Aliases: []string{"gh"},
	// 	Run: func(cmd *cobra.Command, args []string) {},
	// })

	// addCmd(cmd, &cobra.Command{
	// 	Use:   "waypoint",
	// 	Short: "Login to your Hashicorp (Waypoint) account",
	// 	Aliases: []string{"wp"},
	// 	Run: func(cmd *cobra.Command, args []string) {},
	// })

	return cmd
}
