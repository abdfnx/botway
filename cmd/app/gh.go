package app

import (
	"github.com/botwayorg/gh/api"
	"github.com/botwayorg/gh/context"
	"github.com/botwayorg/gh/core/ghrepo"
	aCmd "github.com/botwayorg/gh/pkg/cmd/auth"
	"github.com/botwayorg/gh/pkg/cmd/factory"
	cCmd "github.com/botwayorg/gh/pkg/cmd/gh-config"
	rCmd "github.com/botwayorg/gh/pkg/cmd/gh-repo"
	"github.com/botwayorg/gh/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func Auth(f *cmdutil.Factory) *cobra.Command {
	cmd := aCmd.NewCmdAuth(f)
	return cmd
}

func GHConfig(f *cmdutil.Factory) *cobra.Command {
	cmd := cCmd.NewCmdConfig(f)
	return cmd
}

func Repo(f *cmdutil.Factory) *cobra.Command {
	repoResolvingCmdFactory := *f
	repoResolvingCmdFactory.BaseRepo = resolvedBaseRepo(f)

	cmd := rCmd.NewCmdRepo(&repoResolvingCmdFactory)

	return cmd
}

func resolvedBaseRepo(f *cmdutil.Factory) func() (ghrepo.Interface, error) {
	return func() (ghrepo.Interface, error) {
		httpClient, err := f.HttpClient()
		if err != nil {
			return nil, err
		}

		apiClient := api.NewClientFromHTTP(httpClient)

		remotes, err := f.Remotes()
		if err != nil {
			return nil, err
		}

		repoContext, err := context.ResolveRemotesToRepos(remotes, apiClient, "")
		if err != nil {
			return nil, err
		}

		baseRepo, err := repoContext.BaseRepo(f.IOStreams)
		if err != nil {
			return nil, err
		}

		return baseRepo, nil
	}
}

var NewGHConfigCmd = GHConfig(factory.New())
var NewGHRepoCmd = Repo(factory.New())
var GitHubCmd = Auth(factory.New())
