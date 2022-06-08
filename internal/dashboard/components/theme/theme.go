package theme

import (
	"github.com/abdfnx/botway/constants"
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	LightGray lipgloss.AdaptiveColor
	Blue      lipgloss.AdaptiveColor
	Green     lipgloss.AdaptiveColor
	Red       lipgloss.AdaptiveColor
	Orange    lipgloss.AdaptiveColor

	DefaultTextColor lipgloss.AdaptiveColor
	ErrorColor       lipgloss.AdaptiveColor

	SelectedListItemColor   lipgloss.AdaptiveColor
	UnselectedListItemColor lipgloss.AdaptiveColor
	ActiveBoxBorderColor    lipgloss.AdaptiveColor
	InactiveBoxBorderColor  lipgloss.AdaptiveColor
}

type appColors struct {
	white                  string
	black                  string
	lightGray              string
	green                  string
	orange                 string
	slightlyBrighterOrange string
	red                    string
	slightlyBrighterRed    string
	blue                   string
}

var Colors = appColors{
	white:                  "#fff",
	black:                  "#202124",
	lightGray:              "#ced3d9",
	green:                  "#00bf0d",
	orange:                 "#cf8611",
	slightlyBrighterOrange: "#ffff00",
	red:                    "#cc241d",
	slightlyBrighterRed:    "#ff0a00",
	blue:                   constants.PRIMARY_COLOR_HEX,
}

// themeMap represents the mapping of different themes.
var themeMap = map[string]Theme{
	"default": {
		LightGray: lipgloss.AdaptiveColor{Dark: Colors.lightGray, Light: Colors.lightGray},
		Blue:      lipgloss.AdaptiveColor{Dark: Colors.blue, Light: Colors.blue},
		Green:     lipgloss.AdaptiveColor{Dark: Colors.green, Light: Colors.green},
		Red:       lipgloss.AdaptiveColor{Dark: Colors.slightlyBrighterRed, Light: Colors.red},
		Orange:    lipgloss.AdaptiveColor{Dark: Colors.slightlyBrighterOrange, Light: Colors.orange},

		DefaultTextColor: lipgloss.AdaptiveColor{Dark: Colors.white, Light: Colors.black},
		ErrorColor:       lipgloss.AdaptiveColor{Dark: Colors.slightlyBrighterRed, Light: Colors.red},

		SelectedListItemColor:   lipgloss.AdaptiveColor{Dark: Colors.white, Light: Colors.black},
		UnselectedListItemColor: lipgloss.AdaptiveColor{Dark: Colors.black, Light: Colors.white},

		ActiveBoxBorderColor:   lipgloss.AdaptiveColor{Dark: Colors.blue, Light: Colors.blue},
		InactiveBoxBorderColor: lipgloss.AdaptiveColor{Dark: Colors.white, Light: Colors.black},
	},
}

var AppTheme = themeMap["default"]

func SetTheme(theme string) {
	if val, ok := themeMap[theme]; ok {
		AppTheme = val
	}
}
