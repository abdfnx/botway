package botway

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/abdfnx/botway/cmd/app"
	"github.com/abdfnx/botway/cmd/factory"
	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	"github.com/elewis787/boa"
	"github.com/spf13/cobra"
)

var opts = options.RootOptions{
	Version: false,
}

func Execute(f *factory.Factory, version string, buildDate string) *cobra.Command {
	const desc = `🤖 Generate, build, handle and deploy your own bot with your favorite language, for Discord, or Telegram, or Slack`

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "botway <subcommand> [flags]",
		Version: version,
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
	rootCmd.Flags().BoolVarP(&opts.Version, "version", "v", false, "Print the version of your botway binary.")

	boa.TitleStyle.BorderForeground(constants.PRIMARY_COLOR)
	boa.SelectedItemStyle.Background(constants.PRIMARY_COLOR)

	// Add sub-commands to root command
	rootCmd.AddCommand(
		app.InitCMD(),
		app.NewCMD(),
		app.TokenCMD(),
		app.RemoveCMD(),
		versionCmd,
	)

	return rootCmd
}
