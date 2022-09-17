package render

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	token_shared "github.com/abdfnx/botway/internal/pipes/token"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func (m model) Auth() {
	if _, err := os.Stat(constants.RenderConfigFile); os.IsNotExist(err) {
		renderConfigFile := os.WriteFile(constants.RenderConfigFile, []byte("{}"), 0644)

		if renderConfigFile != nil {
			log.Fatal(renderConfigFile)
		}
	}

	email := strings.ReplaceAll(m.inputs[1].Value(), "@", "%40")

	url := fmt.Sprintf("https://api.render.com/v1/owners?name=%s&email=%s&limit=20", m.inputs[0].Value(), email)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+m.inputs[2].Value())

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	id := gjson.Get(string(body), "0.owner.id").String()

	userName, _ := sjson.Set(string(constants.RenderConfig), "user.name", m.inputs[0].Value())
	userEmail, _ := sjson.Set(userName, "user.email", m.inputs[1].Value())
	userId, _ := sjson.Set(userEmail, "user.id", id)
	apiToken, _ := sjson.Set(userId, "user.api_token", m.inputs[2].Value())
	renderProjects, _ := sjson.Set(apiToken, "projects", map[string]string{})

	remove := os.Remove(constants.RenderConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.RenderConfigFile, []byte(renderProjects), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = token_shared.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Render User Name"
			t.Focus()
			t.PromptStyle = token_shared.FocusedStyle
			t.TextStyle = token_shared.FocusedStyle

		case 1:
			t.Placeholder = "Render User Email"
			t.CharLimit = 64

		case 2:
			t.Placeholder = "Render API Token"
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
				m.Auth()

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

func BotwayRenderAuth() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
