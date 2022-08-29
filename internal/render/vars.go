package render

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
	"github.com/tidwall/gjson"
)

func Vars(isGetEnvCmd bool, args []string) {
	url := "https://api.render.com/v1/services/" + serviceId + "/env-vars?limit=100"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		fmt.Println(constants.HEADING + constants.BOLD.Render(serviceName+" Environment Variables"))

		envs := gjson.Get(string(body), "#.envVar")

		for _, e := range envs.Array() {
			k := gjson.Get(e.String(), "key")
			v := gjson.Get(e.String(), "value")

			if isGetEnvCmd {
				if args[0] == k.String() {
					fmt.Println(v.String())
				}
			} else {
				fmt.Println(constants.BOLD.Render(k.String()) + ": " + v.String())
			}
		}
	} else {
		fmt.Println(string(body))
	}
}

func SetEnvVars(args []string) {
	url := "https://api.render.com/v1/services/" + serviceId + "/env-vars"

	payload := strings.NewReader("")

	var key, value string

	for _, kvPair := range args {
		parts := strings.SplitN(kvPair, "=", 2)

		if len(parts) != 2 {
			panic("invalid variables invocation. See --help")
		}

		key = parts[0]
		value = parts[1]

		payload = strings.NewReader(fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"}]", key, value))
	}

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
		fmt.Println(constants.HEADING + constants.BOLD.Render(fmt.Sprintf("Updated %s for \"%s\"", key, botwaygo.GetBotInfo("bot.name"))))
		fmt.Println(constants.HEADING + constants.BOLD.Render(fmt.Sprintf("%s: %s", key, value)))
	} else {
		fmt.Println(string(body))
	}
}
