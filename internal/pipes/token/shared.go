package token

import (
	"fmt"
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

	HomeDir, err     = dfs.GetHomeDirectory()
	BotwayConfigPath = filepath.Join(HomeDir, ".botway", "botway.yaml")
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
