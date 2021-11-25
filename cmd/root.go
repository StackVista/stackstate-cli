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
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
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
	InjectVerboseLogging(cmd, cli)
	InjectClient(cmd, cli)

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "🎃 %s\n", color.Red(err))
		os.Exit(2)
	}
}

func InjectVerboseLogging(cmd *cobra.Command, cli *di.Context) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	f := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetCount("verbose")
		if verbose > 0 {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
		ret := f(cmd, args)

		log.Ctx(cmd.Context()).Info().Msg(fmt.Sprintf("Loaded config %+v", cli.Config))

		return ret
	}
}

func InjectClient(cmd *cobra.Command, cli *di.Context) {
	f := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		ret := f(cmd, args)

		configuration := stackstate_client.NewConfiguration()
		configuration.Servers[0] = stackstate_client.ServerConfiguration{
			URL:         cli.Config.URL,
			Description: "",
			Variables:   nil,
		}
		client := stackstate_client.NewAPIClient(configuration)
		cli.Client = client

		auth := make(map[string]stackstate_client.APIKey)
		auth["ApiToken"] = stackstate_client.APIKey{
			Key:    cli.Config.ApiToken,
			Prefix: "",
		}
		context.WithValue(
			cmd.Context(),
			stackstate_client.ContextAPIKeys,
			auth,
		)

		return ret
	}
}
