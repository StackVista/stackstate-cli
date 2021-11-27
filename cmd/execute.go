package cmd

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	pr "gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

const (
	ConfigErrorExitCode      = 2
	CommandExecutionExitCode = 1
)

func Execute(ctx context.Context) {
	printer := pr.NewStdPrinter()

	cli := &di.Deps{
		Printer: printer,
	}

	cmd := AllCommands(cli)

	zerolog.SetGlobalLevel(zerolog.Disabled)
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// needs to happen before run, but after execute
		// so flag-config bindings can take hold
		cfg, err := conf.ReadConf(cmd)
		if err != nil {
			printer.PrintErr(err)
			os.Exit(ConfigErrorExitCode)
		}
		cli.Config = &cfg

		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
		return nil
	}

	if err := cmd.ExecuteContext(ctx); err != nil {
		cli.Printer.PrintErr(err)
		os.Exit(CommandExecutionExitCode)
	}
}
