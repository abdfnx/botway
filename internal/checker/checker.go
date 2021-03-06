package checker

import (
	"fmt"
	"strings"

	"github.com/abdfnx/botway/api"
	"github.com/abdfnx/botway/cmd/factory"
	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/looker"
)

func Check(buildVersion string) {
	cliFactory := factory.New()
	stderr := cliFactory.IOStreams.ErrOut

	latestVersion := api.GetLatest()
	isFromHomebrew := isUnderHomebrew()
	isFromUsrBinDir := isUnderUsr()
	isFromScoop := isUnderScoop()
	isFromAppData := isUnderAppData()

	var command = func() string {
		if isFromHomebrew {
			return "brew upgrade botway"
		} else if isFromUsrBinDir {
			return "curl -sL https://bit.ly/botway | bash"
		} else if isFromScoop {
			return "scoop update botway"
		} else if isFromAppData {
			return "iwr -useb https://bit.ly/bw-win | iex"
		}

		return ""
	}

	if buildVersion != latestVersion {
		fmt.Fprintf(stderr, "\n%s %s → %s",
			constants.WARN_FOREGROUND.Render("There's a new version of ")+
				constants.PRIMARY_FOREGROUND.Render("botway")+
				constants.WARN_FOREGROUND.Render(" is avalaible:"),
			constants.PRIMARY_FOREGROUND.Render(buildVersion),
			constants.PRIMARY_FOREGROUND.Render(latestVersion)+
				"\n",
		)

		if command() != "" {
			fmt.Fprintf(stderr, constants.WARN_FOREGROUND.Render("To upgrade, run: %s"), constants.COMMAND_FOREGROUND.Render(command())+"\n")
		}
	}
}

var botwayExe, _ = looker.LookPath("botway")

func isUnderHomebrew() bool {
	return strings.Contains(botwayExe, "brew")
}

func isUnderUsr() bool {
	return strings.Contains(botwayExe, "usr")
}

func isUnderAppData() bool {
	return strings.Contains(botwayExe, "AppData")
}

func isUnderScoop() bool {
	return strings.Contains(botwayExe, "scoop")
}
