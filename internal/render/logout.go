package render

import (
	"log"
	"os"

	"github.com/abdfnx/botway/constants"
	"github.com/tidwall/sjson"
)

func Logout() {
	clearUserConfig, _ := sjson.Delete(string(constants.BotwayConfig), "render.user")

	remove := os.Remove(constants.BotwayConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.BotwayConfigFile, []byte(clearUserConfig), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}
