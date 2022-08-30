package initx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdfnx/botway/constants"
	token_shared "github.com/abdfnx/botway/internal/pipes/token"
	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

func (m model) InitCmd() {
	err := dfs.CreateDirectory(filepath.Join(constants.HomeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(constants.BotwayDirPath)
	viper.SetConfigName("botway")
	viper.SetConfigType("json")

	viper.SetDefault("botway.bots", map[string]string{})
	viper.SetDefault("botway.settings.auto_sync", true)
	viper.SetDefault("botway.settings.check_updates", true)
	viper.SetDefault("botway.bots_names", []string{})
	viper.SetDefault("github.username", m.inputs[0].Value())
	viper.SetDefault("docker.id", m.inputs[1].Value())

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(constants.BotwayConfigFile); err == nil {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Initialization Successful"))
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Initialization Failed, try again"))
	}
}

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = token_shared.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "GitHub Username"
			t.Focus()
			t.PromptStyle = token_shared.FocusedStyle
			t.TextStyle = token_shared.FocusedStyle

		case 1:
			t.Placeholder = "Docker Hub ID"
			t.CharLimit = 64
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

			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.InitCmd()

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

func BotwayInit() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
