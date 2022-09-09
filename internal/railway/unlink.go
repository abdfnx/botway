package railway

import (
	"context"
	"fmt"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) Unlink(ctx context.Context, req *entity.CommandRequest) error {
	projectCfg, _ := h.ctrl.GetProjectConfigs(ctx)

	project, err := h.ctrl.GetProject(ctx, projectCfg.Project)

	if err != nil {
		return err
	}

	err = h.cfg.RemoveProjectConfigs(projectCfg)

	if err != nil {
		return err
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Printf(" 🔓 Disconnected from %s \n", constants.BOLD.Render(project.Name))

	return nil
}
