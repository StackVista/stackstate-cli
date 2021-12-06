package di

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

// Depedency Injection context for the CLI
type Deps struct {
	Config  *conf.Conf
	Client  *sts.APIClient
	Printer printer.Printer
	Context context.Context
}

func NewDeps() Deps {
	return Deps{
		Config:  nil,
		Client:  nil,
		Printer: printer.NewPrinter(),
		Context: nil,
	}
}

func CmdRunEWithDeps(cli *Deps, runFn func(*Deps, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd, args)
	}
}
