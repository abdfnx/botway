package discord_token

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle   = focusedStyle.Copy()
	noStyle       = lipgloss.NewStyle()
	focusedButton = focusedStyle.Copy().Render("[ Done ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Done"))
	opts          = options.TokenAddOptions{
		BotName: "",
	}
)

func Generator() string {
	rand.Seed(time.Now().Unix())
    charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP1234567890")

    var output strings.Builder

	for i := 0; i < 33; i++ {
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

		return string(gcm.Seal(nonce, nonce, text, nil))
	}

	var encryptId = func () string {
		hash := sha256.Sum256([]byte(id))

		return fmt.Sprintf("%x", hash)
	}

	return encryptToken(), encryptId()
}

func (m model) AddToken() {
	homeDir, err := dfs.GetHomeDirectory()
	token, id := EncryptTokens(m.inputs[0].Value(), m.inputs[1].Value())

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

	viper.SetDefault("botway.bots." + opts.BotName + ".bot_token", token)
	viper.SetDefault("botway.bots." + opts.BotName + ".bot_app_id", id)

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
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + opts.BotName + " Discord tokens're added successfully"))
		fmt.Println("Your Secret key -> " + userSecret)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Failed, try again"))
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
				t.Placeholder = "Discord Bot Token"
				t.Focus()
				t.PromptStyle = focusedStyle
				t.TextStyle = focusedStyle

			case 1:
				t.Placeholder = "Discord App (or Client) ID"
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

func BotwayDiscordTokenSetup() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
