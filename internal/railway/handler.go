package railway

import (
	"github.com/railwayapp/cli/configs"
	"github.com/railwayapp/cli/controller"
)

type Handler struct {
	ctrl *controller.Controller
	cfg  *configs.Configs
}

func NewRW() *Handler {
	return &Handler{
		ctrl: controller.New(),
		cfg:  configs.New(),
	}
}
