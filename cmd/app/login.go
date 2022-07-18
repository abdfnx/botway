package app

import (
	"github.com/spf13/cobra"
)

func LogoutCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Authenticate with Railway",
		Aliases: []string{"auth"},
		RunE:    Contextualize(handler.Login, handler.Panic),
	}

	return cmd
}
