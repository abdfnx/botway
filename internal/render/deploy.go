package render

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/abdfnx/botway/constants"
)

func Deploy() {
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

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 201 {
		fmt.Println(constants.HEADING + constants.BOLD.Render("Deploy started 🚀️"))
	} else {
		fmt.Println(string(body))
	}
}
