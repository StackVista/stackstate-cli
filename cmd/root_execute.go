package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

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
	printer := pr.NewPrinter()

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

		cli.Printer.SetUseColor(!cfg.NoColor)

		switch strings.ToUpper(cfg.Output) {
		case "JSON":
			cli.Printer.SetOutputType(pr.JSON)
		case "YAML":
			cli.Printer.SetOutputType(pr.YAML)
		case "AUTO":
			cli.Printer.SetOutputType(pr.Auto)
		default:
			return fmt.Errorf("invalid choice for output flag: %s. Must be JSON, YAML or Auto", cfg.Output)
		}

		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
		return nil
	}

	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.ExecuteContext(ctx); err != nil {
		cli.Printer.PrintErr(err)
		os.Exit(CommandExecutionExitCode)
	}
}
