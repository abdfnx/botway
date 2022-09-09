package railway

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
	"github.com/briandowns/spinner"
)

func avaliableDBs() {
	fmt.Println("\nAvailable databases:")
	fmt.Println("\n1- " + constants.BOLD.Render("PostgreSQL"))
	fmt.Println("2- " + constants.BOLD.Render("MongoDB"))
	fmt.Println("3- " + constants.BOLD.Render("Redis"))
	fmt.Println("4- " + constants.BOLD.Render("MySQL\n"))
}

func (h *Handler) Add(ctx context.Context, req *entity.CommandRequest) error {
	projectCfg, err := h.ctrl.GetProjectConfigs(ctx)
	if err != nil {
		return err
	}

	selectedPlugin := os.Args[3]

	if selectedPlugin == "" {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" You need to specify a database name"))
		avaliableDBs()
	} else if selectedPlugin != "postgres" && selectedPlugin != "mongodb" && selectedPlugin != "redis" && selectedPlugin != "mysql" {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" Invalid database name"))
		avaliableDBs()
	} else {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Suffix = fmt.Sprintf(" 🗄️ Adding %s plugin...", selectedPlugin)
		s.Start()

		plugin, err := h.ctrl.CreatePlugin(ctx, &entity.CreatePluginRequest{
			ProjectID: projectCfg.Project,
			Plugin:    selectedPlugin,
		})

		if err != nil {
			s.Stop()

			fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR") + " ")

			return err
		}

		s.Stop()

		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" 🎉 Created plugin " + constants.BOLD.Render(plugin.Name)))
	}

	return nil
}
