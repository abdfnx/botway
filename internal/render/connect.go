package render

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func ConnectService() {
	id := gjson.Get(string(constants.BotwayConfig), "render.user.id").String()
	apiToken := gjson.Get(string(constants.BotwayConfig), "render.user.api_token").String()

	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	serviceName := viper.GetString("bot.name")

	serviceName = strings.ReplaceAll(serviceName, " ", "%20")

	url := fmt.Sprintf("https://api.render.com/v1/services?name=%s&type=web_service&env=docker&ownerId=%s&limit=20", serviceName, id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	serviceID := gjson.Get(string(body), "0.service.id").String()
	serviceSlug := gjson.Get(string(body), "0.service.slug").String()
	serviceRepo := gjson.Get(string(body), "0.service.repo").String()

	renderPath := "render.projects." + serviceSlug

	service, _ := sjson.Set(string(constants.BotwayConfig), renderPath+".id", serviceID)
	addRepoToservice, _ := sjson.Set(service, renderPath+".repo", serviceRepo)

	remove := os.Remove(constants.BotwayConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.BotwayConfigFile, []byte(addRepoToservice), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}
}
