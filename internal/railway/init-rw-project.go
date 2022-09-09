package railway

import (
	"context"
	"fmt"
	"os"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
	"github.com/botwayorg/railway-api/ui"
)

func (h *Handler) initNew(ctx context.Context, req *entity.CommandRequest) error {
	os.Chdir(os.Args[2])

	project, err := h.ctrl.CreateProject(ctx, &entity.CreateProjectRequest{
		Name: &os.Args[2],
	})

	if err != nil {
		return err
	}

	err = h.cfg.SetNewProject(project.Id)
	if err != nil {
		return err
	}

	environment, err := ui.PromptEnvironments(project.Environments)
	if err != nil {
		return err
	}

	err = h.cfg.SetEnvironment(environment.Id)
	if err != nil {
		return err
	}

	// Check if a .env exists, if so prompt uploading it
	err = h.ctrl.AutoImportDotEnv(ctx)
	if err != nil {
		return err
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + os.Args[2] + " Created succuessfully at Railway Cloud"))

	return nil
}

func (h *Handler) setProject(ctx context.Context, project *entity.Project) error {
	err := h.cfg.SetNewProject(project.Id)
	if err != nil {
		return err
	}

	environment, err := ui.PromptEnvironments(project.Environments)
	if err != nil {
		return err
	}

	err = h.cfg.SetEnvironment(environment.Id)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) Init(ctx context.Context, req *entity.CommandRequest) error {
	// Since init can be called by guests, ensure we can fetch a user first before calling. This prevents
	// us accidentally creating a temporary (guest) project if we have a token locally but our remote
	// session was deleted.
	_, err := h.ctrl.GetUser(ctx)
	if err != nil {
		return fmt.Errorf("%s\nRun %s", constants.FAIL_FOREGROUND.Render("Account required to init project"), ui.Bold("botway login railway"))
	}

	return h.initNew(ctx, req)
}
