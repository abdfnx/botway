package new

import (
	"log"
	"os"

	"github.com/abdfnx/botway/constants"
	"github.com/tidwall/sjson"
)

func AddBotToConfig(botName, botType, botPath, botLang, hostService string) {
	botTypeContent, _ := sjson.Set(string(constants.BotwayConfig), "botway.bots."+botName+".type", botType)
	botPathContent, _ := sjson.Set(botTypeContent, "botway.bots."+botName+".path", botPath)
	botLangContent, _ := sjson.Set(botPathContent, "botway.bots."+botName+".lang", botLang)
	botHostServiceContent, _ := sjson.Set(botLangContent, "botway.bots."+botName+".host_service", hostService)
	addBotToBotsNamesContent, _ := sjson.Set(botHostServiceContent, "botway.bots_names.-1", botName)

	remove := os.Remove(constants.BotwayConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.BotwayConfigFile, []byte(addBotToBotsNamesContent), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}
