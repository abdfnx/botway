package railway

import (
	"context"

	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) Logs(ctx context.Context, req *entity.CommandRequest) error {
	numLines, err := req.Cmd.Flags().GetInt32("lines")

	if err != nil {
		return err
	}

	return h.ctrl.GetActiveDeploymentLogs(ctx, numLines)
}
