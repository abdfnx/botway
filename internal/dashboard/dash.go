package dashboard

import (
	"github.com/abdfnx/botway/internal/dashboard/components"
	"github.com/abdfnx/botway/internal/dashboard/components/keymap"
	"github.com/abdfnx/botway/internal/dashboard/components/style"
	"github.com/abdfnx/botway/internal/dashboard/components/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type ErrorMsg error

type Bubble struct {
	bubbles       Bubbles
	keyMap        keymap.KeyMap
	nav           Nav
	ready         bool
	activeBox     int
	lastActiveBox int
	width         int
	height        int
}

type Bubbles struct {
	primaryPaginator  Paginator
	splashPaginator   Paginator
	secondaryViewport viewport.Viewport
	commandViewport   viewport.Viewport
}

type Nav struct {
	listCursorHide bool
	listCursor     int
	boolCursor     bool
}

func InitialModel() Bubble {
	secondaryVP := viewport.NewViewport(0, 0)
	secondaryVP.Style = style.SecondaryVP

	commandVP := viewport.NewViewport(0, 0)
	commandVP.Style = style.CommandVP

	pages := NewPaginator()
	pages.SetTotalPages(1)

	splashPages := NewPaginator()
	splashPages.SetTotalPages(1)

	nav := Nav{
		listCursorHide: true,
		boolCursor:     false,
	}

	bubs := Bubbles{
		secondaryViewport: secondaryVP,
		commandViewport:   commandVP,
		primaryPaginator:  pages,
		splashPaginator:   splashPages,
	}

	return Bubble{
		bubbles:   bubs,
		ready:     false,
		nav:       nav,
		activeBox: components.BotListView,
		keyMap:    keymap.New(),
	}
}

func (b Bubble) Init() tea.Cmd {
	return nil
}
