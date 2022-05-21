package railway

import (
	"context"
	"fmt"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/briandowns/spinner"
	"github.com/railwayapp/cli/entity"
	"github.com/railwayapp/cli/ui"
)

func (h *Handler) Add(ctx context.Context, req *entity.CommandRequest) error {
	projectCfg, err := h.ctrl.GetProjectConfigs(ctx)
	if err != nil {
		return err
	}

	plugins, err := h.ctrl.GetAvailablePlugins(ctx, projectCfg.Project)

	if err != nil {
		return err
	}

	selectedPlugin, err := ui.PromptPlugins(plugins)
	if err != nil {
		return err
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" 🗄️ Adding %s plugin...", selectedPlugin)
	s.Start()

	s.FinalMSG = constants.SUCCESS_BACKGROUND.Render("SUCCESS") + " Added " + selectedPlugin + " plugin successfully"

	defer s.Stop()

	plugin, err := h.ctrl.CreatePlugin(ctx, &entity.CreatePluginRequest{
		ProjectID: projectCfg.Project,
		Plugin:    selectedPlugin,
	})

	if err != nil {
		return err
	}

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(" 🎉 Created plugin " + constants.INFO_FOREGROUND.Render(plugin.Name)))

	return nil
}
