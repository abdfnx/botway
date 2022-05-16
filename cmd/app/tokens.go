package app

import (
	"fmt"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/token/discord"
	"github.com/abdfnx/botway/internal/pipes/token/discord/guilds"
	"github.com/abdfnx/botway/internal/pipes/token/slack"
	"github.com/abdfnx/botway/internal/pipes/token/telegram"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

func TokenCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage your bots tokens.",
	}

	cmd.AddCommand(TokenSetCMD())
	cmd.AddCommand(TokenGetCMD())
	cmd.AddCommand(TokenAddGuildsCMD())

	return cmd
}

func TokenSetCMD() *cobra.Command {
	opts := &options.TokenAddOptions{
		BotName:  "",
		Discord:  false,
		Slack:    false,
		Telegram: false,
	}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Create or update the value of a bot token.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				if opts.Discord {
					discord_token.BotwayDiscordTokenSetup(args[0])
				} else if opts.Slack {
					slack_token.BotwaySlackTokenSetup(args[0])
				} else if opts.Telegram {
					telegram_token.BotwayTelegramTokenSetup(args[0])
				} else {
					fmt.Println("Bot Type is not found")
				}
			} else {
				fmt.Println("Bot Name is required")
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.Discord, "discord", "d", false, "For discord bot tokens")
	cmd.Flags().BoolVarP(&opts.Slack, "slack", "s", false, "For slack bot tokens")
	cmd.Flags().BoolVarP(&opts.Telegram, "telegram", "t", false, "For telegram bot tokens")

	return cmd
}

func TokenGetCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the value of a bot token.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Println(messageStyle.Render(strings.ToUpper(args[0]) + " Bot Token ==> ") + gjson.Get(string(constants.BotwayConfig), "botway.bots." + args[0] + ".bot_token").String())
			} else {
				fmt.Println("Bot Name is required")
			}
		},
	}

	return cmd
}

func TokenAddGuildsCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-guilds",
		Short: "Add your discord server guild ids to botway config (ONLY FOR DISCORD BOTS).",
		Run: func(cmd *cobra.Command, args []string) {
			guilds.BotwayDiscordGuildIdsSetup()
		},
	}

	return cmd
}
