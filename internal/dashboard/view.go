package dashboard

import (
	"github.com/abdfnx/botway/internal/dashboard/components"
	"github.com/abdfnx/botway/internal/dashboard/components/style"
	"github.com/abdfnx/botway/internal/dashboard/components/theme"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

func (b Bubble) View() string {
	var body string

	splashStyle := style.SplashVP.Render

	if b.outOfBounds() {
		title := style.PrimaryTitle.
			Width(b.bubbles.splashPaginator.Width + 2).
			Render("Out of Bounds")

		body = connectVert(
			title,
			splashStyle(b.getOutOfBoundsView()),
		)

		return connectVert(
			body,
			b.statusBarView(),
		)
	}

	switch b.activeBox {
	default:
		var primaryBox string
		var secondaryBox string

		// set colors
		primaryBoxBorderColor := theme.AppTheme.InactiveBoxBorderColor
		secondaryBoxBorderColor := theme.AppTheme.InactiveBoxBorderColor

		primaryTitle := b.styleTitle("Bot List")
		secondaryTitle := b.styleSecondaryTitle("Botway")

		switch b.activeBox {
		case components.BotListView:
			primaryBoxBorderColor = theme.AppTheme.ActiveBoxBorderColor

		case components.BotInfoView:
			secondaryBoxBorderColor = theme.AppTheme.ActiveBoxBorderColor
		}

		pageStyle := lipgloss.NewStyle().
			PaddingLeft(components.BoxPadding).
			PaddingRight(components.BoxPadding).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryBoxBorderColor).
			Align(lipgloss.Center).Render

		primaryBox = connectVert(
			primaryTitle,
			pageStyle(b.bubbles.primaryPaginator.GetContent()),
		)

		b.bubbles.secondaryViewport.Style = lipgloss.NewStyle().
			PaddingLeft(components.BoxPadding).
			PaddingRight(components.BoxPadding).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondaryBoxBorderColor)
		secondaryBox = connectVert(
			secondaryTitle,
			b.bubbles.secondaryViewport.View(),
			b.bubbles.commandViewport.View(),
		)

		body = connectHorz(
			primaryBox,
			secondaryBox,
		)
	}

	return connectVert(
		body,
		b.statusBarView(),
	)
}

func (b Bubble) styleTitle(s string) string {
	s = trunc(
		s,
		b.bubbles.primaryPaginator.Width-2,
	)

	return style.PrimaryTitle.
		Width(b.bubbles.primaryPaginator.Width + 2).
		Render(s)
}

func (b Bubble) styleSecondaryTitle(s string) string {
	s = trunc(
		s,
		b.bubbles.secondaryViewport.Width-2,
	)

	return style.SecondaryTitle.
		Width(b.bubbles.secondaryViewport.Width + 2).
		Render(s)
}

func (b Bubble) drawKV(k, v string, color bool) string {
	keyStyle := style.KeyStyle.Width((b.bubbles.secondaryViewport.Width / 3))

	valueStyle := style.ValueStyle

	if color {
		return connectHorz(
			keyStyle.Copy().
				Render(k),
			style.ValueStyle.Copy().
				Foreground(theme.AppTheme.UnselectedListItemColor).
				Background(theme.AppTheme.SelectedListItemColor).
				Render(v))
	} else {
		return connectHorz(keyStyle.Render(k), valueStyle.Render(v))
	}
}

func (b Bubble) drawHelpKV(k, v string) string {
	keyStyle := style.KeyStyle.Copy().Width(11)

	valueStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Faint(true).
		PaddingRight(2)

	return connectHorz(keyStyle.Render(k), valueStyle.Render(v))
}

func (b Bubble) outOfBounds() bool {
	if b.width < components.MinWidth || b.height < components.MinHeight {
		return true
	}

	return false
}

func connectHorz(strs ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, strs...)
}

func connectVert(strs ...string) string {
	return lipgloss.JoinVertical(lipgloss.Top, strs...)
}

func trunc(s string, i int) string {
	return truncate.StringWithTail(
		s,
		uint(i),
		components.EllipsisStyle,
	)
}

func styleWidth(i int) lipgloss.Style {
	return lipgloss.NewStyle().Width(i)
}
