package remove

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/initx"
	token_shared "github.com/abdfnx/botway/internal/pipes/token"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var opts = options.CommonOptions{
	BotName: "",
}

func initialModel(o *options.CommonOptions) model {
	m := model{
		inputs:  make([]textinput.Model, 2),
		botName: o.BotName,
	}

	opts.BotName = o.BotName

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = token_shared.CursorStyle

		switch i {
		case 0:
			t.Placeholder = m.botName
			t.Focus()
			t.PromptStyle = token_shared.FocusedStyle
			t.TextStyle = token_shared.FocusedStyle

		case 1:
			t.Placeholder = "delete my bot"
			t.CharLimit = 13
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) RemoveCmd() {
	if constants.Berr != nil {
		panic(constants.Berr)
	}

	// remove project dir
	botPath := gjson.Get(string(constants.BotwayConfig), "botway.bots."+m.botName+".path").String()

	if botPath == "" {
		panic(errors.New("bot path not found"))
	}

	err := os.RemoveAll(botPath)

	if err != nil {
		panic(err)
	}

	botsNamesList := gjson.Get(string(constants.BotwayConfig), "botway.bots_names")
	i := ""

	botsNamesList.ForEach(func(in, value gjson.Result) bool {
		if value.String() == m.botName {
			i = in.String()
		}

		return true
	})

	deleteBotFromBotsNamesList, _ := sjson.Delete(string(constants.BotwayConfig), "botway.bots_names."+i)
	deleteBot, _ := sjson.Delete(deleteBotFromBotsNamesList, "botway.bots."+m.botName)

	remove := os.Remove(constants.BotwayConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.BotwayConfigFile, []byte(deleteBot), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}

	initx.UpdateConfig()

	if _, err := os.Stat(constants.BotwayConfigFile); err == nil {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + m.botName + " Removed Successfully"))
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Failed, try again"))
	}
}

type model struct {
	focusIndex int
	inputs     []textinput.Model
	botName    string
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

			if s == "enter" && m.inputs[0].Value() == m.botName && m.inputs[1].Value() == "delete my bot" {
				m.RemoveCmd()

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
	b.WriteString(constants.FAIL_BACKGROUND.Render("WARNING"))
	b.WriteString(constants.FAIL_FOREGROUND.Render(" This action is not reversible. Please be certain."))
	b.WriteString("\n")
	b.WriteString("\n")

	for i := range m.inputs {
		if i == 0 {
			b.WriteString("Enter the bot name to continue")
		}

		if i == 1 {
			b.WriteString("\n")
			b.WriteString("To verify, type " + token_shared.BoldStyle.Render("delete my bot") + " below")
		}

		b.WriteString("\n")
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

func Remove(o *options.CommonOptions) {
	if err := tea.NewProgram(initialModel(o)).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
