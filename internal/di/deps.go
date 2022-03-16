package di

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

// Depedency Injection context for the CLI
type Deps struct {
	Config    *conf.Conf
	Printer   printer.Printer
	Context   context.Context
	Client    *stackstate_client.APIClient
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

func CmdRunEWithDeps(cli *Deps, runFn func(*Deps, *cobra.Command, []string) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd, args)
	}
}
