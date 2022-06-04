package common

import (
	"fmt"
	"log"

	"github.com/abdfnx/botway/internal/dashboard/components/theme"
	"github.com/charmbracelet/lipgloss"
)

func LogCommand(msg string) {
	log.Print(lipgloss.NewStyle().Foreground(theme.AppTheme.Blue).Render(msg))
}

func LogCommandf(format string, a ...interface{}) {
	LogCommand(fmt.Sprintf(format, a...))
}

func LogSuccess(msg string) {
	log.Print(lipgloss.NewStyle().Foreground(theme.AppTheme.Green).Render(msg))
}

func LogSuccessf(format string, a ...interface{}) {
	LogSuccess(fmt.Sprintf(format, a...))
}

func LogWarning(msg string) {
	log.Print(lipgloss.NewStyle().Foreground(theme.AppTheme.Orange).Render(msg))
}

func LogWarningf(format string, a ...interface{}) {
	LogWarning(fmt.Sprintf(format, a...))
}

func LogError(msg string) {
	log.Print(lipgloss.NewStyle().Foreground(theme.AppTheme.Red).Render(msg))
}

func LogErrorf(format string, a ...interface{}) {
	LogError(fmt.Sprintf(format, a...))
}
