package app

import (
	"github.com/spf13/cobra"
)

func LoginCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout of your Railway account",
		RunE:  Contextualize(handler.Logout, handler.Panic),
	}

	return cmd
}
