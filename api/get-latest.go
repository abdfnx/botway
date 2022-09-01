package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/briandowns/spinner"
)

func GetLatest() string {
	url := "https://get-latest.onrender.com/abdfnx/botway"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error creating request: %s \n", err.Error())
		os.Exit(0)
	}

	suffix := " 🔍 Checking for updates..."

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = suffix
	s.Start()

	client := httpClient.HttpClient()
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error sending request: %s", err.Error())
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error reading response: %s", err.Error())
	}

	s.Stop()

	return string(b)
}
