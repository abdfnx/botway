package pm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/dashboard/icons"
	"github.com/abdfnx/botway/internal/options"
	"github.com/abdfnx/botway/internal/pipes/new/config"
	"github.com/abdfnx/looker"
	"github.com/abdfnx/resto/core/api"
	table "github.com/calyptia/go-bubble-table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/muesli/termenv"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var (
	styleDoc    = lipgloss.NewStyle().PaddingTop(1)
	headerStyle = lipgloss.NewStyle().PaddingRight(3)
	opts        = options.CommonOptions{
		BotName: "",
	}
)

type model struct {
	table    table.Model
	lang     string
	botName  string
	platform string
}

func NewPM(botName, platform, lang string) model {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		w = 80
		h = 24
	}

	top, right, bottom, left := styleDoc.GetPadding()

	w = w - left - right
	h = h - top - bottom

	tbl := table.New([]string{
		headerStyle.Render("Package Manager"),
		headerStyle.Render("Website"),
		headerStyle.Render("Is Installed Locally ?"),
		headerStyle.Render("Langauge"),
	}, w, h)

	tbl.Styles.SelectedRow = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E90FF"))

	r := 1

	if lang == "nodejs" {
		r = 3
	} else if lang == "python" {
		r = 2
	} else if lang == "rust" {
		r = 2	
	}

	rows := make([]table.Row, r)

	for i := 0; i < r; i++ {
		icon, color := icons.GetIcon(lang)

		langIcon := lipgloss.NewStyle().Width(2).Render(fmt.Sprintf("%s%s\033[0m ", color, icon))

		isInstalledLocally := "false"
		installedColor := constants.RED_COLOR

		installedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(installedColor)).SetString(isInstalledLocally).String()

		if lang == "nodejs" {
			npm := "npm"
			yarn := "yarn"
			pnpm := "pnpm"

			_, nerr := looker.LookPath(npm)
			_, yerr := looker.LookPath(yarn)
			_, perr := looker.LookPath(pnpm)

			if nerr == nil || yerr == nil || perr == nil {
				isInstalledLocally = "true"
				color = constants.GREEN_COLOR
			}

			rows[0] = table.SimpleRow{
				npm,
				"npmjs.com",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}

			rows[1] = table.SimpleRow{
				yarn,
				"yarnpkg.com",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}

			rows[2] = table.SimpleRow{
				pnpm,
				"pnpm.io",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}
		} else if lang == "python" {
			pip := "pip"
			pipenv := "pipenv"

			_, perr := looker.LookPath(pip)
			_, ierr := looker.LookPath(pipenv)

			if perr == nil || ierr == nil {
				isInstalledLocally = "true"
				color = constants.GREEN_COLOR
			}

			rows[0] = table.SimpleRow{
				pip,
				"pip.pypa.io",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}

			rows[1] = table.SimpleRow{
				pipenv,
				"pipenv.pypa.io",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}
		} else if lang == "rust" {
			cargo := "cargo"
			fleet := "fleet"

			_, cerr := looker.LookPath(cargo)
			_, ferr := looker.LookPath(fleet)

			if cerr == nil || ferr == nil {
				isInstalledLocally = "true"
				color = constants.GREEN_COLOR
			}

			rows[0] = table.SimpleRow{
				cargo,
				"doc.rust-lang.org/cargo",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}

			rows[1] = table.SimpleRow{
				fleet,
				"fleet.rs",
				installedStyle,
				fmt.Sprintf("%s %s", langIcon, lang),
			}
		}
	}

	if lang == "nodejs" || lang == "python" || lang == "rust" {
		tbl.SetRows(rows)

		return model{
			table: tbl,
			lang: lang,
			botName: botName,
			platform: platform,
		}
	} else {
		return model{
			lang: lang,
			botName: botName,
			platform: platform,
		}
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
					raw := strings.ReplaceAll(fmt.Sprint(m.table.SelectedRow()), "[", "")
					raw = strings.ReplaceAll(raw, "]", "")
					data := strings.Split(raw, " ")

					return m.buildBot(m.botName, m.platform, m.lang, data[0])

				case "ctrl+c":
					return m, tea.Quit
			}
	}

	if m.lang == "nodejs" || m.lang == "python" || m.lang == "rust" {
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)

		return m, cmd
	} else {
		return m.buildBot(m.botName, m.platform, m.lang, "")
	}
}

