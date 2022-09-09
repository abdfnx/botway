package railway

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/abdfnx/botway/constants"
	"github.com/botwayorg/railway-api/entity"
)

func (h *Handler) Connect(ctx context.Context, req *entity.CommandRequest) error {
	projectCfg, _ := h.ctrl.GetProjectConfigs(ctx)

	project, err := h.ctrl.GetProject(ctx, projectCfg.Project)

	if err != nil {
		return err
	}

	environment, err := h.ctrl.GetCurrentEnvironment(ctx)

	if err != nil {
		return err
	}

	fmt.Printf("🎉 Connecting to: %s %s\n", constants.INFO_FOREGROUND.Render(project.Name), constants.BOLD.Render(environment.Name))

	var plugin string

	if len(project.Plugins) > 2 && len(req.Args) == 0 {
		fmt.Print(constants.INFO_BACKGROUND.Render("INFO"))
		fmt.Println(constants.INFO_FOREGROUND.Render(" You've multiple databases, Please select a database to connect to:\n"))

		for _, plugin := range project.Plugins {
			if plugin.Name != "env" {
				fmt.Println("- " + plugin.Name)
			}
		}

		fmt.Println("\nTo connect to a database, run " + constants.COMMAND_FOREGROUND.Render("botway db connect <database>"))

		return nil
	} else {
		if len(req.Args) != 0 {
			plugin = req.Args[0]
		} else if len(project.Plugins) >= 2 {
			plugin = project.Plugins[0].Name
		}

		if !isPluginValid(plugin) {
			return fmt.Errorf("Invalid plugin: %s", plugin)
		}

		envs, err := h.ctrl.GetEnvsForCurrentEnvironment(ctx, nil)

		if err != nil {
			return err
		}

		command, connectEnv := buildConnectCommand(plugin, envs)

		if !commandExistsInPath(command[0]) {
			fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
			fmt.Println(constants.FAIL_FOREGROUND.Render(" " + command[0] + " was not found in $PATH."))

			return nil
		}

		cmd := exec.CommandContext(ctx, command[0], command[1:]...)

		cmd.Env = os.Environ()

		for k, v := range connectEnv {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%+v", k, v))
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		cmd.Stdin = os.Stdin
		catchSignals(ctx, cmd, nil)

		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func commandExistsInPath(cmd string) bool {
	// The error can be safely ignored because it indicates a failure to find the
	// command in $PATH.
	_, err := exec.LookPath(cmd)

	return err == nil
}

func isPluginValid(plugin string) bool {
	switch plugin {
	case "redis":
		fallthrough

	case "psql":
		fallthrough

	case "postgres":
		fallthrough

	case "postgresql":
		fallthrough

	case "mysql":
		fallthrough

	case "mongo":
		fallthrough

	case "mongodb":
		return true

	default:
		return false
	}
}

func buildConnectCommand(plugin string, envs *entity.Envs) ([]string, map[string]string) {
	var command []string
	var connectEnv map[string]string

	switch plugin {
	case "redis":
		command = []string{"redis-cli", "-u", (*envs)["REDIS_URL"]}

	case "psql":
		fallthrough

	case "postgres":
		fallthrough

	case "postgresql":
		connectEnv = map[string]string{
			"PGPASSWORD": (*envs)["PGPASSWORD"],
		}

		command = []string{
			"psql",
			"-U",
			(*envs)["PGUSER"],
			"-h",
			(*envs)["PGHOST"],
			"-p",
			(*envs)["PGPORT"],
			"-d",
			(*envs)["PGDATABASE"],
		}

	case "mongo":
		fallthrough

	case "mongodb":
		command = []string{
			"mongo",
			fmt.Sprintf(
				"mongodb://%s:%s@%s:%s",
				(*envs)["MONGOUSER"],
				(*envs)["MONGOPASSWORD"],
				(*envs)["MONGOHOST"],
				(*envs)["MONGOPORT"],
			),
		}

	case "mysql":
		command = []string{
			"mysql",
			fmt.Sprintf("-h%s", (*envs)["MYSQLHOST"]),
			fmt.Sprintf("-u%s", (*envs)["MYSQLUSER"]),
			fmt.Sprintf("-p%s", (*envs)["MYSQLPASSWORD"]),
			fmt.Sprintf("--port=%s", (*envs)["MYSQLPORT"]),
			"--protocol=TCP",
			(*envs)["MYSQLDATABASE"],
		}
	}

	return command, connectEnv
}
