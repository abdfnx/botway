package new

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	PlatformChoice int
	Platform       bool
	Ticks      	   int
	Frames         int
	Progress       float64
	Loaded         bool
	Lang           bool
	LangChoice     int
	PM             bool // package manager
	PMCoice        int
	Quitting       bool
}

type tickMsg struct{}
type frameMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Duration(time.Hour.Seconds()), func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}
