package guilds

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/initx"
	token_shared "github.com/abdfnx/botway/internal/pipes/token"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func (m model) AddGuildId() {
	checkBotType := gjson.Get(string(constants.BotwayConfig), "botway.bots."+m.inputs[0].Value()+".type").String()

	if checkBotType != "discord" {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" this command/feature only works with discord bots"))
	} else {
		newGuild, _ := sjson.Set(string(constants.BotwayConfig), "botway.bots."+m.inputs[0].Value()+".guilds."+m.inputs[1].Value()+".server_id", m.inputs[2].Value())

		remove := os.Remove(constants.BotwayConfigFile)

		if remove != nil {
			log.Fatal(remove)
		}

		newBotwayConfig := os.WriteFile(constants.BotwayConfigFile, []byte(newGuild), 0644)

		if newBotwayConfig != nil {
			panic(newBotwayConfig)
		}

		bot_path := gjson.Get(string(constants.BotwayConfig), "botway.bots."+m.inputs[0].Value()+".path").String()
		guildsPath := filepath.Join(bot_path, "config", "guilds.json")
		guildsFile, err := os.ReadFile(guildsPath)

		if err != nil {
			panic(err)
		}

		addGuild, _ := sjson.Set(string(guildsFile), "guilds.-1", m.inputs[1].Value())

		removeOldGuilds := os.Remove(guildsPath)

		if removeOldGuilds != nil {
			panic(removeOldGuilds)
		}

		newGuilds := os.WriteFile(guildsPath, []byte(addGuild), 0644)

		if newGuilds != nil {
			panic(newGuilds)
		}

		initx.UpdateConfig()

		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + m.inputs[1].Value() + " server id is added successfully"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render("You can add more server ids by running the same command again"))
	}
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = token_shared.CursorStyle

		switch i {
		case 0:
			t.Placeholder = "Discord Bot Name"
			t.Focus()
			t.PromptStyle = token_shared.FocusedStyle
			t.TextStyle = token_shared.FocusedStyle

		case 1:
			t.Placeholder = "Discord Server Name"

		case 2:
			t.Placeholder = "Discord Server ID"
			t.CharLimit = 18
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
				m.AddGuildId()

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

func BotwayDiscordGuildIdsSetup() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
