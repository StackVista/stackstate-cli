package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hierynomus/taipan"

	home "github.com/mitchellh/go-homedir"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
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
	SetVerboseLogging(cmd, cfg)

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "🎃 %s\n", color.Red(err))
		os.Exit(2)
	}
}

func SetVerboseLogging(cmd *cobra.Command, cfg *config.Config) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	f := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetCount("verbose")
		if verbose > 0 {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
		ret := f(cmd, args)
		log.Ctx(cmd.Context()).Info().Msg(fmt.Sprintf("Loaded config %+v", cfg))
		return ret
	}
}
