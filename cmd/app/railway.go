package app

import (
	"github.com/spf13/cobra"
)

func RailwayCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "railway",
		Short: "Manage your work with Railway",
		Aliases: []string{"rw"},
	}

	cmd.AddCommand(RailwayInfoCMD())
	cmd.AddCommand(RailwayLinkCMD())
	cmd.AddCommand(RailwayUnLinkCMD())

	return cmd
}

func RailwayLinkCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Associate existing project with current directory, may specify projectId as an argument",
		RunE:  Contextualize(handler.Link, handler.Panic),
	}

	return cmd
}

func RailwayUnLinkCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink",
		Short: "Disassociate project from current directory",
		RunE:  Contextualize(handler.Unlink, handler.Panic),
	}

	return cmd
}

func RailwayInfoCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Show information about the current project",
		RunE:  Contextualize(handler.Info, handler.Panic),
	}

	return cmd
}
