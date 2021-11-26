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
	EnvPrefix = "STS"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState Command Line Interface",
	}

	var verbosity int
	cmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Print more verbose logging.")
	cmd.PersistentFlags().String("api-url", "", "StackState API URL.")

	return cmd
}

func AllCommands(cli *di.Deps) *cobra.Command {
	cmd := RootCommand()
	cmd.AddCommand(VersionCommand())
	cmd.AddCommand(ScriptCommand(cli))

	return cmd
}

func Execute(ctx context.Context) {
	printer := pr.NewStdPrinter()

	cfg, err := conf.ReadConf()
	if err != nil {
		printer.PrintErr(err)
		os.Exit(1)
	}

	cli := &di.Deps{
		Config:  &cfg,
		Printer: printer,
	}

	cmd := AllCommands(cli)

	// set verbosity using zerolog
	zerolog.SetGlobalLevel(zerolog.Disabled)
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetCount("verbose")
		if verbose > 0 {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
		return nil
	}

	if err := cmd.ExecuteContext(ctx); err != nil {
		cli.Printer.PrintErr(err)
		os.Exit(2)
	}
}
