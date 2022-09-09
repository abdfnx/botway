package dashboard

import (
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/dashboard/components"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tidwall/gjson"
)

// Handles all key press events
func (b *Bubble) handleKeys(msg tea.KeyMsg) tea.Cmd {
	var cmds []tea.Cmd

	if b.outOfBounds() {
		if key.Matches(msg, b.keyMap.Quit) {
			return tea.Quit
		}

		return tea.Batch(cmds...)
	}

	switch {
	// Quit
	case key.Matches(msg, b.keyMap.Quit):
		return tea.Quit

	// Down
	case key.Matches(msg, b.keyMap.Down):
		b.scrollView("down")

	// Up
	case key.Matches(msg, b.keyMap.Up):
		b.scrollView("up")

	// Left
	case key.Matches(msg, b.keyMap.Left):
		b.scrollView("left")

	// Right
	case key.Matches(msg, b.keyMap.Right):
		b.scrollView("right")

	// Escape
	case key.Matches(msg, b.keyMap.Esc):
		cmds = append(cmds, b.resetView())

	// Swap view
	case key.Matches(msg, b.keyMap.SwapView):
		switch b.activeBox {
		case components.BotListView:
			b.switchActiveView(components.BotInfoView)

		default:
			b.switchActiveView(components.BotListView)
		}

	// Open bot project at Railway or Render
	case key.Matches(msg, b.keyMap.Open):
		if b.botInfo("host_service") == "railway.app" {
			bot_project_id := gjson.Get(string(constants.RailwayConfig), "projects."+strings.ToLower(b.botInfo("path"))+".project").String()

			OpenBrowser("https://railway.app/project/" + bot_project_id)
		} else if b.botInfo("host_service") == "render.com" {
			bot_project_id := gjson.Get(string(constants.RenderConfig), "projects."+b.botInfo("name")+".id").String()

			OpenBrowser("https://dashboard.render.com/web/" + bot_project_id)
		}
	}

	return tea.Batch(cmds...)
}

func (b *Bubble) resetView() tea.Cmd {
	b.nav.boolCursor = false
	b.nav.listCursor = 0
	b.nav.listCursorHide = true
	b.switchActiveView(components.BotListView)
	b.lastActiveBox = components.BotListView

	return nil
}
