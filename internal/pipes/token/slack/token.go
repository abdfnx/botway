package slack_token

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	token_shared "github.com/abdfnx/botway/internal/pipes/token"
	"github.com/abdfnx/botway/tools"
	"github.com/gookit/config/v2"
	yaml "github.com/gookit/config/v2/yaml"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
	botName    string
}

func EncryptTokens(token, id string) (string, string) {
	var encryptToken = func () string {
		text := []byte(token)
		key := []byte(token_shared.UserSecret)

		c, err := aes.NewCipher(key)

		if err != nil {
			fmt.Println(err)
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			fmt.Println(err)
		}

		nonce := make([]byte, gcm.NonceSize())

		if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
			fmt.Println(err)
		}

		return fmt.Sprintf("%x", gcm.Seal(nonce, nonce, text, nil))
	}

	var encryptId = func () string {
		hash := sha256.Sum256([]byte(id))

		return fmt.Sprintf("%x", hash)
	}

	return encryptToken(), encryptId()
}

func (m model) AddToken() {
	botwayConfig, err := ioutil.ReadFile(token_shared.BotwayConfigPath)
	token, id := EncryptTokens(m.inputs[0].Value(), m.inputs[1].Value())

	if err != nil {
		panic(err)
	}

	bc := config.New(".")
	bc.AddDriver(yaml.Driver)
	bcp := bc.LoadSources(config.Yaml, botwayConfig)

	if bcp != nil {
		panic(bcp)
	}

	path := bc.Get("botway.bots." + m.botName + ".path")
	botType := bc.Get("botway.bots." + m.botName + ".type")

	bc.Set("botway.bots." + m.botName + ".bot_token", token)
	bc.Set("botway.bots." + m.botName + ".bot_app_token", id)
	bc.Set("botway.bots." + m.botName + ".path", path)
	bc.Set("botway.bots." + m.botName + ".type", botType)

	remove := os.Remove(token_shared.BotwayConfigPath)

	if remove != nil {
        log.Fatal(remove)
    }

	newBotConfig := os.WriteFile(token_shared.BotwayConfigPath, []byte(string(tools.ToYaml(bc.Data()))), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + m.botName + " Slack tokens're added successfully"))
	fmt.Println("Your Secret key -> " + token_shared.BoldStyle.Render(token_shared.UserSecret) + " Keep it in a safe place")
}

func initialModel(botName string) model {
	m := model{
		inputs: make([]textinput.Model, 2),
		botName: botName,
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = token_shared.CursorStyle

		switch i {
			case 0:
				t.Placeholder = "Slack Token"
				t.Focus()
				t.PromptStyle = token_shared.FocusedStyle
				t.TextStyle = token_shared.FocusedStyle

			case 1:
				t.Placeholder = "Slack App Token"
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

func BotwaySlackTokenSetup(botName string) {
	if err := tea.NewProgram(initialModel(botName)).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
