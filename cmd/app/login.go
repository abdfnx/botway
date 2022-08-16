package app

import (
	"github.com/abdfnx/botway/internal/render"
	"github.com/spf13/cobra"
)

func LoginCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Authenticate with Railway Or Render",
		Aliases: []string{"auth"},
	}

	cmd.AddCommand(RailwayLogin())
	cmd.AddCommand(RenderLogin())

	return cmd
}

func RailwayLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "railway",
		Short:   "Authenticate with Railway",
		Aliases: []string{"rw"},
		RunE:    Contextualize(handler.Login, handler.Panic),
	}

	return cmd
}

func RenderLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "render",
		Short:   "Authenticate with Render",
		Aliases: []string{"rn"},
		Run: func(cmd *cobra.Command, args []string) {
			render.BotwayRenderAuth()
		},
	}

	return cmd
}
