package di

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

type Context struct {
	Config *config.Config
	Client *sts.APIClient
}

func NewContext() Context {
	return Context{
		Config: nil,
		Client: nil,
	}
}

func CmdRunEWithCLIContext(cli *Context, runFn func(*Context, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd, args)
	}
}
