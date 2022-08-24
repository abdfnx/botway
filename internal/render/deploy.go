package render

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/botwaygo"
	"github.com/tidwall/gjson"
)

func Deploy() {
	tools.SetupTokensInDocker()

	serviceId := gjson.Get(string(constants.BotwayConfig), "render.projects."+botwaygo.GetBotInfo("bot.name")+".id").String()

	UpdateTokens(serviceId)

	url := fmt.Sprintf("https://api.render.com/v1/services/%s/deploys", serviceId)

	payload := strings.NewReader("{\"clearCache\":\"do_not_clear\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 201 {
		fmt.Println(constants.HEADING + constants.BOLD.Render("Deploy started 🚀️"))
	} else {
		fmt.Println(string(body))
	}
}
