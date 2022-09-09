package railway

import (
	"context"
	"errors"

	"github.com/abdfnx/botway/internal/railway/project"
	"github.com/botwayorg/railway-api/entity"
	"github.com/botwayorg/railway-api/uuid"
)

func (h *Handler) Link(ctx context.Context, req *entity.CommandRequest) error {
	if len(req.Args) > 0 {
		// projectID provided as argument
		arg := req.Args[0]

		if uuid.IsValidUUID(arg) {
			project, err := h.ctrl.GetProject(ctx, arg)

			if err != nil {
				return err
			}

			return h.setProject(ctx, project)
		}

		project, err := h.ctrl.GetProjectByName(ctx, arg)

		if err != nil {
			return err
		}

		return h.setProject(ctx, project)
	}

	isLoggedIn, err := h.ctrl.IsLoggedIn(ctx)

	if err != nil {
		return err
	}

	if isLoggedIn {
		return h.linkFromID(ctx, req)
	} else {
		return errors.New("You are not logged in to Railway Cloud")
	}
}

func (h *Handler) linkFromID(ctx context.Context, req *entity.CommandRequest) error {
	projectID, err := project.Project()

	if err != nil {
		return err
	}

	project, err := h.ctrl.GetProject(ctx, projectID)

	if err != nil {
		return err
	}

	return h.setProject(ctx, project)
}
