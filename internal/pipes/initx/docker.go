package initx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/botwaygo"
	"github.com/abdfnx/tran/dfs"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func DockerInit() {
	err := dfs.CreateDirectory(filepath.Join(constants.HomeDir, ".botway"))

	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(constants.BotwayDirPath)
	viper.SetConfigName("botway")
	viper.SetConfigType("json")

	botType := botwaygo.GetBotInfo("bot.type")
	bot_token := ""
	app_token := ""
	signing_secret := "SLACK_SIGNING_SECRET"
	cid := ""

	if botType == "discord" {
		bot_token = "DISCORD_TOKEN"
		app_token = "DISCORD_CLIENT_ID"
		cid = "bot_app_id"
	} else if botType == "slack" {
		bot_token = "SLACK_TOKEN"
		app_token = "SLACK_APP_TOKEN"
		cid = "bot_app_token"
	} else if botType == "telegram" {
		bot_token = "TELEGRAM_TOKEN"
	}

	env := viper.New()
	env.SetConfigType("env")

	secretsFile, serr := ioutil.ReadFile(filepath.Join("config", "botway-tokens.env"))

	if serr != nil {
		panic(serr)
	}

	env.ReadConfig(bytes.NewBuffer(secretsFile))

	get := func(value string) string {
		if botwaygo.GetBotInfo("bot.host_service") == "render.com" && tools.IsRunningInContainer() {
			return os.Getenv(value)
		}

		return env.GetString(bot_token)
	}

	viper.SetDefault("botway.bots."+botwaygo.GetBotInfo("bot.name")+".bot_token", get(bot_token))
	viper.SetDefault("botway.bots_names", []string{botwaygo.GetBotInfo("bot.name")})

	if botType != "telegram" {
		viper.SetDefault("botway.bots."+botwaygo.GetBotInfo("bot.name")+"."+cid, get(app_token))
	}

	if botType == "slack" {
		viper.SetDefault("botway.bots."+botwaygo.GetBotInfo("bot.name")+".signing_secret", get(signing_secret))
	}

	if botType == "discord" {
		if constants.Gerr != nil {
			panic(constants.Gerr)
		} else {
			guilds := gjson.Get(string(constants.Guilds), "guilds.#")

			for x := 0; x < int(guilds.Int()); x++ {
				server := gjson.Get(string(constants.Guilds), "guilds."+fmt.Sprint(x)).String()

				sgi := strings.ToUpper(server) + "_GUILD_ID"

				viper.Set("botway.bots."+botwaygo.GetBotInfo("bot.name")+".guilds."+server+".server_id", get(sgi))
			}
		}
	}

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

	fmt.Println(constants.HEADING + constants.BOLD.Render("Done 🐋️"))
}
