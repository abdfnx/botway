package new

import (
	"fmt"

	"github.com/abdfnx/botway/internal/options"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
)

func New(o *options.CommonOptions) {
	m := model{0, false, 10, 0, 0, false, false, 0, false, 0, false}
	p := tea.NewProgram(m)

	opts.BotName = o.BotName

	if err := p.Start(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()

		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	if !m.Platform {
		return updatePlatforms(msg, m)
	}

	if !m.Lang {
		return updateLangs(msg, m)
	}

	if !m.PM {
		return updatePMs(msg, m)
	}

	return buildBot(msg, m, opts.BotName)
}

func (m model) View() string {
	var s string

	if m.Quitting {
		return "\nBye 👋\n\n"
	}

	if !m.Platform {
		s = platformsView(m)
		return indent.String("\n"+s+"\n\n", 2)
	} else if !m.Lang {
		s = langsView(m)
		return indent.String("\n"+s+"\n\n", 2)
	} else if !m.PM {
		s = pmsView(m)
		return indent.String("\n"+s+"\n\n", 2)
	}

	return s
}
