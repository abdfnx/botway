package twitch_token

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/initx"
	token_shared "github.com/abdfnx/botway/internal/pipes/token"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tidwall/sjson"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	botName    string
}

func (m model) AddToken() {
	tokenContent, _ := sjson.Set(string(constants.BotwayConfig), "botway.bots."+m.botName+".bot_token", m.inputs[0].Value())
	clientContent, _ := sjson.Set(tokenContent, "botway.bots."+m.botName+".bot_client_id", m.inputs[1].Value())
	clientSecretContent, _ := sjson.Set(clientContent, "botway.bots."+m.botName+".bot_client_secret", m.inputs[2].Value())

	remove := os.Remove(constants.BotwayConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.BotwayConfigFile, []byte(clientSecretContent), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}

	initx.UpdateConfig()

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + m.botName + " Slack tokens're added successfully"))
	// fmt.Println("Your Secret key -> " + token_shared.BoldStyle.Render(token_shared.UserSecret) + " Keep it in a safe place")
}

func initialModel(botName string) model {
	m := model{
		inputs:  make([]textinput.Model, 3),
		botName: botName,
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = token_shared.CursorStyle

		switch i {
		case 0:
			t.Placeholder = "OAuth Token"
			t.Focus()
			t.PromptStyle = token_shared.FocusedStyle
			t.TextStyle = token_shared.FocusedStyle

		case 1:
			t.Placeholder = "Client ID"

		case 2:
			t.Placeholder = "Client Secret"
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" {
				m.AddToken()

				return m, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))

			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = token_shared.FocusedStyle
					m.inputs[i].TextStyle = token_shared.FocusedStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = token_shared.NoStyle
				m.inputs[i].TextStyle = token_shared.NoStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &token_shared.BlurredButton

	if m.focusIndex == len(m.inputs) {
		button = &token_shared.FocusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func BotwayTwitchTokenSetup(botName string) {
	if err := tea.NewProgram(initialModel(botName)).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
