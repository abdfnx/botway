package railway

import (
	"context"
	"fmt"

	"github.com/abdfnx/botway/constants"
	deleteProject "github.com/abdfnx/botway/internal/railway/delete"
	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) Down(ctx context.Context, req *entity.CommandRequest) error {
	isVerbose, err := req.Cmd.Flags().GetBool("verbose")

	if err != nil {
		// Verbose mode isn't a necessary flag; just default to false.
		isVerbose = false
	}

	if isVerbose {
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" Using verbose mode"))
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" Loading project configuration"))
	}

	projectConfig, err := h.ctrl.GetProjectConfigs(ctx)

	if err != nil {
		return err
	}

	if isVerbose {
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" Loading environment"))
	}

	environmentName, err := req.Cmd.Flags().GetString("environment")

	if err != nil {
		return err
	}

	environment, err := h.getEnvironment(ctx, environmentName)

	if err != nil {
		return err
	}

	if isVerbose {
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" Using environment " + constants.BOLD.Render(environment.Name)))
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" Loading project"))
	}

	project, err := h.ctrl.GetProject(ctx, projectConfig.Project)

	if err != nil {
		return err
	}

	shouldDelete, err := deleteProject.Delete(project.Name)

	if err != nil || !shouldDelete {
		return err
	}

	err = h.ctrl.Down(ctx, &entity.DownRequest{
		ProjectID:     project.Id,
		EnvironmentID: environment.Id,
	})

	if err != nil {
		return err
	}

	fmt.Print(constants.WARN_BACKGROUND.Render("WARN"))

	fmt.Println(constants.WARN_FOREGROUND.Render(" Deleted latest deployment for project " + constants.BOLD.Render(project.Name)))

	return nil
}
