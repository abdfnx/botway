// forked from: https://github.com/charmbracelet/bubbles/blob/master/viewport/viewport.go
package viewport

import (
	"math"
	"strings"

	"github.com/abdfnx/botway/internal/dashboard/components/keymap"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewViewport(width, height int) (m Viewport) {
	m.Width = width
	m.Height = height
	m.setInitialValues()

	return m
}

type Viewport struct {
	Width  int
	Height int
	KeyMap keymap.KeyMap

	MouseWheelEnabled bool

	MouseWheelDelta int

	YOffset int

	YPosition int

	Style lipgloss.Style

	initialized bool
	lines       []string
}

func (m *Viewport) setInitialValues() {
	m.KeyMap = keymap.New()
	m.MouseWheelEnabled = true
	m.MouseWheelDelta = 3
	m.initialized = true
}

func (m Viewport) Init() tea.Cmd {
	return nil
}

func (m Viewport) AtTop() bool {
	return m.YOffset <= 0
}

func (m Viewport) AtBottom() bool {
	return m.YOffset >= m.maxYOffset()
}

func (m Viewport) PastBottom() bool {
	return m.YOffset > m.maxYOffset()
}

func (m Viewport) ScrollPercent() float64 {
	if m.Height >= len(m.lines) {
		return 1.0
	}

	y := float64(m.YOffset)
	h := float64(m.Height)
	t := float64(len(m.lines) - 1)
	v := y / (t - h)

	return math.Max(0.0, math.Min(1.0, v))
}

func (m *Viewport) SetContent(s string) {
	s = strings.ReplaceAll(s, "\r\n", "\n") // normalize line endings
	m.lines = strings.Split(s, "\n")

	if m.YOffset > len(m.lines)-1 {
		m.GotoBottom()
	}
}

func (m Viewport) maxYOffset() int {
	return max(0, len(m.lines)-m.Height)
}

func (m Viewport) visibleLines() (lines []string) {
	if len(m.lines) > 0 {
		top := max(0, m.YOffset)
		bottom := clamp(m.YOffset+m.Height, top, len(m.lines))
		lines = m.lines[top:bottom]
	}

	return lines
}

func (m Viewport) scrollArea() (top, bottom int) {
	top = max(0, m.YPosition)
	bottom = max(top, top+m.Height)

	if top > 0 && bottom > top {
		bottom--
	}

	return top, bottom
}

func (m *Viewport) SetYOffset(n int) {
	m.YOffset = clamp(n, 0, m.maxYOffset())
}

func (m *Viewport) ViewDown() []string {
	if m.AtBottom() {
		return nil
	}

	m.SetYOffset(m.YOffset + m.Height)

	return m.visibleLines()
}

func (m *Viewport) ViewUp() []string {
	if m.AtTop() {
		return nil
	}

	m.SetYOffset(m.YOffset - m.Height)

	return m.visibleLines()
}

func (m *Viewport) HalfViewDown() (lines []string) {
	if m.AtBottom() {
		return nil
	}

	m.SetYOffset(m.YOffset + m.Height/2)

	return m.visibleLines()
}

func (m *Viewport) HalfViewUp() (lines []string) {
	if m.AtTop() {
		return nil
	}

	m.SetYOffset(m.YOffset - m.Height/2)

	return m.visibleLines()
}

func (m *Viewport) LineDown(n int) (lines []string) {
	if m.AtBottom() || n == 0 {
		return nil
	}

	m.SetYOffset(m.YOffset + n)

	return m.visibleLines()
}

func (m *Viewport) LineUp(n int) (lines []string) {
	if m.AtTop() || n == 0 {
		return nil
	}

	m.SetYOffset(m.YOffset - n)

	return m.visibleLines()
}

func (m *Viewport) GotoTop() (lines []string) {
	if m.AtTop() {
		return nil
	}

	m.SetYOffset(0)

	return m.visibleLines()
}

func (m *Viewport) GotoBottom() (lines []string) {
	m.SetYOffset(m.maxYOffset())
	return m.visibleLines()
}

func ViewDown(m Viewport, lines []string) tea.Cmd {
	if len(lines) == 0 {
		return nil
	}

	top, bottom := m.scrollArea()

	return tea.ScrollDown(lines, top, bottom)
}

func ViewUp(m Viewport, lines []string) tea.Cmd {
	if len(lines) == 0 {
		return nil
	}

	top, bottom := m.scrollArea()

	return tea.ScrollUp(lines, top, bottom)
}

func (m Viewport) Update(msg tea.Msg) (Viewport, tea.Cmd) {
	var cmd tea.Cmd

	m, cmd = m.updateAsModel(msg)

	return m, cmd
}

func (m Viewport) updateAsModel(msg tea.Msg) (Viewport, tea.Cmd) {
	if !m.initialized {
		m.setInitialValues()
	}

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.PageDown):
			m.ViewDown()

		case key.Matches(msg, m.KeyMap.PageUp):
			m.ViewUp()

		case key.Matches(msg, m.KeyMap.Down):
			m.LineDown(1)

		case key.Matches(msg, m.KeyMap.Up):
			m.LineUp(1)
		}

	case tea.MouseMsg:
		if !m.MouseWheelEnabled {
			break
		}

		switch msg.Type {
		case tea.MouseWheelUp:
			m.LineUp(m.MouseWheelDelta)

		case tea.MouseWheelDown:
			m.LineDown(m.MouseWheelDelta)
		}
	}

	return m, cmd
}

func (m Viewport) View() string {
	lines := m.visibleLines()

	extraLines := ""

	if len(lines) < m.Height {
		extraLines = strings.Repeat("\n", max(0, m.Height-len(lines)))
	}

	return m.Style.Copy().
		UnsetWidth().
		UnsetHeight().
		Render(strings.Join(lines, "\n") + extraLines)
}

func clamp(v, low, high int) int {
	if high < low {
		low, high = high, low
	}

	return min(high, max(low, v))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
