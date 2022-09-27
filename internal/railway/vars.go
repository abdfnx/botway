package railway

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
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
	fmt.Print(ui.KeyValues(*envs, false))

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

	noRedeployHint, err := req.Cmd.Flags().GetBool("no-redeploy-hint")

	if err != nil {
		noRedeployHint = false
	}

	hidden, err := req.Cmd.Flags().GetBool("hidden")

	if err != nil {
		hidden = false
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

	fmt.Println(constants.HEADING + (fmt.Sprintf("%s %s for \"%s\"", operation, strings.Join(updatedEnvNames, ", "), environment.Name)))
	fmt.Print(ui.KeyValues(*variables, hidden))

	if !skipRedeploy {
		serviceID, err := h.ctrl.GetServiceIdByName(ctx, &serviceName)

		if err != nil {
			return err
		}

		err = h.redeployAfterVariablesChange(ctx, environment, serviceID, noRedeployHint)

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

	noRedeployHint, err := req.Cmd.Flags().GetBool("no-redeploy-hint")

	if err != nil {
		noRedeployHint = false
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

		err = h.redeployAfterVariablesChange(ctx, environment, serviceID, noRedeployHint)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) redeployAfterVariablesChange(ctx context.Context, environment *entity.Environment, serviceID *string, noRedeployHint bool) error {
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
		if !noRedeployHint {
			fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
			fmt.Println(constants.INFO_FOREGROUND.Render(fmt.Sprintf(" Run %s to redeploy your project ", constants.COMMAND_FOREGROUND.Render("botway deploy"))))
		}

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

func UpdateTokens(botPath, botType string) {
	setVarCmd := "botway vars set --no-redeploy-hint --hidden "

	if botType == "discord" {
		setVarCmd += "DISCORD_TOKEN=" + botwaygo.GetToken() + " DISCORD_CLIENT_ID=" + botwaygo.GetAppId()
	} else if botType == "slack" {
		setVarCmd += "SLACK_TOKEN=" + botwaygo.GetToken() + " SLACK_APP_TOKEN=" + botwaygo.GetAppId() + " SIGNING_SECRET=" + botwaygo.GetSecret()
	} else if botType == "telegram" {
		setVarCmd += "TELEGRAM_TOKEN=" + botwaygo.GetToken()
	} else if botType == "twitch" {
		setVarCmd += "TWITCH_OAUTH_TOKEN=" + botwaygo.GetToken() + " TWITCH_CLIENT_ID=" + botwaygo.GetAppId() + " TWITCH_CLIENT_SECRET=" + botwaygo.GetSecret()
	}

	cmd := exec.Command("bash", "-c", setVarCmd)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", setVarCmd)
	}

	cmd.Dir = botPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
