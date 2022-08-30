package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit     key.Binding
	Down     key.Binding
	Up       key.Binding
	Left     key.Binding
	Right    key.Binding
	Esc      key.Binding
	SwapView key.Binding
	Open     key.Binding
	PageDown key.Binding
	PageUp   key.Binding
	New      key.Binding
}

func New() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+q", "ctrl+c"),
			key.WithHelp("ctrl+q", "quit"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "move down"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "move up"),
		),
		Left: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("↓", "move down"),
		),
		Right: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("↑", "move up"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("escape", "Reset"),
		),
		SwapView: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "select"),
		),
		Open: key.NewBinding(
			key.WithKeys("ctrl+o"),
			key.WithHelp("ctrl+o", "Open bot project at Host Service"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("pgdown", " ", "f"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup", "b"),
		),
		New: key.NewBinding(
			key.WithKeys("n"),
		),
	}
}
