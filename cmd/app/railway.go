package app

import (
	"github.com/abdfnx/botway/tools"
	"github.com/spf13/cobra"
)

func RailwayCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "railway",
		Short:   "Manage your work with Railway",
		Aliases: []string{"rw"},
	}

	cmd.AddCommand(RailwayLogoutCMD())
	cmd.AddCommand(RailwayInfoCMD())
	cmd.AddCommand(RailwayLinkCMD())
	cmd.AddCommand(RailwayUnLinkCMD())

	return cmd
}

func RailwayLogoutCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout of your Railway account",
		RunE:  Contextualize(handler.Logout, handler.Panic),
	}

	return cmd
}

func RailwayLinkCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "link",
		Short:  "Associate existing project with current directory, may specify projectId as an argument",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		RunE:   Contextualize(handler.Link, handler.Panic),
	}

	return cmd
}

func RailwayUnLinkCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "unlink",
		Short:  "Disassociate project from current directory",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		RunE:   Contextualize(handler.Unlink, handler.Panic),
	}

	return cmd
}

func RailwayInfoCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "info",
		Short:  "Show information about the current project",
		PreRun: func(cmd *cobra.Command, args []string) { tools.CheckDir() },
		RunE:   Contextualize(handler.Info, handler.Panic),
	}

	return cmd
}
