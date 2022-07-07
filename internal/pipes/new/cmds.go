package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/new/config"
	"github.com/abdfnx/resto/core/api"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/spf13/viper"
)

func updatePlatforms(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.PlatformChoice += 1
			if m.PlatformChoice > 2 {
				m.PlatformChoice = 3
			}

		case "k", "up":
			m.PlatformChoice -= 1
			if m.PlatformChoice < 0 {
				m.PlatformChoice = 0
			}

		case "enter":
			m.Platform = true
			return m, frame()
		}
	}

	return m, nil
}

func updateLangs(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.LangChoice += 1

			if m.PlatformChoice == 2 {
				if m.LangChoice > 1 {
					m.LangChoice = 1
				}
			} else {
				if m.PlatformChoice == 0 {
					if m.LangChoice > 8 {
						m.LangChoice = 8
					}
				} else {
					if m.LangChoice > 7 {
						m.LangChoice = 7
					}
				}
			}

		case "k", "up":
			m.LangChoice -= 1

			if m.LangChoice < 0 {
				m.LangChoice = 0
			}

		case "enter":
			m.Lang = true

			return m, frame()
		}
	}

	return m, nil
}

func updatePMs(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.PMCoice += 1

			if m.LangChoice == 1 {
				if m.PlatformChoice == 2 {
					if m.PMCoice > 3 {
						m.PMCoice = 3
					}
				} else {
					if m.PMCoice > 0 {
						m.PMCoice = 0
					}
				}
			} else if m.LangChoice == 0 || m.LangChoice == 2 {
				if m.PMCoice > 3 {
					m.PMCoice = 3
				}
			} else {
				if m.PMCoice > 1 {
					m.PMCoice = 1
				}
			}

		case "k", "up":
			m.PMCoice -= 1
			if m.PMCoice < 0 {
				m.PMCoice = 0
			}

		case "enter":
			m.PM = true
			return m, frame()
		}
	}

	return m, nil
}

func buildBot(msg tea.Msg, m model, botName string) (tea.Model, tea.Cmd) {
	fmt.Println(finalView(m))

	var conf = koanf.New(".")

	l := ""

	if m.LangChoice == 0 {
		l = "Python"
	} else if m.LangChoice == 1 {
		if m.PlatformChoice == 2 {
			l = "Node"
		} else {
			l = "Go"
		}
	} else if m.LangChoice == 2 || m.LangChoice == 5 {
		l = "Node"
	} else if m.LangChoice == 3 {
		l = "Ruby"
	} else if m.LangChoice == 4 {
		l = "Rust"
	} else if m.LangChoice == 6 {
		l = "C#"
	} else if m.LangChoice == 7 {
		l = "Dart"
	} else if m.LangChoice == 8 {
		l = "Crystal"
	}

	if err := conf.Load(file.Provider(constants.BotwayConfigFile), json.Parser()); err != nil {
		log.Fatal(err)
	} else {
		if err := os.Mkdir(opts.BotName, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(opts.BotName, "src"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(opts.BotName, "config"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		dockerImage := conf.String("user.docker_id") + "/" + opts.BotName

		botwayConfig := viper.New()

		botwayConfig.AddConfigPath(opts.BotName)
		botwayConfig.SetConfigName(".botway")
		botwayConfig.SetConfigType("yaml")

		botwayConfig.SetDefault("author", conf.String("user.github_username"))
		botwayConfig.SetDefault("bot.lang", BotLang(m))
		botwayConfig.SetDefault("bot.name", opts.BotName)
		botwayConfig.SetDefault("bot.package_manager", BotPM(m))
		botwayConfig.SetDefault("bot.type", BotType(m))
		botwayConfig.SetDefault("bot.start_cmd", BotStartCmd(m))
		botwayConfig.SetDefault("bot.version", "0.1.0")

		botwayConfig.SetDefault("docker.image", dockerImage)
		botwayConfig.SetDefault("docker.enable_buildkit", false)
		botwayConfig.SetDefault("docker.cmds.build", "docker build -t "+dockerImage+" .")
		botwayConfig.SetDefault("docker.cmds.run", "docker run -it "+dockerImage)

		if m.PlatformChoice == 0 {
			guildsFile := os.WriteFile(filepath.Join(opts.BotName, "config", "guilds.json"), []byte("{}"), 0644)

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

		if err := botwayConfig.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal(err)
			}
		}

		respone, status, _, err := api.BasicGet("https://raw.githubusercontent.com/github/gitignore/main/"+l+".gitignore", "GET", "", "", "", "", true, 0, nil)

		if err != nil {
			fmt.Println(err.Error())
		}

		if status == "404" || status == "401" {
			fmt.Println("404")
			os.Exit(0)
		}

		if m.LangChoice == 5 {
			respone += "fileloader.ts"
		}

		if m.LangChoice == 6 {
			respone = CSharpGitIgnore()
		}

		if m.LangChoice == 8 {
			respone = `/docs/
/lib/
/bin/
/.shards/
*.dwarf

# Libraries don't need dependency lock
# Dependencies will be locked in applications that use them
/shard.lock`
		}

		dotGitIgnoreFileContent := respone + "\n*.lock"

		if BotLang(m) == "rust" && BotPM(m) == "fleet" {
			dotGitIgnoreFileContent += "\n.cargo\nfleet.toml"
		}

		dotGitIgnoreFile := os.WriteFile(filepath.Join(opts.BotName, ".gitignore"), []byte(dotGitIgnoreFileContent), 0644)

		if dotGitIgnoreFile != nil {
			log.Fatal(dotGitIgnoreFile)
		}

		DiscordHandler(m)
		TelegramHandler(m)
		SlackHandler(m)

		pwd, _ := os.Getwd()

		pwd = filepath.Join(pwd, botName)

		config.AddBotToConfig(opts.BotName, BotType(m), pwd, BotLang(m))
	}

	return m, tea.Quit
}
