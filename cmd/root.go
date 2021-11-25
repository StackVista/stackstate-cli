package cmd

import (
	"context"
	"os"

	"github.com/hierynomus/taipan"

	home "github.com/mitchellh/go-homedir"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
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

	return cmd
}

func AllCommands(cli *di.Deps) *cobra.Command {
	cmd := RootCommand()
	cmd.AddCommand(VersionCommand())
	cmd.AddCommand(ScriptCommand(cli))

	return cmd
}

func Execute(ctx context.Context) {
	cfg := &config.Config{}
	cli := &di.Deps{
		Config:  cfg,
		Client:  nil,
		Printer: printer.NewStdPrinter(),
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

	homeFolder, err := home.Expand("~/.stackstate")
	if err != nil {
		cli.Printer.PrintErr(err)
		os.Exit(1)
	}

	taipanConfig := &taipan.Config{
		DefaultConfigName:  "cli",
		ConfigurationPaths: []string{".", homeFolder},
		EnvironmentPrefix:  EnvPrefix,
		AddConfigFlag:      true,
		ConfigObject:       cfg,
		PrefixCommands:     true,
	}

	t := taipan.New(taipanConfig)
	t.Inject(cmd)

	if err := cmd.ExecuteContext(ctx); err != nil {
		cli.Printer.PrintErr(err)
		os.Exit(2)
	}
}
