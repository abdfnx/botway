package app

import (
	"log"
	"path/filepath"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/pipes/initx"
	"github.com/abdfnx/botway/tools"
	"github.com/abdfnx/tran/dfs"
	"github.com/spf13/cobra"
)

func DockerInitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker-init",
		Short: "Initialize ~/.botway for docker containers",
		Run: func(cmd *cobra.Command, args []string) {
			if opts.CopyFile {
				err := dfs.CreateDirectory(filepath.Join(constants.HomeDir, ".botway"))

				if err != nil {
					log.Fatal(err)
				}

				tools.Copy("botway.json", constants.BotwayConfigFile)
			} else {
				initx.DockerInit()
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.CopyFile, "copy-file", "", false, "Copy config file")

	return cmd
}
