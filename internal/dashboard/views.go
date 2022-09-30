package dashboard

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/config"
	"github.com/abdfnx/botway/internal/dashboard/components/style"
	"github.com/abdfnx/botway/internal/dashboard/components/theme"
	"github.com/abdfnx/botway/internal/dashboard/icons"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var (
	keyStyle = func(b Bubble) lipgloss.Style {
		return style.KeyStyle.Width((b.bubbles.secondaryViewport.Width / 4) + 3)
	}

	valueStyle = style.ValueStyle

	drawKV = func(k, v string, b Bubble) string {
		return connectHorz(keyStyle(b).Render(k), valueStyle.Render(v))
	}

	drawKVColor = func(k, v string, color lipgloss.AdaptiveColor, b Bubble) string {
		return connectHorz(
			keyStyle(b).Render(k),
			valueStyle.Copy().
				Foreground(color).
				Render(v))
	}

	styles = func(b Bubble) lipgloss.Style {
		return lipgloss.NewStyle().Width(b.bubbles.secondaryViewport.Width)
	}
)

func (b Bubble) botListView() string {
	pageStyle := styleWidth(b.bubbles.primaryPaginator.Width).
		Height(b.bubbles.primaryPaginator.PerPage + 1).Render

	pagerStyle := styleWidth(b.bubbles.primaryPaginator.Width).
		Align(lipgloss.Center).Render

	page := ""
	v := ""

	bots.ForEach(func(i, value gjson.Result) bool {
		v = value.String()

		line := trunc(value.String(), b.bubbles.primaryPaginator.Width-2)

		lang := gjson.Get(string(constants.BotwayConfig), "botway.bots."+v+".lang").String()

		if lang == "php" || lang == "crystal" {
			lang = "other"
		}

		icon, color := icons.GetIcon(lang)

		fileIcon := lipgloss.NewStyle().Width(2).Render(fmt.Sprintf("%s%s\033[0m ", color, icon))

		line = fmt.Sprintf("%s %s", line, fileIcon)

		if b.bubbles.primaryPaginator.Cursor == int(i.Int()) && !b.nav.listCursorHide {
			page += style.ListSelected.
				Width(b.bubbles.primaryPaginator.Width).
				Render(line)
		} else {
			page += line
		}

		page += "\n"

		return true
	})

	page = connectVert(
		pageStyle(page),
		pagerStyle(b.bubbles.primaryPaginator.View()),
	)

	return styleWidth(b.bubbles.primaryPaginator.Width).
		Height(b.bubbles.primaryPaginator.Height - 3).
		Render(page)
}

func (b Bubble) botInfoView() string {
	if !b.nav.listCursorHide {
		bot_config := viper.New()
		bot_config.SetConfigType("yaml")

		content, err := os.ReadFile(filepath.Join(b.botInfo("path"), ".botway.yaml"))

		if err != nil {
			panic(err)
		}

		bot_config.ReadConfig(bytes.NewBuffer(content))

		bType := drawKVColor("Bot Type", b.botInfo("type"), theme.AppTheme.Blue, b)

		lang := drawKVColor("Language", b.botInfo("lang"), theme.AppTheme.LightGray, b)

		bPath := trunc(b.botInfo("path"), (b.bubbles.secondaryViewport.Width*2/3)-3)

		bPath = drawKVColor("Path", bPath, theme.AppTheme.Blue, b)

		// bRepo := drawKV("Repo", bot_config.GetString("bot.repo"), b)

		is_bot_tokens_been_set := b.botInfo("token")
		var color lipgloss.AdaptiveColor

		if is_bot_tokens_been_set == "" {
			is_bot_tokens_been_set = "No"
			color = theme.AppTheme.Red
		} else {
			is_bot_tokens_been_set = "Yes"
			color = theme.AppTheme.Green
		}

		tokens := drawKVColor("The Tokens been set?", is_bot_tokens_been_set, color, b)

		return connectVert(
			styles(b).Render(""),
			bType,
			"\n",
			lang,
			bPath,
			"\n",
			tokens,
		)
	}

	return styleWidth(b.bubbles.secondaryViewport.Width).
		Height(b.bubbles.secondaryViewport.Height - 3).
		Render(b.homeView())
}

func (b Bubble) statusBarView() string {
	count := fmt.Sprintf("Bot: %d/%d", b.bubbles.primaryPaginator.GetCursorIndex()+1, int(bots_count))

	count = style.StatusBar.Copy().
		Align(lipgloss.Right).
		PaddingRight(6).
		PaddingLeft(2).
		Width(22).
		Render(count)

	return count
}

func (b Bubble) commandView() string {
	commandStyle := lipgloss.NewStyle().
		Align(lipgloss.Center)

	return commandStyle.Render(b.helpView())
}

func (b Bubble) homeView() string {
	is_logged_in := drawKVColor("Is Logged In ?", "No", theme.AppTheme.Red, b)

	if 1 > 0 {
		is_logged_in = drawKVColor("Is Logged In ?", "Yes", theme.AppTheme.Green, b)
	}

	gh_username := drawKVColor("GitHub Username", config.Get("github.username"), theme.AppTheme.LightGray, b)
	docker_id := drawKVColor("Docker ID", config.Get("docker.id"), theme.AppTheme.Blue, b)
	total_bots := drawKV("Total Bots", fmt.Sprint(bots_count), b)

	style := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Bold(true).
		PaddingLeft(2)

	return connectVert(
		style.Render("\n🤖 Generate, build, handle and deploy your own bot with your\nfavorite language, for Discord, or Telegram, or Slack, or even Twitch\n\n"),
		gh_username,
		docker_id,
		"\n",
		is_logged_in,
		"\n",
		total_bots,
	)
}

func (b Bubble) helpView() string {
	leftColumn := []string{
		b.drawHelpKV("up", "Move up"),
		b.drawHelpKV("down", "Move down"),
		b.drawHelpKV("tab", "Swap windows"),
		b.drawHelpKV("ctrl+o", "Open bot project at Host Service"),
		b.drawHelpKV("esc", "Reset"),
	}

	rightColumn := []string{
		b.drawHelpKV("ctrl+q", "Quit"),
	}

	var content string

	if b.bubbles.commandViewport.Width >= 49 {
		content = connectHorz(connectVert(leftColumn...), connectVert(rightColumn...))
	} else {
		content = connectVert(connectVert(leftColumn...), connectVert(rightColumn...))
	}

	content = lipgloss.NewStyle().
		Margin(1, 0).
		Render(content)

	return styleWidth(b.bubbles.commandViewport.Width).
		Align(lipgloss.Left).
		Padding(0, 1).
		Render(content)
}

func (b Bubble) getOutOfBoundsView() string {
	content := lipgloss.NewStyle().
		Width(b.bubbles.splashPaginator.Width).
		Height(b.bubbles.splashPaginator.Height + 2).
		Align(lipgloss.Center).
		Render("Please adjust window size")

	return content
}
