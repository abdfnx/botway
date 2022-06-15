package new

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/botway/internal/dashboard/icons"
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/new/pm"
	"github.com/abdfnx/resto/core/api"
	table "github.com/calyptia/go-bubble-table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"golang.org/x/term"
)

var (
	styleDoc = lipgloss.NewStyle().PaddingTop(1)
	headerStyle = lipgloss.NewStyle().PaddingRight(3)
	opts        = options.CommonOptions{
		BotName: "",
	}
)

func New(o *options.CommonOptions) model {
	url := "https://api.github.com/users/botwayorg/repos"
	respone, status, _, err := api.BasicGet(url, "GET", "", "", "", "", false, 0, nil)

	opts.BotName = o.BotName

	if err != nil {
		fmt.Println(err.Error())
	}

	if status == "404" || status == "401" || strings.Contains(respone, "404") {
		fmt.Println("404")
		os.Exit(0)
	}

	repos := gjson.Get(respone, "#.name")
	repo_size := gjson.Get(respone, "#.size")

	rStr := repos.String()
	rsStr := repo_size.String()

	for i := 0; i < 2; i++ {
		rStr, _ = sjson.Delete(rStr, "6")
		rsStr, _ = sjson.Delete(rsStr, "6")
	}

	w, h, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		w = 80
		h = 24
	}

	top, right, bottom, left := styleDoc.GetPadding()

	w = w - left - right
	h = h - top - bottom

	tbl := table.New([]string{
		headerStyle.Render("Template Name"),
		headerStyle.Render("Size (MB)"),
		headerStyle.Render("Platform"),
		headerStyle.Render("Language"),
	}, w, h)

	tbl.Styles.SelectedRow = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E90FF"))

	rows := make([]table.Row, int(gjson.Get(rStr, "#").Int()))

	for i := 0; i < int(gjson.Get(rStr, "#").Int()); i++ {
		r := gjson.Get(rStr, fmt.Sprint(i))

		rInfo := strings.Split(r.Str, "-")

		if rInfo[1] == "deno" {
			rInfo[1] = "typescript"
		}

		icon, color := icons.GetIcon(rInfo[1])
	
		langIcon := lipgloss.NewStyle().Width(2).Render(fmt.Sprintf("%s%s\033[0m ", color, icon))

		rows[i] = table.SimpleRow{
			r,
			gjson.Get(rsStr, fmt.Sprint(i)),
			rInfo[0],
			fmt.Sprintf("%s %s", langIcon, rInfo[1]),
		}
	}

	tbl.SetRows(rows)

	return model{table: tbl}
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	raw := strings.ReplaceAll(fmt.Sprint(m.table.SelectedRow()), "[", "")
	raw = strings.ReplaceAll(raw, "]", "")
	raw = strings.ReplaceAll(raw, "  ", " ")
	data := strings.Split(raw, " ")

	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			top, right, bottom, left := styleDoc.GetPadding()

			m.table.SetSize(
				msg.Width-left-right,
				msg.Height-top-bottom,
			)

		case tea.KeyMsg:
			switch msg.String() {
				case "enter":
					return m.buildBot(opts.BotName, data[2], data[4])

				case "ctrl+c":
					return m, tea.Quit
			}
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return styleDoc.Render(
		m.table.View(),
	)
}

func (m model) buildBot(botName, platform, lang string) (tea.Model, tea.Cmd) {
	l := strings.Split(lang, " ")[0]

	termenv.AltScreen()

	errx := tea.NewProgram(pm.NewPM(botName, platform, l), tea.WithAltScreen()).Start()
	if errx != nil {
		fmt.Fprintln(os.Stderr, errx)
		os.Exit(1)
	}

	return m, tea.Quit
}
