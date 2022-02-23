package di

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_old_client"
)

// Depedency Injection context for the CLI
type Deps struct {
	Config    *conf.Conf
	Printer   printer.Printer
	Context   context.Context
	IsVerBose bool
}

func NewDeps() Deps {
	return Deps{
		Config:  nil,
		Printer: printer.NewPrinter(),
		Context: nil,
	}
}

func CmdRunEWithDeps(cli *Deps, runFn func(*Deps, *cobra.Command, []string) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd, args)
	}
}

func (deps *Deps) NewStsClient() (stackstate_client.APIClient, context.Context) {
	configuration := stackstate_client.NewConfiguration()
	configuration.Servers[0] = stackstate_client.ServerConfiguration{
		URL:         deps.Config.ApiUrl,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = deps.IsVerBose

	client := stackstate_client.NewAPIClient(configuration)

	auth := make(map[string]stackstate_client.APIKey)
	auth["ApiToken"] = stackstate_client.APIKey{
		Key:    deps.Config.ApiToken,
		Prefix: "",
	}
	ctx := context.WithValue(
		deps.Context,
		stackstate_client.ContextAPIKeys,
		auth,
	)

	return *client, ctx
}

func (deps *Deps) NewOldStsClient() (stackstate_old_client.APIClient, context.Context) {
	configuration := stackstate_old_client.NewConfiguration()
	configuration.Servers[0] = stackstate_old_client.ServerConfiguration{
		URL:         deps.Config.ApiUrl,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = deps.IsVerBose

	client := stackstate_old_client.NewAPIClient(configuration)

	auth := make(map[string]stackstate_old_client.APIKey)
	auth["ApiToken"] = stackstate_old_client.APIKey{
		Key:    deps.Config.ApiToken,
		Prefix: "",
	}
	ctx := context.WithValue(
		deps.Context,
		stackstate_old_client.ContextAPIKeys,
		auth,
	)

	return *client, ctx
}
