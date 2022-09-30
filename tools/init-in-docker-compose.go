package tools

import (
	"bytes"
	"fmt"
	"os"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
	"github.com/spf13/viper"
)

func InitInDockerCompose(db string) {
	viper.SetConfigType("yaml")

	dockerCompose, err := os.ReadFile("docker-compose.yaml")

	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(dockerCompose))

	image := ""
	ports := ""

	if db == "pocketbase" {
		image = "botwayorg/pocketbase"
		ports = "8090:8090"
	} else if db == "surrealdb" {
		image = "surrealdb/surrealdb"
		ports = "8000:8000"
	}

	viper.Set("services."+botwaygo.GetBotInfo("bot.name")+".depends_on", []string{db})
	viper.Set("services."+db+".image", image)
	viper.Set("services."+db+".ports", []string{ports})

	err = viper.WriteConfigAs("docker-compose.yaml")

	if err != nil {
		panic(err)
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(" Initialized Successfully"))
	fmt.Print(constants.INFO_BACKGROUND.Render("NEXT STEP"))
	fmt.Println(constants.INFO_FOREGROUND.Render(" Add " + db + " SDK package to your bot 📦️"))
}
