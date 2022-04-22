package di

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

// Dependency Injection context for the CLI
type Deps struct {
	Config    *conf.Conf
	Printer   printer.Printer
	Context   context.Context
	Client    StackStateClient
	IsVerBose bool
}

func NewDeps() Deps {
	return Deps{
		Client:  nil,
		Config:  nil,
		Printer: printer.NewPrinter(),
		Context: nil,
	}
}

func (cli *Deps) CmdRunE(runFn func(*Deps, *cobra.Command) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd)
	}
}

type CmdWithApiFn = func(
	*cobra.Command,
	*Deps,
	*stackstate_client.APIClient,
	*stackstate_client.ServerInfo,
) common.CLIError

func (cli *Deps) CmdRunEWithApi(
	runFn CmdWithApiFn) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		api, serverInfo, err := cli.Client.Connect()
		if err != nil {
			return common.NewConnectError(err)
		}

		return runFn(cmd, cli, api, serverInfo)
	}
}
