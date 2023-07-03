package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdfnx/botway/internal/config"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/abdfnx/resto/core/api"
	"github.com/botwayorg/templates"
	"github.com/spf13/viper"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content("dockerfiles/blank.dockerfile", "botway", botName, "")
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

	botConfig := viper.New()

	dockerImage := "botway-local/" + opts.BotName

	botConfig.AddConfigPath(opts.BotName)
	botConfig.SetConfigName(".botway")
	botConfig.SetConfigType("yaml")

	author := config.Get("github.username")

	if author == "" {
		author = "botway"
	}

	botConfig.SetDefault("author", author)
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

		if lang == 6 {
			respone += "fileloader.ts"
		}

		if lang == 7 {
			respone = CSharpGitIgnore()
		}

		if lang == 8 || lang == 11 {
			respone += "\n.gradle\nbuild"
		}

		if lang == 14 && platform == 1 {
			respone += "\n.build\nPackage.resolved"
		}

		if lang == 15 {
			respone = `/docs/
/lib/
/bin/
/.shards/
*.dwarf

# Libraries don't need dependency lock
# Dependencies will be locked in applications that use them
/shard.lock`
		}

		dotGitIgnoreFileContent = respone + "\n*.lock\nbotway-tokens.env\n/botway.json"

		DiscordHandler(m)
		TelegramHandler(m)
		SlackHandler(m)
		TwitchHandler(m)
	} else {
		dockerFile := os.WriteFile(filepath.Join(opts.BotName, "Dockerfile"), []byte(DockerfileContent(opts.BotName, HostServiceName(m))), 0644)

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}
	}

	dotGitIgnoreFile := os.WriteFile(filepath.Join(opts.BotName, ".gitignore"), []byte(dotGitIgnoreFileContent), 0644)
	dotGitKeepFile := os.WriteFile(filepath.Join(opts.BotName, "config", ".gitkeep"), []byte(""), 0644)
	readmeFile := os.WriteFile(filepath.Join(opts.BotName, "README.md"), []byte(templates.Content("bot-readme.md", "resources", "", "")), 0644)
	dockerComposeFile := os.WriteFile(filepath.Join(opts.BotName, "docker-compose.yaml"), []byte(templates.Content("dockerfiles/compose/docker-compose.yaml", "botway", "", "")), 0644)

	if dotGitIgnoreFile != nil {
		log.Fatal(dotGitIgnoreFile)
	}

	if readmeFile != nil {
		log.Fatal(readmeFile)
	}

	if dotGitKeepFile != nil {
		log.Fatal(dotGitKeepFile)
	}

	if dockerComposeFile != nil {
		log.Fatal(dockerComposeFile)
	}

	pwd, _ := os.Getwd()

	pwd = filepath.Join(pwd, opts.BotName)

	AddBotToConfig(opts.BotName, BotType(m), pwd, BotLang(m), HostService(m))

	initx.UpdateConfig()
}
