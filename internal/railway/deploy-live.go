package railway

import (
	"context"

	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) OpenApp(ctx context.Context, req *entity.CommandRequest) error {
	projectId, err := h.cfg.GetProject()

	if err != nil {
		return err
	}

	environmentId, err := h.cfg.GetCurrentEnvironment()

	if err != nil {
		return err
	}

	deployment, err := h.ctrl.GetLatestDeploymentForEnvironment(ctx, projectId, environmentId)

	if err != nil {
		return err
	}

	return h.ctrl.OpenStaticUrlInBrowser(deployment.StaticUrl)
}
