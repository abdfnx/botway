package railway

import (
	"context"
	"fmt"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) Login(ctx context.Context, req *entity.CommandRequest) error {
	user, err := h.ctrl.Login(ctx, true)

	if err != nil {
		return err
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Printf(" Logged in as %s (%s) 🎉\n", constants.BOLD.Render(user.Name), user.Email)

	return nil
}
