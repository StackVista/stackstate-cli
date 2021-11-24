package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hierynomus/taipan"

	home "github.com/mitchellh/go-homedir"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

const (
	EnvPrefix = "STS"
)

func RootCommand() *cobra.Command {
	var verbosity int

	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState Command Line Interface",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			switch verbosity {
			case 0: //nolint:gomnd
				// Nothing to do
			case 1: //nolint:gomnd
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			case 2: //nolint:gomnd
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			default:
				zerolog.SetGlobalLevel(zerolog.TraceLevel)
			}

			return nil
		},
	}

	cmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Print more verbose logging.")

	return cmd
}

func AllCommands(cfg *config.Config) *cobra.Command {
	cmd := RootCommand()
	cmd.AddCommand(VersionCommand(cfg))
	cmd.AddCommand(ScriptCommand(cfg))

	return cmd
}

func Execute(ctx context.Context) {
	cfg := &config.Config{}
	cmd := AllCommands(cfg)

	homeFolder, err := home.Expand("~/.stackstate")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", color.Red(err))
		os.Exit(1)
	}

	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

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
		fmt.Fprintf(os.Stderr, "🎃 %s\n", color.Red(err))
		os.Exit(2)
	}
}
