package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/internal/pipes/new/config"
	"github.com/abdfnx/resto/core/api"
	"github.com/abdfnx/tran/dfs"
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
						if m.LangChoice > 3 {
							m.LangChoice = 3
						}
					} else if m.PlatformChoice == 0 {
						if m.LangChoice > 5 {
							m.LangChoice = 5
						}
					} else {
						if m.LangChoice > 4 {
							m.LangChoice = 4
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
						if m.PMCoice > 0 {
							m.PMCoice = 0
						}
					} else if m.LangChoice == 2 {
						if m.PMCoice > 2 {
							m.PMCoice = 2
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

	homeDir, err := dfs.GetHomeDirectory()

	if err != nil {
		log.Fatal(err)
	}

	var conf = koanf.New(".")

	l := ""

	if m.LangChoice == 0 {
		l = "Python"
	} else if m.LangChoice == 1 {
		l = "Go"
	} else if m.LangChoice == 2 || m.LangChoice == 5 {
		l = "Node"
	} else if m.LangChoice == 3 {
		l = "Ruby"
	} else if m.LangChoice == 4 {
		l = "Rust"
	}

	if err := conf.Load(file.Provider(filepath.Join(homeDir, ".botway", "botway.json")), json.Parser()); err != nil {
		log.Fatal(err)
	} else {
		if err := os.Mkdir(opts.BotName, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(opts.BotName, "src"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(opts.BotName)
		viper.SetConfigName(".botway")
		viper.SetConfigType("yaml")

		viper.SetDefault("author", conf.String("user.github_username"))
		viper.SetDefault("bot.name", opts.BotName)
		viper.SetDefault("bot.version", "0.1.0")
		viper.SetDefault("bot.type", BotType(m))
		viper.SetDefault("bot.lang", BotLang(m))
		viper.SetDefault("bot.package_manager", BotPM(m))
		viper.SetDefault("docker.image", conf.String("user.docker_id") + "/" + opts.BotName)

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

		respone, status, _, err := api.BasicGet("https://raw.githubusercontent.com/github/gitignore/main/" + l + ".gitignore", "GET", "", "", "", "", true, 0, nil)

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

		dotGitIgnoreFileContent := []byte(respone)
		herokuFileContent := []byte(HerokuFile())

		dotGitIgnoreFile := os.WriteFile(filepath.Join(opts.BotName, ".gitignore"), dotGitIgnoreFileContent, 0644)
		dotDockerIgnoreFile := os.WriteFile(filepath.Join(opts.BotName, ".dockerignore"), dotGitIgnoreFileContent, 0644)
		herokuFile := os.WriteFile(filepath.Join(opts.BotName, "heroku.yaml"), herokuFileContent, 0644)

		if dotGitIgnoreFile != nil {
			log.Fatal(dotGitIgnoreFile)
		}

		if dotDockerIgnoreFile != nil {
			log.Fatal(dotDockerIgnoreFile)
		}

		if herokuFile != nil {
			log.Fatal(herokuFile)
		}

		DiscordHandler(m)
		TelegramHandler(m)
		SlackHandler(m)

		pwd, _ := os.Getwd()

		pwd = filepath.Join(pwd, botName)

		config.AddBotToConfig(opts.BotName, BotType(m), pwd)
	}

	return m, tea.Quit
}
