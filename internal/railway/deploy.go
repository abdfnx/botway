package railway

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
	"github.com/botwayorg/railway-api/entity"
	CLIErrors "github.com/botwayorg/railway-api/errors"
	"github.com/botwayorg/railway-api/ui"
	"github.com/briandowns/spinner"
)

func (h *Handler) Delpoy(ctx context.Context, req *entity.CommandRequest) error {
	CheckBuildKit()

	isVerbose, err := req.Cmd.Flags().GetBool("verbose")
	if err != nil {
		// Verbose mode isn't a necessary flag; just default to false.
		isVerbose = false
	}

	serviceName, err := req.Cmd.Flags().GetString("service")
	if err != nil {
		return err
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

	src := projectConfig.ProjectPath

	if src == "" {
		// When deploying with a project token, the project path is empty
		src = "."
	}

	UpdateTokens(src, botwaygo.GetBotInfo("bot.type"))

	fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
	fmt.Println(constants.INFO_FOREGROUND.Render(" Uploading directory " + constants.BOLD.Render(src)))

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

	serviceId := ""

	if serviceName != "" {
		for _, service := range project.Services {
			if service.Name == serviceName {
				serviceId = service.ID
			}
		}

		if serviceId == "" {
			return CLIErrors.ServiceNotFound
		}
	}

	// If service has not been provided via flag, prompt for it
	if serviceId == "" {
		if isVerbose {
			fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
			fmt.Println(constants.INFO_FOREGROUND.Render(" Loading services"))
		}

		service, err := ui.PromptServices(project.Services)

		if err != nil {
			return err
		}

		if service != nil {
			serviceId = service.ID
		}
	}

	_, err = os.ReadFile(".railwayignore")

	if err == nil {
		if isVerbose {
			fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
			fmt.Println(constants.INFO_FOREGROUND.Render(" Using ignore file .railwayignore"))
		}
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 📡 Laying tracks in the clouds..."
	s.Start()

	res, err := h.ctrl.Upload(ctx, &entity.UploadRequest{
		ProjectID:     projectConfig.Project,
		EnvironmentID: environment.Id,
		ServiceID:     serviceId,
		RootDir:       src,
	})

	if err != nil {
		return err
	} else {
		s.FinalMSG = constants.SUCCESS_BACKGROUND.Render("SUCCESS") + " ☁️ Build logs available at " + constants.BOLD.Render(res.URL) + "\n"
		s.Stop()
	}

	detach, err := req.Cmd.Flags().GetBool("detach")

	if err != nil {
		return err
	}

	if detach {
		return nil
	}

	for i := 0; i < 3; i++ {
		err = h.ctrl.GetActiveBuildLogs(ctx, 0)

		if err == nil {
			break
		}

		time.Sleep(time.Duration(i) * 250 * time.Millisecond)
	}

	fmt.Println(constants.SUCCESS_FOREGROUND.Render("\n======= Build Completed  ======\n\n"))

	err = h.ctrl.GetActiveDeploymentLogs(ctx, 1000)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
	fmt.Println(constants.INFO_FOREGROUND.Render(" Deployment logs available at " + constants.COMMAND_FOREGROUND.Render(res.URL)))
	fmt.Println(constants.INFO_FOREGROUND.Render("OR run `botway logs` to tail them here\n"))

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))

	if res.DeploymentDomain != "" {
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" ☁️ Deployment live at " + constants.COMMAND_FOREGROUND.Render(h.ctrl.GetFullUrlFromStaticUrl(res.DeploymentDomain))))
	} else {
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" ☁️ Deployment is live"))
	}

	os.RemoveAll("botway.json")

	return nil
}
