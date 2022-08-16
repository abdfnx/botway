package new

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	PlatformChoice 	  int
	Platform       	  bool
	Ticks          	  int
	Frames            int
	Progress       	  float64
	Loaded            bool
	Lang           	  bool
	LangChoice     	  int
	PM             	  bool // package manager
	PMChoice       	  int
	HostService    	  bool
	HostServiceChoice int
	Quitting       	  bool
}

type frameMsg struct{}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}
