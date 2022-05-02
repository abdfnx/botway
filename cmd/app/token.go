package app

import (
	"fmt"

	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/token/discord"
	"github.com/abdfnx/botway/internal/pipes/token/discord/guilds"
	"github.com/abdfnx/botway/internal/pipes/token/slack"
	"github.com/abdfnx/botway/internal/pipes/token/telegram"
	"github.com/spf13/cobra"
)

func TokenCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage your bots tokens.",
	}

	cmd.AddCommand(TokenAddCMD())
	cmd.AddCommand(TokenAddGuildsCMD())

	return cmd
}

func TokenAddCMD() *cobra.Command {
	opts := &options.TokenAddOptions{
		BotName:  "",
		Discord:  false,
		Slack:    false,
		Telegram: false,
	}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add new bot tokens.",
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

	cmd.Flags().BoolVarP(&opts.Discord, "discord", "d", false, "Add discord bot tokens")
	cmd.Flags().BoolVarP(&opts.Slack, "slack", "s", false, "Add slack bot tokens")
	cmd.Flags().BoolVarP(&opts.Telegram, "telegram", "t", false, "Add telegram bot tokens")

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
