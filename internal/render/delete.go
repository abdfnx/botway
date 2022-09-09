package render

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/abdfnx/botway/constants"
	"github.com/tidwall/gjson"
)

func DeleteRenderService(serviceName string) {
	serviceId := gjson.Get(string(constants.RenderConfig), "projects."+serviceName+".id").String()

	url := "https://api.render.com/v1/services/" + serviceId

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 204 {
		fmt.Println(constants.HEADING + constants.BOLD.Render("Service Deleted successfully"))
	} else {
		fmt.Println(string(body))
	}
}
