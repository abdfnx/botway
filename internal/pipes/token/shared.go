package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/lipgloss"
)

var (
	FocusedStyle  = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	BlurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	BoldStyle     = lipgloss.NewStyle().Bold(true)
	CursorStyle   = FocusedStyle.Copy()
	NoStyle       = lipgloss.NewStyle()
	FocusedButton = FocusedStyle.Copy().Render("[ Done ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Done"))

	HomeDir, _       = dfs.GetHomeDirectory()
	BotwayConfigPath = filepath.Join(HomeDir, ".botway", "botway.json")
	UserSecret 	     = Generator()
)

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

func EncryptTokens(token, id string) (string, string) {
	var encryptToken = func () string {
		text := []byte(token)
		key := []byte(UserSecret)

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
