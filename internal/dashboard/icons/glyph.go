// forked from: https://github.com/knipferrc/teacup/blob/main/icons/glyphs.go

package icons

import "fmt"

type IconInfo struct {
	icon  string
	color [3]uint8
}

func (i *IconInfo) GetGlyph() string {
	return i.icon
}

func (i *IconInfo) GetColor(f uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", i.color[0], i.color[1], i.color[2])
}

var IconSet = map[string]*IconInfo{
	"go":         {icon: "\ufcd1", color: [3]uint8{32, 173, 194}},
	"python":     {icon: "\uf81f", color: [3]uint8{52, 102, 143}},
	"ruby":       {icon: "\ue739", color: [3]uint8{229, 61, 58}},
	"rust":       {icon: "\ue7a8", color: [3]uint8{250, 111, 66}},
	"nodejs":     {icon: "\ue74e", color: [3]uint8{255, 202, 61}},
	"typescript": {icon: "\ue628", color: [3]uint8{3, 136, 209}},
	"deno":       {icon: "\ue628", color: [3]uint8{3, 136, 209}},
	"csharp":     {icon: "\uf81a", color: [3]uint8{2, 119, 189}},
	"dart":       {icon: "\ue798", color: [3]uint8{87, 182, 240}},
	"java":       {icon: "\uf675", color: [3]uint8{244, 68, 62}},
	"kotlin":     {icon: "\ue70e", color: [3]uint8{139, 195, 74}},
	"c":          {icon: "\ufb70", color: [3]uint8{2, 119, 189}},
	"cpp":        {icon: "\ufb71", color: [3]uint8{2, 119, 189}},
	"nim":        {icon: "\uf6a4", color: [3]uint8{255, 202, 61}},
	"lua":        {icon: "\ue620", color: [3]uint8{66, 165, 245}},
	"swift":      {icon: "\ufbe3", color: [3]uint8{249, 95, 63}},
	"other":      {icon: "\ufba7", color: [3]uint8{249, 89, 63}},
}