func (m model) View() string {
	return styleDoc.Render(
		m.table.View(),
	)
}

func (m model) buildBot(botName, platform, lang, pm string) (tea.Model, tea.Cmd) {
	termenv.AltScreen()

	var conf = koanf.New(".")

	l := strings.Split(lang, " ")[0]

	if l == "typescript" || l == "nodejs" {
		l = "Node"
	} 

	if err := conf.Load(file.Provider(constants.BotwayConfigFile), json.Parser()); err != nil {
		log.Fatal(err)
	} else {
		if err := os.Mkdir(botName, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "src"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "config"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		botwayConfig := viper.New()
		deployConfig := viper.New()

		botwayConfig.AddConfigPath(botName)
		botwayConfig.SetConfigName(".botway")
		botwayConfig.SetConfigType("yaml")

		deployConfig.AddConfigPath(botName)
		deployConfig.SetConfigName("deploy")
		deployConfig.SetConfigType("hcl")

		botwayConfig.SetDefault("author", conf.String("user.github_username"))
		botwayConfig.SetDefault("bot.lang", l)
		botwayConfig.SetDefault("bot.name", botName)
		botwayConfig.SetDefault("bot.package_manager", pm)
		botwayConfig.SetDefault("bot.type", platform)
		botwayConfig.SetDefault("bot.start_cmd", "la")
		botwayConfig.SetDefault("bot.version", "0.1.0")

		dockerImage := conf.String("user.docker_id") + "/" + botName

		deployConfig.SetDefault("docker.image", dockerImage)
		deployConfig.SetDefault("docker.cmds.build", "docker build -t " + dockerImage + " .")
		deployConfig.SetDefault("docker.cmds.run", "docker run -it " + dockerImage)

		if platform == "discord" {
			guildsFile := os.WriteFile(filepath.Join(botName, "config", "guilds.json"), []byte("{}"), 0644)

			if guildsFile != nil {
				panic(guildsFile)
			}
		}

		if err := botwayConfig.SafeWriteConfig(); err != nil {
			if os.IsNotExist(err) {
				err = botwayConfig.WriteConfig()

				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if err := deployConfig.SafeWriteConfig(); err != nil {
			if os.IsNotExist(err) {
				err = deployConfig.WriteConfig()

				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if err := botwayConfig.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal(err)
			}
		}

		if err := deployConfig.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal(err)
			}
		}

		respone, status, _, err := api.BasicGet("https://raw.githubusercontent.com/github/gitignore/main/" + strings.Title(l) + ".gitignore", "GET", "", "", "", "", true, 0, nil)

		if err != nil {
			fmt.Println(err.Error())
		}

		if status == "404" || status == "401" {
			fmt.Println("404")
			os.Exit(0)
		}

		if l == "typescript" {
			respone += "fileloader.ts"
		}

		dotGitIgnoreFileContent := respone + "\n*.lock"

		if l == "rust" && pm == "fleet" {
			dotGitIgnoreFileContent += "\n.cargo\nfleet.toml"
		}

		dotGitIgnoreFile := os.WriteFile(filepath.Join(botName, ".gitignore"), []byte(dotGitIgnoreFileContent), 0644)

		if dotGitIgnoreFile != nil {
			log.Fatal(dotGitIgnoreFile)
		}

		DiscordHandler(m, botName, pm)
		SlackHandler(m, botName, pm)
		TelegramHandler(m, botName, pm)

		pwd, _ := os.Getwd()

		pwd = filepath.Join(pwd, botName)

		config.AddBotToConfig(botName, platform, pwd)
	}

	return m, tea.Quit
}
