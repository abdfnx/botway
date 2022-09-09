package railway

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
	"github.com/botwayorg/railway-api/ui"
	"github.com/briandowns/spinner"
)

func (h *Handler) Variables(ctx context.Context, req *entity.CommandRequest) error {
	serviceName, err := req.Cmd.Flags().GetString("service")
	if err != nil {
		return err
	}

	envs, err := h.ctrl.GetEnvsForCurrentEnvironment(ctx, &serviceName)
	if err != nil {
		return err
	}

	environment, err := h.ctrl.GetCurrentEnvironment(ctx)

	if err != nil {
		return err
	}

	fmt.Print(constants.HEADING + (fmt.Sprintf("%s Environment Variables", environment.Name)))
	fmt.Print(ui.KeyValues(*envs))

	return nil
}

func (h *Handler) VariablesGet(ctx context.Context, req *entity.CommandRequest) error {
	serviceName, err := req.Cmd.Flags().GetString("service")
	if err != nil {
		return err
	}

	envs, err := h.ctrl.GetEnvsForCurrentEnvironment(ctx, &serviceName)

	if err != nil {
		return err
	}

	for _, key := range req.Args {
		fmt.Println(envs.Get(key))
	}

	return nil
}

func (h *Handler) VariablesSet(ctx context.Context, req *entity.CommandRequest) error {
	serviceName, err := req.Cmd.Flags().GetString("service")

	if err != nil {
		return err
	}

	skipRedeploy, err := req.Cmd.Flags().GetBool("skip-redeploy")

	if err != nil {
		// The flag is optional; default to false.
		skipRedeploy = false
	}

	replace, err := req.Cmd.Flags().GetBool("replace")

	if err != nil {
		// The flag is optional; default to false.
		replace = false
	}

	yes, err := req.Cmd.Flags().GetBool("yes")

	if err != nil {
		// The flag is optional; default to false.
		yes = false
	}

	if replace && !yes {
		fmt.Println(ui.Bold(ui.RedText(fmt.Sprintf("Warning! You are about to fully replace all your variables for the service '%s'.", serviceName)).String()))

		confirm, err := ui.PromptYesNo("Continue?")

		if err != nil {
			return err
		}

		if !confirm {
			return nil
		}
	}

	variables := &entity.Envs{}
	updatedEnvNames := make([]string, 0)

	for _, kvPair := range req.Args {
		parts := strings.SplitN(kvPair, "=", 2)

		if len(parts) != 2 {
			return errors.New("invalid variables invocation. See --help")
		}

		key := parts[0]
		value := parts[1]

		variables.Set(key, value)
		updatedEnvNames = append(updatedEnvNames, key)
	}

	err = h.ctrl.UpdateEnvs(ctx, variables, &serviceName, replace)

	if err != nil {
		return err
	}

	environment, err := h.ctrl.GetCurrentEnvironment(ctx)
	if err != nil {
		return err
	}

	operation := "Updated"

	if replace {
		operation = "Replaced existing variables with"
	}

	fmt.Print(constants.HEADING + (fmt.Sprintf("%s %s for \"%s\"", operation, strings.Join(updatedEnvNames, ", "), environment.Name)))
	fmt.Print(ui.KeyValues(*variables))

	if !skipRedeploy {
		serviceID, err := h.ctrl.GetServiceIdByName(ctx, &serviceName)

		if err != nil {
			return err
		}

		err = h.redeployAfterVariablesChange(ctx, environment, serviceID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) VariablesDelete(ctx context.Context, req *entity.CommandRequest) error {
	serviceName, err := req.Cmd.Flags().GetString("service")
	if err != nil {
		return err
	}

	skipRedeploy, err := req.Cmd.Flags().GetBool("skip-redeploy")
	if err != nil {
		// The flag is optional; default to false.
		skipRedeploy = false
	}

	err = h.ctrl.DeleteEnvs(ctx, req.Args, &serviceName)
	if err != nil {
		return err
	}

	environment, err := h.ctrl.GetCurrentEnvironment(ctx)
	if err != nil {
		return err
	}

	fmt.Print(constants.HEADING + (fmt.Sprintf("Deleted %s for \"%s\"", strings.Join(req.Args, ", "), environment.Name)))

	if !skipRedeploy {
		serviceID, err := h.ctrl.GetServiceIdByName(ctx, &serviceName)
		if err != nil {
			return err
		}

		err = h.redeployAfterVariablesChange(ctx, environment, serviceID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) redeployAfterVariablesChange(ctx context.Context, environment *entity.Environment, serviceID *string) error {
	deployments, err := h.ctrl.GetDeployments(ctx)
	if err != nil {
		return err
	}

	// Don't redeploy if we don't yet have any deployments
	if len(deployments) == 0 {
		return nil
	}

	// Don't redeploy if the latest deploy for environment came from up
	latestDeploy := deployments[0]
	if latestDeploy.Meta == nil || latestDeploy.Meta.Repo == "" {
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(fmt.Sprintf(" Run %s to redeploy your project ", constants.COMMAND_FOREGROUND.Render("botway deploy"))))

		return nil
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" 📟 Redeploying \"%s\" with new variables", environment.Name)
	s.Start()

	err = h.ctrl.DeployEnvironmentTriggers(ctx, serviceID)
	if err != nil {
		return err
	}

	s.FinalMSG = constants.SUCCESS_BACKGROUND.Render("SUCCESS") + " Deploy triggered"
	s.Stop()

	fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
	fmt.Println(constants.SUCCESS_FOREGROUND.Render(fmt.Sprintf("☁️ Deploy Logs available at %s\n", constants.COMMAND_FOREGROUND.Render(h.ctrl.GetProjectDeploymentsURL(ctx, latestDeploy.ProjectID)))))

	return nil
}
