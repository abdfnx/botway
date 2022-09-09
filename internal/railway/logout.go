package railway

import (
	"context"

	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) Logout(ctx context.Context, req *entity.CommandRequest) error {
	return h.ctrl.Logout(ctx)
}
