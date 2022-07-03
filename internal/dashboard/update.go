package dashboard

import (
	"log"

	"github.com/abdfnx/botway/internal/dashboard/components"
	"github.com/abdfnx/botway/internal/dashboard/components/common"
	tea "github.com/charmbracelet/bubbletea"
)

// Do computations for TUI app
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case ErrorMsg:
		b.ready = true
		common.LogErrorf("error: %v", msg)

	case tea.WindowSizeMsg:
		b.width = msg.Width
		b.height = msg.Height

		b.bubbles.commandViewport.Width = (msg.Width / 3) - b.bubbles.commandViewport.Style.GetHorizontalFrameSize() + 8
		b.bubbles.commandViewport.Height = 7

		b.bubbles.secondaryViewport.Width = (msg.Width / 3) - b.bubbles.secondaryViewport.Style.GetHorizontalFrameSize() + 8
		b.bubbles.secondaryViewport.Height = msg.Height - b.bubbles.commandViewport.Height - components.StatusBarHeight - b.bubbles.secondaryViewport.Style.GetVerticalFrameSize() - 7

		b.bubbles.primaryPaginator.SetWidth(msg.Width - b.bubbles.secondaryViewport.Width - 8)
		b.bubbles.primaryPaginator.SetHeight(b.bubbles.secondaryViewport.Height + b.bubbles.commandViewport.Height)

		b.bubbles.splashPaginator.SetWidth(msg.Width - 4)
		b.bubbles.splashPaginator.SetHeight(msg.Height - components.StatusBarHeight - 9)

	case tea.KeyMsg:
		cmds = append(cmds, b.handleKeys(msg))

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseWheelUp:
			b.scrollView("up")

		case tea.MouseWheelDown:
			b.scrollView("down")
		}

	default:
		log.Printf("%T", msg)
	}

	cmds = append(cmds, b.updateActiveView(msg))

	return b, tea.Batch(cmds...)
}

// Update content for active view
func (b *Bubble) updateActiveView(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	b.bubbles.primaryPaginator.SetTotalPages(int(bots_count))
	b.checkActiveViewPortBounds()

	b.bubbles.commandViewport.SetContent(b.commandView())

	switch b.activeBox {
	case components.BotListView:
		b.bubbles.primaryPaginator.SetContent(b.botListView())
		b.bubbles.secondaryViewport.SetContent(b.botInfoView())

	case components.BotInfoView:
		b.bubbles.primaryPaginator.SetContent(b.botListView())
		b.bubbles.secondaryViewport, cmd = b.bubbles.secondaryViewport.Update(msg)
		b.bubbles.secondaryViewport.SetContent(b.botInfoView())
	}

	return cmd
}

func (b *Bubble) switchActiveView(newView int) {
	b.bubbles.primaryPaginator.GoToStart()
	b.bubbles.splashPaginator.GoToStart()
	b.lastActiveBox = b.activeBox
	b.activeBox = newView
}

// Handles wrapping and button scrolling in the viewport
func (b *Bubble) checkActiveViewPortBounds() {
	switch b.activeBox {
	case components.BotInfoView:
		if b.bubbles.secondaryViewport.AtBottom() {
			b.bubbles.secondaryViewport.GotoBottom()
		} else if b.bubbles.secondaryViewport.AtTop() {
			b.bubbles.secondaryViewport.GotoTop()
		}
	}
}

// Handles mouse scrolling in the viewport
func (b *Bubble) scrollView(dir string) {
	switch dir {
	case "up":
		switch b.activeBox {
		case components.BotListView:
			if b.nav.listCursorHide {
				b.nav.listCursorHide = false
			} else {
				b.bubbles.primaryPaginator.LineUp()
			}

		case components.BotInfoView:
			b.bubbles.secondaryViewport.LineUp(1)
		}

	case "down":
		switch b.activeBox {
		case components.BotListView:
			if b.nav.listCursorHide {
				b.nav.listCursorHide = false
			} else {
				b.bubbles.primaryPaginator.LineDown()
			}

		case components.BotInfoView:
			b.bubbles.secondaryViewport.LineDown(1)
		}

	case "left":
		switch b.activeBox {
		case components.BotListView:
			if b.nav.listCursorHide {
				b.nav.listCursorHide = false
			} else {
				b.bubbles.primaryPaginator.PrevPage()
			}
		}

	case "right":
		switch b.activeBox {
		case components.BotListView:
			if b.nav.listCursorHide {
				b.nav.listCursorHide = !b.nav.listCursorHide
			} else {
				b.bubbles.primaryPaginator.NextPage()
			}
		}

	default:
		log.Panic("Invalid scroll direction: " + dir)
	}
}
