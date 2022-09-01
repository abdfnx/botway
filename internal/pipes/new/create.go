package new

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/new/config"
	"github.com/abdfnx/botway/templates"
	"github.com/abdfnx/resto/core/api"
	"github.com/spf13/viper"
)

func DockerfileContent(botName string) string {
	return templates.Content("dockerfiles/blank.dockerfile", "botway", botName)
}

func NewBot(m model, l string, platform, lang int) {
	if err := os.Mkdir(opts.BotName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(opts.BotName, "src"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if l == "Swift" || l == "Java" {
		os.RemoveAll(filepath.Join(opts.BotName, "src"))
	}

	if err := os.Mkdir(filepath.Join(opts.BotName, "config"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	botwayConfig := viper.New()
	botConfig := viper.New()

	botwayConfig.SetConfigType("json")
	botwayConfig.ReadConfig(bytes.NewBuffer(constants.BotwayConfig))

	dockerImage := botwayConfig.GetString("docker.id") + "/" + opts.BotName

	botConfig.AddConfigPath(opts.BotName)
	botConfig.SetConfigName(".botway")
	botConfig.SetConfigType("yaml")

	botConfig.SetDefault("author", botwayConfig.GetString("github.username"))
	botConfig.SetDefault("bot.lang", BotLang(m))
	botConfig.SetDefault("bot.name", opts.BotName)
	botConfig.SetDefault("bot.host_service", HostService(m))

	if BotPM(m) != "continue" {
		botConfig.SetDefault("bot.package_manager", BotPM(m))
	}

	botConfig.SetDefault("bot.type", BotType(m))
	botConfig.SetDefault("bot.start_cmd", BotStartCmd(m))

	botConfig.SetDefault("docker.image", dockerImage)
	botConfig.SetDefault("docker.enable_buildkit", true)
	botConfig.SetDefault("docker.cmds.build", "docker build -t "+dockerImage+" .")
	botConfig.SetDefault("docker.cmds.run", "docker run -it "+dockerImage)

	if platform == 0 {
		guildsFile := os.WriteFile(filepath.Join(opts.BotName, "config", "guilds.json"), []byte("{}"), 0644)

		if guildsFile != nil {
			panic(guildsFile)
		}
	}

	if err := botConfig.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = botConfig.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := botConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}

	dotGitIgnoreFileContent := ""

	if botConfig.GetString("bot.lang") != blankLangMessage {
		respone, status, _, err := api.BasicGet("https://raw.githubusercontent.com/github/gitignore/main/"+l+".gitignore", "GET", "", "", "", "", true, 0, nil)

		if err != nil {
			fmt.Println(err.Error())
		}

		if status == "404" || status == "401" {
			fmt.Println("404")
			os.Exit(0)
		}

		if lang == 5 {
			respone += "fileloader.ts"
		}

		if lang == 6 {
			respone = CSharpGitIgnore()
		}

		if lang == 9 || lang == 10 {
			respone += "\n.gradle\nbuild"
		}

		if lang == 13 && platform == 1 {
			respone += "\n.build\nPackage.resolved"
		}

		if lang == 14 {
			respone = `/docs/
/lib/
/bin/
/.shards/
*.dwarf

# Libraries don't need dependency lock
# Dependencies will be locked in applications that use them
/shard.lock`
		}

		dotGitIgnoreFileContent = respone + "\n*.lock\nbotway-tokens.env"

		if BotLang(m) == "rust" && BotPM(m) == "fleet" {
			dotGitIgnoreFileContent += "\nfleet.toml"
		}

		DiscordHandler(m)
		TelegramHandler(m)
		SlackHandler(m)
	} else {
		dockerFile := os.WriteFile(filepath.Join(opts.BotName, "Dockerfile"), []byte(DockerfileContent(opts.BotName)), 0644)

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}
	}

	dotGitIgnoreFile := os.WriteFile(filepath.Join(opts.BotName, ".gitignore"), []byte(dotGitIgnoreFileContent), 0644)

	if dotGitIgnoreFile != nil {
		log.Fatal(dotGitIgnoreFile)
	}

	readmeFile := os.WriteFile(filepath.Join(opts.BotName, "README.md"), []byte(templates.Content("bot-readme.md", "resources", "")), 0644)

	if readmeFile != nil {
		log.Fatal(readmeFile)
	}

	pwd, _ := os.Getwd()

	pwd = filepath.Join(pwd, opts.BotName)

	config.AddBotToConfig(opts.BotName, BotType(m), pwd, BotLang(m), HostService(m))
}
