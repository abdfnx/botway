package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	discord_token "github.com/abdfnx/botway/internal/pipes/token/discord"
	"github.com/abdfnx/botway/internal/pipes/token/discord/guilds"
	slack_token "github.com/abdfnx/botway/internal/pipes/token/slack"
	telegram_token "github.com/abdfnx/botway/internal/pipes/token/telegram"
	twitch_token "github.com/abdfnx/botway/internal/pipes/token/twitch"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func TokenCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage your bots tokens",
	}

	cmd.AddCommand(TokenSetCMD())
	cmd.AddCommand(TokenGetCMD())
	cmd.AddCommand(TokenRemoveCMD())
	cmd.AddCommand(TokenAddGuildsCMD())

	return cmd
}

func TokenSetCMD() *cobra.Command {
	opts := &options.TokenAddOptions{
		BotName:  "",
		Discord:  false,
		Slack:    false,
		Telegram: false,
		Twitch:   false,
	}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Create or update the value of a bot token.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				if opts.Discord {
					discord_token.BotwayDiscordTokenSetup(args[0])
				} else if opts.Slack {
					slack_token.BotwaySlackTokenSetup(args[0])
				} else if opts.Telegram {
					telegram_token.BotwayTelegramTokenSetup(args[0])
				} else if opts.Twitch {
					twitch_token.BotwayTwitchTokenSetup(args[0])
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
	cmd.Flags().BoolVarP(&opts.Twitch, "twitch", "w", false, "For twitch bot tokens")

	return cmd
}

func TokenGetCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the value of a bot token.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Println(messageStyle.Render(strings.ToUpper(args[0])+" Bot Token ==> ") + gjson.Get(string(constants.BotwayConfig), "botway.bots."+args[0]+".bot_token").String())
			} else {
				fmt.Println("Bot Name is required")
			}
		},
	}

	return cmd
}

func TokenRemoveCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   "Remove a bot token.",
		Aliases: []string{"rm", "delete"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				remove, err := sjson.Delete(string(constants.BotwayConfig), "botway.bots."+args[0]+".bot_token")

				if err != nil {
					fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
					fmt.Print(" ")
					panic(constants.FAIL_FOREGROUND.Render(err.Error()))
				}

				os.Remove(constants.BotwayConfigFile)

				os.WriteFile(constants.BotwayConfigFile, []byte(remove), 0644)

				fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
				fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Token removed successfully"))
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
