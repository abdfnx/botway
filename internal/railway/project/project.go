package project

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var proID string

type tickMsg struct{}
type errMsg error

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 36
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			os.Exit(0)

			return m, tea.Quit

		case tea.KeyEnter:
			id := m.textInput.Value()

			if len(strings.TrimSpace(id)) < 1 {
				m.textInput.Placeholder = "Project ID is required"
				m.textInput.SetValue("")
			} else if len(strings.TrimSpace(id)) > 36 {
				m.textInput.Placeholder = "Project ID is too long"
				m.textInput.SetValue("")
			} else if strings.Contains(id, " ") {
				m.textInput.Placeholder = "Project ID cannot contain spaces"
				m.textInput.SetValue("")
			} else if len(strings.TrimSpace(id)) < 36 {
				m.textInput.Placeholder = "Project ID is too short, please enter a valid project id"
				m.textInput.SetValue("")
			} else {
				proID = id
				return m, tea.Quit
			}
		}

	case errMsg:
		m.err = msg

		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	projectID := lipgloss.NewStyle().Bold(true).SetString("Enter your project id:").String()

	return fmt.Sprintf(
		"%s%s\n",
		projectID,
		m.textInput.View(),
	) + "\n"
}

func Project() (string, error) {
	err := tea.NewProgram(initialModel()).Start()

	return proID, err
}
