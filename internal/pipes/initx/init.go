package initx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func (m model) InitCmd() {
	var err error

	homeDir, err := dfs.GetHomeDirectory()

	if err != nil {
		log.Fatal(err)
	}

	err = dfs.CreateDirectory(filepath.Join(homeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	botwayDirPath := ""

	if runtime.GOOS == "windows" {
		botwayDirPath = `$HOME\\.botway`
	} else {
		botwayDirPath = `$HOME/.botway`
	}

	viper.AddConfigPath(botwayDirPath)
	viper.SetConfigName("botway")
	viper.SetConfigType("yaml")

	viper.SetDefault("botway.bots", "")
	viper.SetDefault("user.github_username", m.inputs[0].Value())
	viper.SetDefault("user.docker_id", m.inputs[1].Value())

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

	botwayConfigFile := filepath.Join(homeDir, ".botway", "botway.yaml")

	if _, err := os.Stat(botwayConfigFile); err == nil {
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
	cursorMode textinput.CursorMode
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
			case 0:
				t.Placeholder = "GitHub Username"
				t.Focus()
				t.PromptStyle = focusedStyle
				t.TextStyle = focusedStyle

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
							m.inputs[i].PromptStyle = focusedStyle
							m.inputs[i].TextStyle = focusedStyle
							continue
						}

						m.inputs[i].Blur()
						m.inputs[i].PromptStyle = noStyle
						m.inputs[i].TextStyle = noStyle
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

	button := &blurredButton

	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
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
