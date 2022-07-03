package dashboard

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/abdfnx/botway/internal/dashboard/components/common"
)

func openInLinuxBrowser(url string) error {
	var err error
	providers := []string{"xdg-open", "x-www-browser", "www-browser", "wslview"}

	for _, provider := range providers {
		if _, err = exec.LookPath(provider); err == nil {
			err = exec.Command(provider, url).Start()

			if err != nil {
				return err
			}

			return nil
		}
	}

	return &exec.Error{Name: strings.Join(providers, ","), Err: exec.ErrNotFound}
}

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = openInLinuxBrowser(url)

	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	case "darwin":
		err = exec.Command("open", url).Start()

	case "android":
		err = exec.Command("termux-open-url", url).Start()

	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		common.LogErrorf("could not open browser: %s", err)
	}
}
