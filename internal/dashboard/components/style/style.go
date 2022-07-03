package style

import (
	"github.com/abdfnx/botway/internal/dashboard/components"
	"github.com/abdfnx/botway/internal/dashboard/components/theme"
	"github.com/charmbracelet/lipgloss"
)

var (
	SecondaryVP = lipgloss.NewStyle().
			PaddingLeft(components.BoxPadding).
			PaddingRight(components.BoxPadding).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.AppTheme.InactiveBoxBorderColor)

	CommandVP = lipgloss.NewStyle().
			PaddingLeft(components.BoxPadding).
			PaddingRight(components.BoxPadding).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.AppTheme.InactiveBoxBorderColor)

	SplashVP = lipgloss.NewStyle().
			PaddingLeft(components.BoxPadding).
			PaddingRight(components.BoxPadding).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.AppTheme.ActiveBoxBorderColor).
			Align(lipgloss.Center)

	PrimaryTitle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).
			Height(3).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 0)

	SecondaryTitle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).
			Height(3).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 0)

	PaginatorActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")

	PaginatorInactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")

	ListSelected = lipgloss.NewStyle().
			Background(theme.AppTheme.SelectedListItemColor).
			Foreground(theme.AppTheme.UnselectedListItemColor)

	AbstractInfo = lipgloss.NewStyle().
			Bold(false).
			Align(lipgloss.Center).
			Height(3).
			Padding(1, 2)

	KeyStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Bold(true).
			Padding(0, 2)

	ValueStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Padding(0, 2)

	StatusBar = lipgloss.NewStyle().
			Height(components.StatusBarHeight)
)
