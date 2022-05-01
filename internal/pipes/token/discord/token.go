package discord_token

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/tran/dfs"
	"github.com/gookit/config/v2"
	yaml "github.com/gookit/config/v2/yaml"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle   = focusedStyle.Copy()
	noStyle       = lipgloss.NewStyle()
	focusedButton = focusedStyle.Copy().Render("[ Done ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Done"))

	homeDir, err     = dfs.GetHomeDirectory()
	botwayConfigPath = filepath.Join(homeDir, ".botway", "botway.yaml")
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
	botName    string
}

func Generator() string {
	rand.Seed(time.Now().Unix())
    charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP1234567890")

    var output strings.Builder

	for i := 0; i < 32; i++ {
        random := rand.Intn(len(charSet))
        randomChar := charSet[random]
        output.WriteRune(randomChar)
    }

	return output.String()
}

var userSecret = Generator()

func EncryptTokens(token, id string) (string, string) {
	var encryptToken = func () string {
		text := []byte(token)
		key := []byte(userSecret)

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
	botwayConfig, err := ioutil.ReadFile(botwayConfigPath)
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
	bc.Set("botway.bots." + m.botName + ".bot_app_id", id)
	bc.Set("botway.bots." + m.botName + ".path", path)
	bc.Set("botway.bots." + m.botName + ".type", botType)

	remove := os.Remove(botwayConfigPath)

	if remove != nil {
        log.Fatal(remove)
    }

	newBotConfig := os.WriteFile(botwayConfigPath, []byte(string(tools.ToYaml(bc.Data()))), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + m.botName + " Discord tokens're added successfully"))
	fmt.Println("Your Secret key -> " + userSecret + " , Save it in save place")
}

func initialModel(botName string) model {
	m := model{
		inputs: make([]textinput.Model, 2),
		botName: botName,
	}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle

		switch i {
			case 0:
				t.Placeholder = "Discord Bot Token"
				t.Focus()
				t.PromptStyle = focusedStyle
				t.TextStyle = focusedStyle
				t.CharLimit = 65

			case 1:
				t.Placeholder = "Discord App (or Client) ID"
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

func BotwayDiscordTokenSetup(botName string) {
	if err := tea.NewProgram(initialModel(botName)).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
