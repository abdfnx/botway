package app

import (
	"github.com/abdfnx/botway/internal/pipes/start"
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func StartCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "start",
		Short:  "Start Running your bot",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		Run: func(cmd *cobra.Command, args []string) {
			start.Start()
		},
	}

	return cmd
}
