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
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
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

func AllCommands(cli *di.Context) *cobra.Command {
	cmd := RootCommand()
	cmd.AddCommand(VersionCommand())
	cmd.AddCommand(ScriptCommand(cli))

	return cmd
}

func Execute(ctx context.Context) {
	cfg := &config.Config{}
	cli := &di.Context{
		Config: cfg,
		Client: nil,
	}

	cmd := AllCommands(cli)

	// set verbosity using zerolog
	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetCount("verbose")
		if verbose > 0 {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		}
	}

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

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "🎃 %s\n", color.Red(err))
		os.Exit(2)
	}
}
