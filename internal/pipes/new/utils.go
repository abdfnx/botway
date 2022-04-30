package new

import (
	"fmt"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	"github.com/charmbracelet/lipgloss"
)

var (
	prim   = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	subtle = lipgloss.NewStyle().Foreground(constants.GRAY_COLOR)
	dot    = lipgloss.NewStyle().Foreground(constants.GRAY_COLOR).SetString(" • ").String()
	opts   = options.NewOptions{
		BotName: "",
	}
)

func checkbox(label string, checked bool) string {
	if checked {
		return prim.Render("[✔] " + label)
	}

	return fmt.Sprintf("[ ] %s", label)
}

func BotType(m model) string {
	if m.PlatformChoice == 0 {
		return "discord"
	} else if m.PlatformChoice == 1 {
		return "telegram"
	} else if m.PlatformChoice == 2 {
		return "messenger"
	} else if m.PlatformChoice == 3 {
		return "slack"
	}

	return ""
}

func BotLang(m model) string {
	if m.LangChoice == 0 {
		return "python"
	} else if m.LangChoice == 1 {
		return "go"
	} else if m.LangChoice == 2 {
		return "nodejs"
	} else if m.LangChoice == 3 {
		return "ruby"
	} else if m.LangChoice == 4 {
		return "rust"
	} else if m.LangChoice == 5 {
		return "deno"
	}

	return ""
}

func BotPM(m model) string {
	if m.LangChoice == 0 && m.PMCoice == 0 {
		return "pip"
	} else if m.LangChoice == 0 && m.PMCoice == 1 {
		return "pipenv"
	} else if m.LangChoice == 1  {
		return "go mod"
	} else if m.LangChoice == 2 && m.PMCoice == 0 {
		return "npm"
	} else if m.LangChoice == 2 && m.PMCoice == 1 {
		return "yarn"
	} else if m.LangChoice == 2 && m.PMCoice == 2 {
		return "pnpm"
	} else if m.LangChoice == 3 {
		return "bundler"
	} else if m.LangChoice == 4 && m.PMCoice == 0 {
		return "cargo"
	} else if m.LangChoice == 4 && m.PMCoice == 1 {
		return "fleet"
	} else if m.LangChoice == 5 {
		return "deno"
	}

	return ""
}
