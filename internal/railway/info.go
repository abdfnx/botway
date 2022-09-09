package railway

import (
	"context"
	"fmt"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
	"github.com/botwayorg/railway-api/errors"
)

func (h *Handler) Info(ctx context.Context, req *entity.CommandRequest) error {
	projectCfg, err := h.ctrl.GetProjectConfigs(ctx)

	if err != nil {
		return err
	}

	project, err := h.ctrl.GetProject(ctx, projectCfg.Project)

	if err != nil {
		return err
	}

	if project != nil {
		fmt.Printf("%s: %s\n", constants.BOLD.Render("Projects"), constants.PRIMARY_FOREGROUND.Render(project.Name))

		environment, err := h.ctrl.GetCurrentEnvironment(ctx)

		if err != nil {
			return err
		}

		fmt.Printf("%s: %s\n", constants.BOLD.Render("Enviroment"), constants.SUCCESS_FOREGROUND.Render(environment.Name))

		if len(project.Plugins) > 0 {
			fmt.Println(constants.BOLD.Render("Plugins:"))

			for i := range project.Plugins {
				plugin := project.Plugins[i]

				if plugin.Name == "env" {
					// legacy plugin
					continue
				}

				fmt.Printf("%s\n", constants.SUCCESS_FOREGROUND.Render(plugin.Name))
			}
		}

		if len(project.Services) > 0 {
			fmt.Println(constants.BOLD.Render("Services:"))

			for i := range project.Services {
				fmt.Printf("%s\n", constants.SUCCESS_FOREGROUND.Render(project.Services[i].Name))
			}
		}
	} else {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(" " + constants.FAIL_FOREGROUND.Render(fmt.Sprint(errors.ProjectConfigNotFound)))
	}

	return nil
}
