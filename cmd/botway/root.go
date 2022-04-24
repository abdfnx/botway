package botway

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/MakeNowJust/heredoc"
	"github.com/abdfnx/botway/cmd/factory"
	"github.com/elewis787/boa"
)

func Execute(f *factory.Factory, version string, buildDate string) *cobra.Command {
	const desc = `🤖 A bot framework to build and handle your own bot, for Telegram, or Discord, or Slack`

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "botway <subcommand> [flags]",
		Short:  desc,
		Example: heredoc.Doc(""),
		Annotations: map[string]string{
			"help:tellus": heredoc.Doc(`
				Open an issue at https://github.com/abdfnx/botway/issues
			`),
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Aliases: []string{"ver"},
		Short: "Print the version of your botway binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("botway version " + version + " " + buildDate)
		},
	}

	rootCmd.SetOut(f.IOStreams.Out)
	rootCmd.SetErr(f.IOStreams.ErrOut)

	rootCmd.PersistentFlags().Bool("help", false, "Help for botway")
	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)

	// Add sub-commands to root command
	rootCmd.AddCommand(versionCmd)

	return rootCmd
}
