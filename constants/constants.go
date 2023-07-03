package constants

import (
	"os"
	"path/filepath"

	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/lipgloss"
)

var (
	PRIMARY_COLOR_HEX = "#1d4ed8"
	PRIMARY_COLOR     = lipgloss.Color(PRIMARY_COLOR_HEX)
	CYAN_COLOR        = lipgloss.Color("#00FFFF")
	GREEN_COLOR       = "#04B575"
	RED_COLOR         = "#FF4141"
	YELLOW_COLOR      = lipgloss.Color("178")
	GRAY_COLOR        = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}

	BOLD = lipgloss.NewStyle().Bold(true)

	SUCCESS_BACKGROUND = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFF")).
				Background(lipgloss.Color(GREEN_COLOR)).
				PaddingLeft(1).
				PaddingRight(1)
	FAIL_BACKGROUND = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(RED_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)
	INFO_BACKGROUND = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(CYAN_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)
	WARN_BACKGROUND = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(YELLOW_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)
	PRIMARY_FOREGROUND = lipgloss.NewStyle().Bold(true).Foreground(PRIMARY_COLOR)
	SUCCESS_FOREGROUND = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(GREEN_COLOR))
	FAIL_FOREGROUND    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(RED_COLOR))
	INFO_FOREGROUND    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(CYAN_COLOR))
	COMMAND_FOREGROUND = lipgloss.NewStyle().Bold(true).Foreground(GRAY_COLOR)
	WARN_FOREGROUND    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(YELLOW_COLOR))
	HEADING            = lipgloss.NewStyle().Foreground(lipgloss.Color(PRIMARY_COLOR)).SetString("==> ").String()

	HomeDir, _             = dfs.GetHomeDirectory()
	BotwayDirPath          = filepath.Join(HomeDir, ".botway")
	BotwayConfigFile       = filepath.Join(BotwayDirPath, "botway.json")
	BotwayConfig, Berr     = os.ReadFile(BotwayConfigFile)
	BWDBConfigFile         = filepath.Join(BotwayDirPath, "bwdb.json")
	BWDBConfig, Derr       = os.ReadFile(BWDBConfigFile)
	BotConfig, Oerr        = os.ReadFile(".botway.yaml")
	BotComposeConfig, Cerr = os.ReadFile(".botway-compose.yaml")
	Guilds, Gerr           = os.ReadFile(filepath.Join("config", "guilds.json"))

	RailwayConfigFile   = filepath.Join(HomeDir, ".botway", "railway-config.json")
	RailwayConfig, Rerr = os.ReadFile(RailwayConfigFile)

	RenderConfigFile   = filepath.Join(HomeDir, ".botway", "render-config.json")
	RenderConfig, Nerr = os.ReadFile(RenderConfigFile)

	RAIL_PORT = 4411
)
