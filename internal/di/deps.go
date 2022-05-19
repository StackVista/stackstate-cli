package di

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/client"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

// Dependency Injection context for the CLI
type Deps struct {
	Config    *conf.Conf
	Printer   printer.Printer
	Context   context.Context
	Client    client.StackStateClient
	Clock     pflags.Clock
	IsVerBose bool
	IsJson    bool
	NoColor   bool
	Version   string
	Commit    string
	BuildDate string
	CLIType   string
}

func (cli *Deps) CmdRunE(runFn func(*Deps, *cobra.Command) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd)
	}
}

type CmdWithApiFn = func(
	*cobra.Command,
	*Deps,
	*stackstate_api.APIClient,
	*stackstate_api.ServerInfo,
) common.CLIError

func (cli *Deps) CmdRunEWithApi(
	runFn CmdWithApiFn) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cli.Config == nil {
			err := cli.LoadConfig(cmd)
			if err != nil {
				return err
			}
		}

		if cli.Client == nil {
			err := cli.LoadClient(cmd, cli.Config.URL, cli.Config.ApiPath, cli.Config.ApiToken)
			if err != nil {
				return err
			}
		}

		api, serverInfo, err := cli.Client.Connect()
		if err != nil {
			return err
		}

		return runFn(cmd, cli, api, serverInfo)
	}
}

// needs to happen before command run, but after cobra command execute
// so flag-config bindings can take hold
func (cli *Deps) LoadConfig(cmd *cobra.Command) common.CLIError {
	cfg, err := conf.ReadConf(cmd)
	if err != nil {
		return err
	}
	cli.Config = &cfg
	log.Info().Msg(fmt.Sprintf("Loaded config %+v", cli.Config))
	return nil
}

func (cli *Deps) LoadClient(cmd *cobra.Command, apiURL string, apiPath string, apiToken string) common.CLIError {
	cli.Client, cli.Context = client.NewStackStateClient(cmd.Context(), cli.IsVerBose, cli.Printer, apiURL, apiPath, apiToken)
	return nil
}
