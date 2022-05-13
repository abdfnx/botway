package app

import (
	"github.com/abdfnx/botway/internal/options"
	"github.com/spf13/cobra"
)

var loginOpts = options.LoginOptions{
	Railway: false,
}

func LoginCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with Railway or GitHub.",
		Aliases: []string{"auth"},
	}


	addCmd(cmd, &cobra.Command{
		Use:   "railway",
		Short: "Login to your Railway account",
		RunE:  Contextualize(handler.Login, handler.Panic),
	})

	return cmd
}
