package render

import (
	"log"
	"os"

	"github.com/abdfnx/botway/constants"
	"github.com/tidwall/sjson"
)

func Logout() {
	clearUserConfig, _ := sjson.Delete(string(constants.RenderConfig), "user")

	remove := os.Remove(constants.RenderConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.RenderConfigFile, []byte(clearUserConfig), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}
