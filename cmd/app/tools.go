package app

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/railway"
	rw_constants "github.com/botwayorg/railway-api/constants"
	"github.com/botwayorg/railway-api/entity"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	handler      = railway.NewRW()
	messageStyle = lipgloss.NewStyle().Foreground(constants.CYAN_COLOR)
)

func Contextualize(fn entity.HandlerFunction, panicFn entity.PanicFunction) entity.CobraFunction {
	return func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		defer func() {
			if rw_constants.IsDevVersion() {
				return
			}

			if r := recover(); r != nil {
				err := panicFn(ctx, fmt.Sprint(r), string(debug.Stack()), cmd.Name(), args)
				if err != nil {
					fmt.Println("Unable to relay panic to server. Are you connected to the internet?")
				}
			}
		}()

		req := &entity.CommandRequest{
			Cmd:  cmd,
			Args: args,
		}

		err := fn(ctx, req)

		if err != nil {
			// TODO: Make it *pretty*
			fmt.Println(err.Error())
			os.Exit(1) // Set non-success exit code on error
		}

		return nil
	}
}
