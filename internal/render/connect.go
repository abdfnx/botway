package render

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/abdfnx/botwaygo"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func UpdateTokens(serviceId string) {
	url := fmt.Sprintf("https://api.render.com/v1/services/%s/env-vars", serviceId)

	botType := botwaygo.GetBotInfo("bot.type")

	bot_token := ""
	app_token := ""
	secret_value := ""
	payload_content := ""

	if botType == "discord" {
		bot_token = "DISCORD_TOKEN"
		app_token = "DISCORD_CLIENT_ID"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken(), app_token, botwaygo.GetAppId())
	} else if botType == "slack" {
		bot_token = "SLACK_TOKEN"
		app_token = "SLACK_APP_TOKEN"
		secret_value = "SIGNING_SECRET"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken(), app_token, botwaygo.GetAppId(), secret_value, botwaygo.GetSecret())
	} else if botType == "telegram" {
		bot_token = "TELEGRAM_TOKEN"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken())
	} else if botType == "twitch" {
		bot_token = "TWITCH_OAUTH_TOKEN"
		app_token = "TWITCH_CLIENT_ID"
		secret_value = "TWITCH_CLIENT_SECRET"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken(), app_token, botwaygo.GetAppId(), secret_value, botwaygo.GetSecret())
	}

	payload := strings.NewReader(payload_content)
	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		fmt.Println(constants.HEADING + constants.BOLD.Render("Tokens updated successfully 🔑️"))
	} else {
		fmt.Println(string(body))
	}
}

func ConnectService() {
	url := fmt.Sprintf("https://api.render.com/v1/services?name=%s&type=web_service&env=docker&ownerId=%s&limit=20", serviceName, id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	body, _ := ioutil.ReadAll(res.Body)

	serviceId := gjson.Get(string(body), "0.service.id").String()
	serviceSlug := gjson.Get(string(body), "0.service.slug").String()
	serviceRepo := gjson.Get(string(body), "0.service.repo").String()

	renderPath := "projects." + serviceName

	service, _ := sjson.Set(string(constants.RenderConfig), renderPath+".id", serviceId)
	addSlug, _ := sjson.Set(service, renderPath+".slug", serviceSlug)
	addRepoToService, _ := sjson.Set(addSlug, renderPath+".repo", serviceRepo)

	remove := os.Remove(constants.RenderConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.RenderConfigFile, []byte(addRepoToService), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}

	UpdateTokens(serviceId)
	initx.UpdateConfig()
}
