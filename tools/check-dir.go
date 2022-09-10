package tools

import (
	"fmt"
	"os"

	"github.com/abdfnx/botway/constants"
)

func CheckDir() {
	if _, err := os.Stat(".botway.yaml"); err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Print(" ")
		panic(constants.FAIL_FOREGROUND.Render("You need to run this command in your bot directory"))
	}
}
